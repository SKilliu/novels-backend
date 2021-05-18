package db

import (
	"fmt"

	"github.com/SKilliu/novels-backend/internal/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

var getReadyForVoteQuery = `SELECT * FROM ready_for_vote WHERE novels_pool_id = (SELECT id FROM novels_pool WHERE id IN (SELECT novels_pool_id FROM ready_for_vote WHERE user_id = '%s' AND is_viewed = 'false') ORDER BY views_amount ASC LIMIT 1) AND user_id = '%s' LIMIT 1;`

// UsersQ query interface, which provide access to DB functions.
type ReadyForVoteQ interface {
	Insert(rfv models.ReadyForVote) error
	Update(rfv models.ReadyForVote) error
	Delete(rfv models.ReadyForVote) error
	GetByUserAndCompetitionIDs(uid, cid string) (models.ReadyForVote, error)
	GetForVote(userid string) (models.ReadyForVote, error)
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

func (c *ReadyForVoteWrapper) GetForVote(userid string) (models.ReadyForVote, error) {
	var res models.ReadyForVote
	err := c.parent.db.NewQuery(fmt.Sprintf(getReadyForVoteQuery, userid, userid)).One(&res)
	return res, err
}

func (r *ReadyForVoteWrapper) Delete(rfv models.ReadyForVote) error {
	return r.parent.db.Model(&rfv).Delete()
}
