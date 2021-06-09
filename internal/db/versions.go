package db

import (
	"github.com/SKilliu/novels-backend/internal/db/models"
)

// UsersQ query interface, which provide access to DB functions.
type VersionsQ interface {
	Insert(us models.Versions) error
	Update(us models.Versions) error
	Get() (models.Versions, error)
}

// UsersWrapper wraps interface.
type VersionsWrapper struct {
	parent *DB
}

// UsersQ query interface getter.
func (d *DB) VersionsQ() VersionsQ {
	return &VersionsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (v *VersionsWrapper) Insert(versions models.Versions) error {
	return v.parent.db.Model(&versions).Insert()
}

func (v *VersionsWrapper) Update(versions models.Versions) error {
	return v.parent.db.Model(&versions).Update()
}

func (v VersionsWrapper) Get() (models.Versions, error) {
	var result models.Versions
	err := v.parent.db.Select().From(models.VersionsTableName).Limit(1).One(&result)
	return result, err
}
