package models

import "time"

type Role struct {
	ID        string    `gorm:"size:100;primary_key;not null" json:"id"`
	Name      string    `gorm:"size:250;" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `gorm:"size:100" json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `gorm:"size:100" json:"updated_by"`
}

func (role *Role) TableName() string {
	return "za_roles"
}

type RoleCreate struct {
	Name string `form:"name"`
}
