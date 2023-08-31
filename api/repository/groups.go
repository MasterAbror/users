package repository

import (
	"html"
	"time"

	"github.com/MasterAbror/users/database"
	"github.com/MasterAbror/users/models"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

// UserRepository -> UserRepository resposible for accessing database
type GroupRepository struct {
	db database.Database
}

// NewGroupRepository -> creates a instance on GroupRepository
func NewGroupRepository(db database.Database) GroupRepository {
	return GroupRepository{
		db: db,
	}
}

func (u GroupRepository) FindAll(group models.Group, keyword string) (*[]models.Group, int64, error) {
	var groups []models.Group
	var totalRows int64 = 0

	queryBuider := u.db.DB.Order("name asc").Model(&models.Group{})

	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			u.db.DB.Where("groups.name LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(group).
		Find(&groups).
		Count(&totalRows).Error
	return &groups, totalRows, err
}

// CreateGroup -> method for saving Group to database
func (u GroupRepository) CreateGroup(group models.GroupCreate, user models.UserRedis) error {
	esc := bluemonday.UGCPolicy()
	CreatedAt := time.Now()
	ID := uuid.Must(uuid.NewRandom()).String()
	var dbGroup models.Group
	dbGroup.ID = ID
	dbGroup.Name = html.EscapeString(esc.Sanitize(group.Name))
	dbGroup.Acronym = html.EscapeString(esc.Sanitize(group.Acronym))
	dbGroup.Parent = html.EscapeString(esc.Sanitize(group.Parent))
	dbGroup.Icon = html.EscapeString(esc.Sanitize(group.Icon))
	dbGroup.CreatedBy = user.ID
	dbGroup.CreatedAt = CreatedAt
	return u.db.DB.Create(&dbGroup).Error
}

func (u GroupRepository) ReadGroup(group models.Group) (models.Group, error) {
	var groups models.Group
	err := u.db.DB.
		Debug().
		Model(&models.Group{}).
		Where(&group).
		Take(&groups).Error
	return groups, err
}

func (u GroupRepository) UpdateGroup(group models.Group, user models.UserRedis) error {
	esc := bluemonday.UGCPolicy()
	UpdatedAt := time.Now()
	var dbGroup models.Group
	dbGroup.ID = group.ID
	dbGroup.Name = html.EscapeString(esc.Sanitize(group.Name))
	dbGroup.Acronym = html.EscapeString(esc.Sanitize(group.Acronym))
	dbGroup.Parent = html.EscapeString(esc.Sanitize(group.Parent))
	dbGroup.Icon = html.EscapeString(esc.Sanitize(group.Icon))
	dbGroup.UpdatedBy = user.ID
	dbGroup.UpdatedAt = UpdatedAt
	return u.db.DB.Debug().Where("id = ?", dbGroup.ID).Take(&group).UpdateColumns(&dbGroup).Error
}

func (u GroupRepository) DeleteGroup(group models.Group) error {
	return u.db.DB.Delete(&group).Error
}
