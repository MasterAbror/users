package models

import "time"

type Level struct {
	ID        int       `gorm:"primary_key;auto_increment;not null" json:"id"`
	Name      string    `gorm:"size:250;" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `gorm:"size:100" json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `gorm:"size:100" json:"updated_by"`
}

func (level *Level) TableName() string {
	return "za_levels"
}

type LevelCreate struct {
	Name string `form:"name"`
}
