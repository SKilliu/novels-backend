package db

import (
	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type ResetPassRequestsQ interface {
	Insert(req models.ResetPassRequest) error
	Delete(req models.ResetPassRequest) error
	GetByID(reqID string) (models.ResetPassRequest, error)
}

type ResetPassRequestsWrapper struct {
	parent *DB
}

func (d *DB) ResetPassRequestsQ() ResetPassRequestsQ {
	return &ResetPassRequestsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (c *ResetPassRequestsWrapper) Insert(req models.ResetPassRequest) error {
	return c.parent.db.Model(&req).Insert()
}

func (c *ResetPassRequestsWrapper) Delete(req models.ResetPassRequest) error {
	return c.parent.db.Model(&req).Delete()
}

func (c *ResetPassRequestsWrapper) GetByID(reqID string) (models.ResetPassRequest, error) {
	var req models.ResetPassRequest
	err := c.parent.db.Select().Where(dbx.HashExp{"id": reqID}).From(models.ResetPasswordRequestsTableName).One(&req)
	return req, err
}
