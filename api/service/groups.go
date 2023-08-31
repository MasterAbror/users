package service

import (
	"github.com/MasterAbror/users/api/repository"
	"github.com/MasterAbror/users/models"
)

type GroupService struct {
	repo repository.GroupRepository
}

func NewGroupService(repo repository.GroupRepository) GroupService {
	return GroupService{
		repo: repo,
	}
}

func (u GroupService) CreateValidation(group models.GroupCreate) any {
	if group.Name == "" {
		var msg = "Nama wajib diisi!"
		return msg
	}
	if group.Acronym == "" {
		var msg = "Nama pendek/Kode group wajib diisi!"
		return msg
	}
	return nil
}

func (u GroupService) UpdateValidation(group models.GroupUpdate) any {
	if group.Name == "" {
		var msg = "Nama wajib diisi!"
		return msg
	}
	if group.Acronym == "" {
		var msg = "Nama pendek/Kode group wajib diisi!"
		return msg
	}
	return nil
}

func (u GroupService) FindAll(group models.Group, keyword string) (*[]models.Group, int64, error) {
	return u.repo.FindAll(group, keyword)
}

func (u GroupService) CreateGroup(group models.GroupCreate, user models.UserRedis) error {
	return u.repo.CreateGroup(group, user)
}

func (u GroupService) ReadGroup(group models.Group) (models.Group, error) {
	return u.repo.ReadGroup(group)
}

func (u GroupService) UpdateGroup(group models.Group, user models.UserRedis) error {
	return u.repo.UpdateGroup(group, user)
}

func (u GroupService) DeleteGroup(group models.Group) error {
	return u.repo.DeleteGroup(group)
}
