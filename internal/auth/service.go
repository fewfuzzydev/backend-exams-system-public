package auth

import "exams/internal/modules/users"

type Service interface {
	Login(CheckUserLogin) (*users.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Login(req CheckUserLogin) (*users.User, error) {
	return s.repo.Login(req)
}
