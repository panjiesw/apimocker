package db

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/panjiesw/apimocker/errs"
)

// User represents apimocker users
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *DB) UserSave(u *User) error {
	err := d.Update(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		uub := ub.Bucket(UserUsernameBucket)
		ueb := ub.Bucket(UserEmailBucket)

		if ok, err := d.UserUsernameExist(u.Username); err != nil {
			return err
		} else if ok {
			return errs.ErrUsernameExists
		}

		if ok, err := d.UserEmailExist(u.Email); err != nil {
			return err
		} else if ok {
			return errs.ErrEmailExists
		}

		id, err := ub.NextSequence()
		if err != nil {
			return err
		}

		u.ID = id
		bID := Itob(u.ID)
		b, err := json.Marshal(u)
		if err != nil {
			return err
		}

		if err = ub.Put(bID, b); err != nil {
			return err
		}

		if err = uub.Put([]byte(u.Username), bID); err != nil {
			return err
		}

		if err = ueb.Put([]byte(u.Email), bID); err != nil {
			return err
		}

		return nil
	})
	return err
}

func (d *DB) UserUsernameExist(username string) (exist bool, err error) {
	err = d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		uub := ub.Bucket(UserUsernameBucket)
		if u := uub.Get([]byte(username)); u != nil {
			exist = true
		}
		return nil
	})
	return
}

func (d *DB) UserEmailExist(email string) (exist bool, err error) {
	err = d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		ueb := ub.Bucket(UserEmailBucket)
		if u := ueb.Get([]byte(email)); u != nil {
			exist = true
		}
		return nil
	})
	return
}

func (d *DB) UserGetByUsername(username string) (*User, error) {
	var user User
	err := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		uub := ub.Bucket(UserUsernameBucket)
		if u := uub.Get([]byte(username)); u != nil {
			if err := json.Unmarshal(u, &user); err != nil {
				return err
			}
		} else {
			return errs.ErrUsernameNotExists
		}
		return nil
	})
	return &user, err
}

func (d *DB) UserGetByEmail(email string) (*User, error) {
	var user User
	err := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		ueb := ub.Bucket(UserEmailBucket)
		if u := ueb.Get([]byte(email)); u != nil {
			if err := json.Unmarshal(u, &user); err != nil {
				return err
			}
		} else {
			return errs.ErrEmailNotExists
		}
		return nil
	})
	return &user, err
}

func (d *DB) UserGetByID(id uint64) (*User, error) {
	var user User
	err := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		if u := ub.Get(Itob(id)); u != nil {
			if err := json.Unmarshal(u, &user); err != nil {
				return err
			}
		} else {
			return errs.ErrIDNotExists
		}
		return nil
	})
	return &user, err
}
