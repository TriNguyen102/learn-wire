package repository

import (
	"database/sql"
)

func NewMyRepository(db *sql.DB) *MyRepository {
	return &MyRepository{db: db}
}

type MyRepository struct {
	db *sql.DB
}

func (mr *MyRepository) CheckConnection() error {
	return nil
}

func (mr *MyRepository) IsExisted() bool {
	return true
}
