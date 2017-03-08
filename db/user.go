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

func (d *DB) UserSave(u *User) *errs.AError {
	if err := d.Update(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		uub := ub.Bucket(UserUsernameBucket)
		ueb := ub.Bucket(UserEmailBucket)

		if ok, err := d.UserUsernameExist(u.Username); err != nil {
			return err
		} else if ok {
			return errs.ErrDBUsernameExists
		}

		if ok, err := d.UserEmailExist(u.Email); err != nil {
			return err
		} else if ok {
			return errs.ErrDBEmailExists
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
	}); err != nil {
		switch e := err.(type) {
		case *errs.AError:
			return e
		default:
			return errs.New("db", e.Error(), 500)
		}
	}
	return nil
}

func (d *DB) UserUsernameExist(username string) (exist bool, err *errs.AError) {
	if er := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		uub := ub.Bucket(UserUsernameBucket)
		if u := uub.Get([]byte(username)); u != nil {
			exist = true
		}
		return nil
	}); er != nil {
		switch e := er.(type) {
		case *errs.AError:
			err = e
		default:
			err = errs.New("db", e.Error(), 500)
		}
	}
	return
}

func (d *DB) UserEmailExist(email string) (exist bool, err *errs.AError) {
	if er := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		ueb := ub.Bucket(UserEmailBucket)
		if u := ueb.Get([]byte(email)); u != nil {
			exist = true
		}
		return nil
	}); er != nil {
		switch e := er.(type) {
		case *errs.AError:
			err = e
		default:
			err = errs.New("db", e.Error(), 500)
		}
	}
	return
}

func (d *DB) UserGetByUsername(username string, user *User) *errs.AError {
	if er := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		uub := ub.Bucket(UserUsernameBucket)
		if uid := uub.Get([]byte(username)); uid != nil {
			if u := ub.Get(uid); u != nil {
				if err := json.Unmarshal(u, user); err != nil {
					return err
				}
			}
		} else {
			return errs.ErrDBUsernameNotExists
		}
		return nil
	}); er != nil {
		switch e := er.(type) {
		case *errs.AError:
			return e
		default:
			return errs.New("db", e.Error(), 500)
		}
	}
	return nil
}

func (d *DB) UserGetByEmail(email string, user *User) *errs.AError {
	if er := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		ueb := ub.Bucket(UserEmailBucket)
		if uid := ueb.Get([]byte(email)); uid != nil {
			if u := ub.Get(uid); u != nil {
				if err := json.Unmarshal(u, &user); err != nil {
					return err
				}
			}
		} else {
			return errs.ErrDBEmailNotExists
		}
		return nil
	}); er != nil {
		switch e := er.(type) {
		case *errs.AError:
			return e
		default:
			return errs.New("db", e.Error(), 500)
		}
	}
	return nil
}

func (d *DB) UserGetByID(id uint64, user *User) *errs.AError {
	if er := d.View(func(tx *bolt.Tx) error {
		ub := tx.Bucket(UserBucket)
		if u := ub.Get(Itob(id)); u != nil {
			if err := json.Unmarshal(u, &user); err != nil {
				return err
			}
		} else {
			return errs.ErrDBIDNotExists
		}
		return nil
	}); er != nil {
		switch e := er.(type) {
		case *errs.AError:
			return e
		default:
			return errs.New("db", e.Error(), 500)
		}
	}
	return nil
}

func (d *DB) UserList(wrapper *Wrapper) *errs.AError {
	if err := d.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(UserBucket)
		c := b.Cursor()
		var offset = wrapper.Meta.Offset
		var us []User
		var usb = []byte("[")

		for k, v := c.First(); k != nil && len(us) < int(wrapper.Meta.Limit); k, v = c.Next() {
			if offset <= uint(0) {
				usb = append(usb, v...)
				usb = append(usb, comma...)
			}
			offset--
		}
		usb = append(usb[:len(usb)-1], []byte("]")...)

		if err := json.Unmarshal(usb, us); err != nil {
			return err
		}

		wrapper.Data = us
		wrapper.Meta.Count = uint(len(us))

		return nil
	}); err != nil {
		switch e := err.(type) {
		case *errs.AError:
			return e
		default:
			return errs.New("db", e.Error(), 500)
		}
	}
	return nil
}
