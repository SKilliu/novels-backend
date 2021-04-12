package db

import (
	"fmt"

	"github.com/SKilliu/novels-backend/internal/db/models"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

var checkUserByUsernameQuery = `SELECT EXISTS (SELECT * FROM users WHERE username = '%s' LIMIT 1);`

// UsersQ query interface, which provide access to DB functions.
type UsersQ interface {
	Insert(user models.User) error
	GetByEmail(email string) (models.User, error)
	GetByUsername(username string) (models.User, error)
	CheckUserByUsername(username string) (models.IsExists, error)
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

func (u *UsersWrapper) GetByUsername(username string) (models.User, error) {
	var user models.User
	err := u.parent.db.Select().Where(dbx.HashExp{"username": username}).From(models.UsersTableName).One(&user)
	return user, err
}

func (u *UsersWrapper) CheckUserByUsername(username string) (models.IsExists, error) {
	var ex models.IsExists
	err := u.parent.db.NewQuery(fmt.Sprintf(checkUserByUsernameQuery, username)).One(&ex)
	fmt.Println(ex)
	return ex, err
}
