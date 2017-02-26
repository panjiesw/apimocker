package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
	"github.com/panjiesw/apimocker/errs"
)

var (
	UserBucket         = []byte("users")
	UserUsernameBucket = []byte("users.username")
	UserEmailBucket    = []byte("users.email")
	ProjectBucket      = []byte("projects")
)

type Datastore interface {
	UserSave(u *User) *errs.AError
	UserUsernameExist(username string) (bool, *errs.AError)
	UserEmailExist(email string) (bool, *errs.AError)
	UserGetByUsername(username string) (*User, *errs.AError)
	UserGetByEmail(email string) (*User, *errs.AError)
	UserGetByID(id uint64) (*User, *errs.AError)
}

type DB struct {
	*bolt.DB
}

func Open(path string) (*DB, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	dB := &DB{db}
	if dB.initialize(); err != nil {
		return nil, err
	}

	return dB, nil
}

func (d *DB) initialize() error {
	tx, err := d.Begin(true)
	if err != nil {
		return err
	}

	ub, err := tx.CreateBucketIfNotExists(UserBucket)
	if err != nil {
		return err
	}

	if _, err = tx.CreateBucketIfNotExists(ProjectBucket); err != nil {
		return err
	}

	if _, err = ub.CreateBucketIfNotExists(UserUsernameBucket); err != nil {
		return err
	}

	if _, err = ub.CreateBucketIfNotExists(UserEmailBucket); err != nil {
		return err
	}

	return tx.Commit()
}

// Itob returns an 8-byte big endian representation of v.
func Itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
