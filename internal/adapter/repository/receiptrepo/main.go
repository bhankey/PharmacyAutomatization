package receiptrepo

import "github.com/jmoiron/sqlx"

type Repository struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

func NewReceiptRepo(master *sqlx.DB, slave *sqlx.DB) *Repository {
	return &Repository{
		master: master,
		slave:  slave,
	}
}
