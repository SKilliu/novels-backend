package db

import (
	"fmt"

	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

var (
	getCompetitionsListWithParamQuery = `SELECT * FROM novels WHERE data LIKE %s ORDER BY %s %s OFFSET %d LIMIT %d;`
	getOpponentQuery                  = "SELECT * FROM users WHERE id = (SELECT user_id FROM novels WHERE id IN (SELECT novel_one_id FROM competitions WHERE status = 'waiting_for_opponent') LIMIT 1) LIMIT 1;"
)

// UsersQ query interface, which provide access to DB functions.
type CompetitionsQ interface {
	Insert(competition models.Competition) error
	Update(competition models.Competition) error
	Delete(competition models.Competition) error
	GetByID(nid string) (models.Competition, error)
	GetOpponentForStory(uid string) (models.Competition, error)
	GetListWithParam(param, orderColumn, order string, offset, limit int) ([]models.Competition, error)
}

// UsersWrapper wraps interface.
type CompetitionsWrapper struct {
	parent *DB
}

// UsersQ query interface getter.
func (d *DB) CompetitionsQ() CompetitionsQ {
	return &CompetitionsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (c *CompetitionsWrapper) Insert(competition models.Competition) error {
	return c.parent.db.Model(&competition).Insert()
}

func (c *CompetitionsWrapper) Update(competition models.Competition) error {
	return c.parent.db.Model(&competition).Update()
}

func (c *CompetitionsWrapper) Delete(competition models.Competition) error {
	return c.parent.db.Model(&competition).Delete()
}

func (c *CompetitionsWrapper) GetByID(nid string) (models.Competition, error) {
	var novel models.Competition
	err := c.parent.db.Select().Where(dbx.HashExp{"id": nid}).From(models.CompetitionsTableName).One(&novel)
	return novel, err
}

func (c *CompetitionsWrapper) GetListWithParam(param, orderColumn, order string, offset, limit int) ([]models.Competition, error) {
	var res []models.Competition
	err := c.parent.db.NewQuery(fmt.Sprintf(getCompetitionsListWithParamQuery, param, orderColumn, order, offset, limit)).All(&res)
	return res, err
}

func (c *CompetitionsWrapper) GetOpponentForStory(uid string) (models.Competition, error) {
	var res models.Competition
	return res, nil
}

func (c *CompetitionsWrapper) GetWaitingForOpponentByUserID(uid string) (models.Competition, error) {
	var res models.Competition
	err := c.parent.db.Select().From(models.CompetitionsTableName).Where(dbx.HashExp{""})
}
