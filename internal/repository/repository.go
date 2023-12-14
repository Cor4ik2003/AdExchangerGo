package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
}

type Authorization interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return nil
}
