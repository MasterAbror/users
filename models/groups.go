package models

import "time"

type Group struct {
	ID        string    `gorm:"size:100;primary_key;not null" json:"id"`
	Name      string    `gorm:"size:250;" json:"name"`
	Acronym   string    `gorm:"size:100;" json:"acronym"`
	Parent    string    `gorm:"size:100;" json:"parent"`
	Icon      string    `gorm:"size:100;" json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `gorm:"size:100" json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `gorm:"size:100" json:"updated_by"`
}

func (group *Group) TableName() string {
	return "za_groups"
}

type GroupCreate struct {
	Name    string `form:"name"`
	Acronym string `form:"acronym"`
	Parent  string `form:"parent"`
	Icon    string `form:"icon"`
}

type GroupUpdate struct {
	ID      string `form:"id"`
	Name    string `form:"name"`
	Acronym string `form:"acronym"`
	Parent  string `form:"parent"`
	Icon    string `form:"icon"`
}

func (group *Group) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = group.ID
	resp["name"] = group.Name
	resp["acronym"] = group.Acronym
	resp["parent"] = group.Parent
	resp["icon"] = group.Icon
	resp["created_at"] = group.CreatedAt
	resp["created_by"] = group.CreatedBy
	resp["updated_at"] = group.UpdatedAt
	resp["updated_by"] = group.UpdatedBy
	return resp
}
