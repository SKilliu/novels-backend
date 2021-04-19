package db

import (
	"fmt"

	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

const getListWithParamQuery = `SELECT * FROM novels WHERE data LIKE %s ORDER BY %s %s OFFSET %d LIMIT %d;`

// UsersQ query interface, which provide access to DB functions.
type NovelsQ interface {
	Insert(novel models.Novel) error
	Update(novel models.Novel) error
	Delete(novel models.Novel) error
	GetByID(nid string) (models.Novel, error)
	GetListWithParam(param, orderColumn, order string, offset, limit int) ([]models.Novel, error)
}

// UsersWrapper wraps interface.
type NovelsWrapper struct {
	parent *DB
}

// UsersQ query interface getter.
func (d *DB) NovelsQ() NovelsQ {
	return &NovelsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (n *NovelsWrapper) Insert(novel models.Novel) error {
	return n.parent.db.Model(&novel).Insert()
}

func (n *NovelsWrapper) Update(novel models.Novel) error {
	return n.parent.db.Model(&novel).Update()
}

func (n *NovelsWrapper) Delete(novel models.Novel) error {
	return n.parent.db.Model(&novel).Delete()
}

func (n *NovelsWrapper) GetByID(nid string) (models.Novel, error) {
	var novel models.Novel
	err := n.parent.db.Select().Where(dbx.HashExp{"id": nid}).From(models.NovelsTableName).One(&novel)
	return novel, err
}

func (n *NovelsWrapper) GetListWithParam(param, orderColumn, order string, offset, limit int) ([]models.Novel, error) {
	var res []models.Novel
	err := n.parent.db.NewQuery(fmt.Sprintf(getListWithParamQuery, param, orderColumn, order, offset, limit)).All(&res)
	return res, err
}
