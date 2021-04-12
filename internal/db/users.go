package db

import (
	"github.com/SKilliu/novels-backend/db/models"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// UsersQ query interface, which provide access to DB functions.
type UsersQ interface {
	Insert(user models.User) error
	GetByEmail(email string) (models.User, error)
}

// UsersWrapper wraps interface.
type UsersWrapper struct {
	parent *DB
}

// UsersQ query interface getter.
func (d *DB) UsersQ() UsersQ {
	return &UsersWrapper{
		parent: &DB{d.db.Clone()},
	}
}

// Insert new user into database.
func (u *UsersWrapper) Insert(user models.User) error {
	return u.parent.db.Model(&user).Insert()
}

// GetByEmail finds the user in database by email
func (u *UsersWrapper) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := u.parent.db.Select().Where(dbx.HashExp{"email": email}).From(models.UsersTableName).One(&user)
	return user, err
}
