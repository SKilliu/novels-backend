package db

import (
	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

var checkBySocialID = `SELECT EXISTS (SELECT * FROM user_socials WHERE social_id = '%s' LIMIT 1);`

// UsersQ query interface, which provide access to DB functions.
type UserSocialsQ interface {
	Insert(us models.UserSocial) error
	Update(us models.UserSocial) error
	Delete(us models.UserSocial) error
	GetByID(uid string) (models.UserSocial, error)
}

// UsersWrapper wraps interface.
type UserSocialsWrapper struct {
	parent *DB
}

// UsersQ query interface getter.
func (d *DB) UserSocialsQ() UserSocialsQ {
	return &UserSocialsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (s *UserSocialsWrapper) Insert(us models.UserSocial) error {
	return s.parent.db.Model(&us).Insert()
}

func (s *UserSocialsWrapper) Update(us models.UserSocial) error {
	return s.parent.db.Model(&us).Update()
}

func (s *UserSocialsWrapper) Delete(us models.UserSocial) error {
	return s.parent.db.Model(&us).Delete()
}

func (s *UserSocialsWrapper) GetByID(sid string) (models.UserSocial, error) {
	var us models.UserSocial
	err := s.parent.db.Select().Where(dbx.HashExp{"social_id": sid}).From(models.UserSocialsTableName).One(&us)
	return us, err
}
