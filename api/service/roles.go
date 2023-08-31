package service

import (
	"github.com/MasterAbror/users/api/repository"
	"github.com/MasterAbror/users/models"
)

type RoleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return RoleService{
		repo: repo,
	}
}

func (u RoleService) CreateValidation(role models.RoleCreate) any {
	if role.Name == "" {
		var msg = "Nama wajib diisi!"
		return msg
	}
	return nil
}

func (u RoleService) CreateRole(role models.RoleCreate) error {
	return u.repo.CreateRole(role)
}
