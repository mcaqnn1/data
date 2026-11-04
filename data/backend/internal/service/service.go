package service

import (
	"context"

	"github.com/deev/data/backend/internal/repository"
)

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{repo: repo}
}

type LoginData struct {
	ID       int   `json:"id"`
	UName    string `json:"user_name"`
	Password string `json:"password"`
}

func (s *Service) Login(ctx context.Context,d LoginData) error {
	return s.repo.Login(ctx,repository.LoginData(d))
}