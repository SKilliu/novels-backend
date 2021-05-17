package db

import (
	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// UsersQ query interface, which provide access to DB functions.
type ReadyForVoteQ interface {
	Insert(rfv models.ReadyForVote) error
	Update(rfv models.ReadyForVote) error
	DeleteByNovelID(novelID string) error
	GetByUserAndCompetitionIDs(uid, cid string) (models.ReadyForVote, error)
	GetForVote() (models.ReadyForVote, error)
}

// UsersWrapper wraps interface.
type ReadyForVoteWrapper struct {
	parent *DB
}

// UsersQ query interface getter.
func (d *DB) ReadyForVoteQ() ReadyForVoteQ {
	return &ReadyForVoteWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (r *ReadyForVoteWrapper) Insert(rfv models.ReadyForVote) error {
	return r.parent.db.Model(&rfv).Insert()
}

func (r *ReadyForVoteWrapper) Update(rfv models.ReadyForVote) error {
	return r.parent.db.Model(&rfv).Update()
}

func (r *ReadyForVoteWrapper) GetByUserAndCompetitionIDs(uid, cid string) (models.ReadyForVote, error) {
	var res models.ReadyForVote
	err := r.parent.db.Select().From(models.ReadyForVoteTableName).Where(dbx.HashExp{"user_id": uid, "novels_pool_id": cid}).One(&res)
	return res, err
}

func (c *ReadyForVoteWrapper) GetForVote() (models.ReadyForVote, error) {
	var res models.ReadyForVote
	err := c.parent.db.Select().From(models.ReadyForVoteTableName).Where(dbx.HashExp{"is_voted": false}).OrderBy("views_amount ASC").Limit(1).One(&res)
	return res, err
}

func (c *ReadyForVoteWrapper) DeleteByNovelID(novelID string) error {
	_, err := c.parent.db.Delete(models.ReadyForVoteTableName, dbx.HashExp{"novel_id": novelID}).Execute()
	return err
}
