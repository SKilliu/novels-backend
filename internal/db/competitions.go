package db

import (
	"fmt"

	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

var (
	getCompetitionsListWithParamQuery = `SELECT * FROM novels_pool WHERE (status LIKE %s AND user_one_id = '%s') OR (status LIKE %s AND user_two_id = '%s') ORDER BY %s %s OFFSET %d LIMIT %d;`
	getOpponentQuery                  = "SELECT * FROM novels_pool WHERE novel_one_id IN (SELECT id FROM novels WHERE user_id IN (SELECT id FROM users WHERE rate = %d AND NOT id = '%s')) AND status = 'waiting_for_opponent' LIMIT 1;"
)

// UsersQ query interface, which provide access to DB functions.
type CompetitionsQ interface {
	Insert(competition models.Competition) error
	Update(competition models.Competition) error
	Delete(competition models.Competition) error
	GetByID(nid string) (models.Competition, error)
	GetCompetitionOpponent(userRate int, userID string) (models.Competition, error)
	GetByNovelOneID(nid string) (models.Competition, error)
	GetByNovelID(nid string) (models.Competition, error)
	GetListWithParam(param, uid, orderColumn, order string, offset, limit int) ([]models.Competition, error)
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

func (c *CompetitionsWrapper) GetListWithParam(param, uid, orderColumn, order string, offset, limit int) ([]models.Competition, error) {
	var res []models.Competition
	err := c.parent.db.NewQuery(fmt.Sprintf(getCompetitionsListWithParamQuery, param, uid, param, uid, orderColumn, order, offset, limit)).All(&res)
	return res, err
}

func (c *CompetitionsWrapper) GetCompetitionOpponent(userRate int, userID string) (models.Competition, error) {
	var res models.Competition
	err := c.parent.db.NewQuery(fmt.Sprintf(getOpponentQuery, userRate, userID)).One(&res)
	return res, err
}

func (c *CompetitionsWrapper) GetByNovelOneID(nid string) (models.Competition, error) {
	var res models.Competition
	err := c.parent.db.Select().From(models.CompetitionsTableName).Where(dbx.HashExp{"novel_one_id": nid}).One(&res)
	return res, err
}

func (c *CompetitionsWrapper) GetByNovelID(nid string) (models.Competition, error) {
	var res models.Competition
	err := c.parent.db.Select().From(models.CompetitionsTableName).Where(dbx.HashExp{"novel_one_id": nid}).OrWhere(dbx.HashExp{"novel_two_id": nid}).One(&res)
	return res, err
}
