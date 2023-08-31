package repository

import (
	"html"

	"github.com/MasterAbror/users/database"
	"github.com/MasterAbror/users/models"
	"github.com/microcosm-cc/bluemonday"
)

type LevelRepository struct {
	db database.Database
}

func NewLevelRepository(db database.Database) LevelRepository {
	return LevelRepository{
		db: db,
	}
}

func (u LevelRepository) CreateLevel(level models.LevelCreate) error {
	esc := bluemonday.UGCPolicy()
	var dbLevel models.Level
	dbLevel.Name = html.EscapeString(esc.Sanitize(level.Name))
	return u.db.DB.Create(&dbLevel).Error
}
