package userrepo

import "github.com/jmoiron/sqlx"

type UserRepo struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

func NewUserRepo(master *sqlx.DB, slave *sqlx.DB) *UserRepo {
	return &UserRepo{
		master: master,
		slave:  slave,
	}
}
