package repository

import (
	"html"
	"time"

	"github.com/MasterAbror/users/database"
	"github.com/MasterAbror/users/models"
	"github.com/MasterAbror/users/util"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

// UserRepository -> UserRepository resposible for accessing database
type UserRepository struct {
	db database.Database
}

// NewUserRepository -> creates a instance on UserRepository
func NewUserRepository(db database.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// CreateUser -> method for saving user to database
func (u UserRepository) CreateUser(user models.UserRegister) error {
	esc := bluemonday.UGCPolicy()
	ID := uuid.Must(uuid.NewRandom()).String()
	CreatedAt := time.Now()
	var dbUser models.User
	dbUser.ID = ID
	dbUser.Fullname = html.EscapeString(esc.Sanitize(user.Fullname))
	dbUser.Email = html.EscapeString(esc.Sanitize(user.Email))
	dbUser.NIK = html.EscapeString(esc.Sanitize(user.NIK))
	dbUser.Mobile = html.EscapeString(esc.Sanitize(user.Mobile))
	dbUser.Password = html.EscapeString(esc.Sanitize(user.Password))
	dbUser.RoleID = html.EscapeString(esc.Sanitize(user.RoleID))
	dbUser.GroupID = html.EscapeString(esc.Sanitize(user.GroupID))
	dbUser.LevelID = user.LevelID
	dbUser.CreatedBy = ID
	dbUser.CreatedAt = CreatedAt
	dbUser.IsActive = true
	dbUser.IsBanned = false
	return u.db.DB.Create(&dbUser).Error
}

// LoginUser -> method for returning user
func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {

	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := util.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}

func (u UserRepository) Find(user models.User) (models.User, error) {
	var users models.User
	err := u.db.DB.
		Debug().
		Model(&models.User{}).
		Where(&user).
		Preload("Role").
		Preload("Group").
		Preload("Level").
		Take(&users).Error
	return users, err
}

func (u UserRepository) FindByID(user models.User) (*models.User, error) {
	var dbUser models.User
	ID := user.ID
	err := u.db.DB.Where("id = ?", ID).First(&dbUser).Error
	return &dbUser, err
}
