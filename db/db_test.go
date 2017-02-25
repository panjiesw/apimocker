package db_test

import (
	"log"
	"os"
	"testing"

	"io/ioutil"

	"encoding/json"

	"github.com/panjiesw/apimocker/db"
)

var d *testDB

func setupDB() {
	var err error
	d, err = newTestDB()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := d.Begin(true)
	if err != nil {
		log.Fatal(err)
	}

	ub := tx.Bucket(db.UserBucket)
	uub := ub.Bucket(db.UserUsernameBucket)
	ueb := ub.Bucket(db.UserEmailBucket)

	u1, err := json.Marshal(db.User{Username: "user1", Email: "user1@bar.com"})
	if err != nil {
		log.Fatal(err)
	}
	if err = ub.Put(db.Itob(1000), u1); err != nil {
		log.Fatal(err)
	}
	if err = uub.Put([]byte("user1"), db.Itob(1000)); err != nil {
		log.Fatal(err)
	}
	if err = ueb.Put([]byte("user1@bar.com"), db.Itob(1000)); err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

type testDB struct {
	*db.DB
}

func newTestDB() (*testDB, error) {
	fileName := tempFile()
	DB, err := db.Open(fileName)
	if err != nil {
		return nil, err
	}
	return &testDB{DB}, nil
}

func (d *testDB) closeTestDB() error {
	defer os.Remove(d.Path())
	return d.DB.Close()
}

func (d *testDB) mustClose() {
	if err := d.closeTestDB(); err != nil {
		panic(err)
	}
}

func tempFile() string {
	file, err := ioutil.TempFile("", "apimocker-")
	if err != nil {
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	if err := os.Remove(file.Name()); err != nil {
		log.Fatal(err)
	}

	return file.Name()
}

func TestMain(m *testing.M) {
	setupDB()
	result := m.Run()
	d.mustClose()
	os.Exit(result)
}
