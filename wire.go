package main

import (
	"database/sql"
	"learn-wire/handler"
	"learn-wire/repository"
	"learn-wire/service"

	"github.com/google/wire"
)

func Initialize(db *sql.DB) (*handler.MyHandler, error) {
	wire.Build(
		repository.NewMyRepository,
		service.NewMyService,
		handler.NewMyHandler,
	)

	return &handler.MyHandler{}, nil
}
