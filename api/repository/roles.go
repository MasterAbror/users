package repository

import (
	"html"

	"github.com/MasterAbror/users/database"
	"github.com/MasterAbror/users/models"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

type RoleRepository struct {
	db database.Database
}

func NewRoleRepository(db database.Database) RoleRepository {
	return RoleRepository{
		db: db,
	}
}

func (u RoleRepository) CreateRole(role models.RoleCreate) error {
	esc := bluemonday.UGCPolicy()
	ID := uuid.Must(uuid.NewRandom()).String()
	var dbRole models.Role
	dbRole.ID = ID
	dbRole.Name = html.EscapeString(esc.Sanitize(role.Name))
	return u.db.DB.Create(&dbRole).Error
}
