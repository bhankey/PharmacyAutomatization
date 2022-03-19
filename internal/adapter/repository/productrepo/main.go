package productrepo

import "github.com/jmoiron/sqlx"

type Repository struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

func NewProductRepo(master *sqlx.DB, slave *sqlx.DB) *Repository {
	return &Repository{
		master: master,
		slave:  slave,
	}
}
