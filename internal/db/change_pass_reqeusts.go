package db

import (
	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type ChangePassRequestsQ interface {
	Insert(req models.ChangePassRequest) error
	Delete(req models.ChangePassRequest) error
	GetByID(reqID string) (models.ChangePassRequest, error)
}

type ChangePassRequestsWrapper struct {
	parent *DB
}

func (d *DB) ChangePassRequestsQ() ChangePassRequestsQ {
	return &ChangePassRequestsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (c *ChangePassRequestsWrapper) Insert(req models.ChangePassRequest) error {
	return c.parent.db.Model(&req).Insert()
}

func (c *ChangePassRequestsWrapper) Delete(req models.ChangePassRequest) error {
	return c.parent.db.Model(&req).Delete()
}

func (c *ChangePassRequestsWrapper) GetByID(reqID string) (models.ChangePassRequest, error) {
	var req models.ChangePassRequest
	err := c.parent.db.Select().Where(dbx.HashExp{"id": reqID}).From(models.ChangePasswordRequestsTableName).One(&req)
	return req, err
}
