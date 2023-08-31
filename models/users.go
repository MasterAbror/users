package models

import "time"

// User -> User struct to save user on database
type User struct {
	ID        string    `gorm:"size:100;primary_key;not null" json:"id"`
	Fullname  string    `gorm:"size:250;" json:"fullname"`
	Nickname  string    `gorm:"size:50;" json:"nickname"`
	NIK       string    `gorm:"size:20;" json:"nik"`
	Phone     string    `gorm:"size:20;" json:"phone"`
	Mobile    string    `gorm:"size:20;" json:"mobile"`
	Email     string    `gorm:"size:250;" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Photo     string    `gorm:"size:100" json:"photo"`
	IsActive  bool      `json:"is_active"`
	IsBanned  bool      `json:"is_banned"`
	Token     string    `gorm:"size:250;not null" json:"token"`
	RoleID    string    `gorm:"size:100;not null" json:"role_id"`
	GroupID   string    `gorm:"size:100;not null" json:"group_id"`
	LevelID   int       `gorm:"size:11;not null" json:"level_id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `gorm:"size:100" json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `gorm:"size:100" json:"updated_by"`
	Role      Role      `gorm:"Foreignkey:RoleID;association_foreignkey:ID;" json:"role"`
	Group     Group     `gorm:"Foreignkey:GroupID;association_foreignkey:ID;" json:"group"`
	Level     Level     `gorm:"Foreignkey:LevelID;association_foreignkey:ID;" json:"level"`
}

type UserRedis struct {
	ID       string `json:"id"`
	Token    string `json:"token"`
	Fullname string `json:"fullname"`
	GroupID  string `json:"group_id"`
	LevelID  string `json:"level_id"`
	RoleID   string `json:"role_id"`
}

// TableName -> returns the table name of User Model
func (user *User) TableName() string {
	return "za_users"
}

// UserLogin -> Request Binding for User Login
type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// UserRegister -> Request Binding for User Register
type UserRegister struct {
	Fullname   string `form:"fullname"`
	Email      string `form:"email" json:"email" binding:"required"`
	NIK        string `form:"nik" json:"nik"`
	Mobile     string `form:"mobile" json:"mobile"`
	RoleID     string `form:"role_id" json:"role_id"`
	GroupID    string `form:"group_id" json:"group_id"`
	LevelID    int    `form:"level_id" json:"level_id"`
	Password   string `form:"password" json:"password" binding:"required"`
	RePassword string `form:"repassword" json:"repassword" binding:"required"`
}

// ResponseMap -> response map method of User
func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["email"] = user.Email
	resp["fullname"] = user.Fullname
	resp["nickname"] = user.Nickname
	resp["is_active"] = user.IsActive
	resp["created_at"] = user.CreatedAt
	resp["updated_at"] = user.UpdatedAt
	return resp
}
