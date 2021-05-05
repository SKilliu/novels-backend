package db

import (
	"fmt"

	"github.com/SKilliu/novels-backend/internal/db/models"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

var (
	checkUserByUsernameQuery = `SELECT EXISTS (SELECT * FROM users WHERE username = '%s' LIMIT 1);`
	checkUserByEmailQuery    = `SELECT EXISTS (SELECT * FROM users WHERE email = '%s' LIMIT 1);`
)

// UsersQ query interface, which provide access to DB functions.
type UsersQ interface {
	Insert(user models.User) error
	Update(user models.User) error
	GetByEmail(email string) (models.User, error)
	GetByUsername(username string) (models.User, error)
	CheckUserByUsername(username string) (models.IsExists, error)
	CheckUserByEmail(email string) (models.IsExists, error)
	GetByID(uid string) (models.User, error)
	GetByDeviceID(deviceID string) (models.User, error)
	GetAllForVote(userOneID, userTwoID string) ([]models.User, error)
	DropAll() error
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

func (u *UsersWrapper) Update(user models.User) error {
	return u.parent.db.Model(&user).Update()
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

func (u *UsersWrapper) CheckUserByEmail(email string) (models.IsExists, error) {
	var ex models.IsExists
	err := u.parent.db.NewQuery(fmt.Sprintf(checkUserByEmailQuery, email)).One(&ex)
	fmt.Println(ex)
	return ex, err
}

func (u *UsersWrapper) GetByID(uid string) (models.User, error) {
	var user models.User
	err := u.parent.db.Select().Where(dbx.HashExp{"id": uid}).From(models.UsersTableName).One(&user)
	return user, err
}

func (u UsersWrapper) GetByDeviceID(deviceID string) (models.User, error) {
	var user models.User
	err := u.parent.db.Select().Where(dbx.HashExp{"device_id": deviceID}).From(models.UsersTableName).One(&user)
	return user, err
}

func (u UsersWrapper) GetAllForVote(userOneID, userTwoID string) ([]models.User, error) {
	var res []models.User
	err := u.parent.db.Select().From(models.UsersTableName).Where(dbx.NotIn("id", userOneID, userTwoID)).All(&res)
	return res, err
}

func (u *UsersWrapper) DropAll() error {
	_, err := u.parent.db.Delete(models.UsersTableName, dbx.HashExp{}).Execute()
	return err
}
