package db

import (
	"fmt"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/utils"

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
	GetByToken(token, key string) (models.User, error)
	DropAll() error
	GetAll() ([]models.User, error)
	DeleteByID(userid string) error
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

func (u *UsersWrapper) GetByToken(token, key string) (models.User, error) {
	var (
		res models.User
		all []models.User
	)

	err := u.parent.db.Select().From(models.UsersTableName).All(&all)
	if err != nil {
		return res, err
	}

	for _, u := range all {
		userToken, err := utils.GenerateJWT(u.ID, "user", key)
		if err != nil {
			return res, err
		}

		if userToken == token {
			res = u
			return res, err
		}
	}

	return res, errs.UserWithTokenNotFoundErr.ToError()
}

func (u *UsersWrapper) GetAll() ([]models.User, error) {
	var res []models.User
	err := u.parent.db.Select().From(models.UsersTableName).All(&res)
	return res, err
}

func (u *UsersWrapper) DeleteByID(userid string) error {
	_, err := u.parent.db.Delete(models.UsersTableName, dbx.HashExp{"id": userid}).Execute()

	_, err = u.parent.db.Delete(models.ReadyForVoteTableName, dbx.HashExp{"user_id": userid}).Execute()

	_, err = u.parent.db.Delete(models.CompetitionsTableName, dbx.HashExp{"user_one_id": userid, "user_two_id": userid}).Execute()

	return err
}
