package db_test

import (
	"testing"

	"github.com/panjiesw/apimocker/db"
)

func TestDB_UserSave(t *testing.T) {

	type args struct {
		u *db.User
	}

	tests := []struct {
		name    string
		args    *args
		wantErr bool
	}{
		{
			name:    "default",
			args:    &args{u: &db.User{Username: "foo", Email: "foo@bar.com"}},
			wantErr: false,
		},
		{
			name:    "username exist",
			args:    &args{u: &db.User{Username: "user1", Email: "foo2@bar.com"}},
			wantErr: true,
		},
		{
			name:    "email exists",
			args:    &args{u: &db.User{Username: "user2", Email: "user1@bar.com"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := d.UserSave(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("DB.UserSave() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && tt.args.u.ID == 0 {
				t.Error("DB.UserSave() no id")
			}
		})
	}
}

func TestDB_UserUsernameExist(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
		wantErr   bool
	}{
		{
			name:      "exists",
			args:      args{username: "user1"},
			wantExist: true,
			wantErr:   false,
		},
		{
			name:      "doesn't exist",
			args:      args{username: "user2"},
			wantExist: false,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := d.UserUsernameExist(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.UserUsernameExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExist != tt.wantExist {
				t.Errorf("DB.UserUsernameExist() = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func TestDB_UserEmailExist(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
		wantErr   bool
	}{
		{
			name:      "exists",
			args:      args{email: "user1@bar.com"},
			wantExist: true,
			wantErr:   false,
		},
		{
			name:      "doesn't exist",
			args:      args{email: "user2@bar.com"},
			wantExist: false,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, err := d.UserEmailExist(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.UserEmailExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExist != tt.wantExist {
				t.Errorf("DB.UserEmailExist() = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

// func TestDB_UserGetByUsername(t *testing.T) {
// 	type fields struct {
// 		DB *bolt.DB
// 	}
// 	type args struct {
// 		username string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *db.User
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &db.DB{
// 				DB: tt.fields.DB,
// 			}
// 			got, err := d.UserGetByUsername(tt.args.username)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DB.UserGetByUsername() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DB.UserGetByUsername() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDB_UserGetByEmail(t *testing.T) {
// 	type fields struct {
// 		DB *bolt.DB
// 	}
// 	type args struct {
// 		email string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *db.User
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &db.DB{
// 				DB: tt.fields.DB,
// 			}
// 			got, err := d.UserGetByEmail(tt.args.email)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DB.UserGetByEmail() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DB.UserGetByEmail() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDB_UserGetByID(t *testing.T) {
// 	type fields struct {
// 		DB *bolt.DB
// 	}
// 	type args struct {
// 		id uint64
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *db.User
// 		wantErr bool
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &db.DB{
// 				DB: tt.fields.DB,
// 			}
// 			got, err := d.UserGetByID(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DB.UserGetByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DB.UserGetByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
