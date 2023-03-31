package service

import (
	"learn-wire/repository"
)

func NewMyService(repo *repository.MyRepository) *MyService {
	return &MyService{repo: repo}
}

type MyServiceInterface interface {
	SomeMethod() (bool, error)
}

type MyService struct {
	repo *repository.MyRepository
}

func (s *MyService) SomeMethod() (bool, error) {
	// Check existed database and Connect to DB return error
	return s.repo.IsExisted(), s.repo.CheckConnection()
}
