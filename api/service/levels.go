package service

import (
	"github.com/MasterAbror/users/api/repository"
	"github.com/MasterAbror/users/models"
)

type LevelService struct {
	repo repository.LevelRepository
}

func NewLevelService(repo repository.LevelRepository) LevelService {
	return LevelService{
		repo: repo,
	}
}

func (u LevelService) CreateValidation(level models.LevelCreate) any {
	if level.Name == "" {
		var msg = "Nama wajib diisi!"
		return msg
	}
	return nil
}

func (u LevelService) CreateLevel(level models.LevelCreate) error {
	return u.repo.CreateLevel(level)
}
