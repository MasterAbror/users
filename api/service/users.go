package service

import (
	"github.com/MasterAbror/users/api/repository"
	"github.com/MasterAbror/users/models"
)

// UserService UserService struct
type UserService struct {
	repo repository.UserRepository
}

// NewUserService : get injected user repo
func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

// Save -> saves users entity
func (u UserService) CreateValidation(user models.UserRegister) any {
	if user.Fullname == "" {
		var msg = "Nama lengkap wajib diisi!"
		return msg
	}
	if user.Email == "" {
		var msg = "Email wajib diisi!"
		return msg
	}
	if user.NIK == "" {
		var msg = "NIK wajib diisi!"
		return msg
	}
	if user.Mobile == "" {
		var msg = "Nomor Ponsel wajib diisi!"
		return msg
	}
	if user.RoleID == "" {
		var msg = "Peran wajib diisi!"
		return msg
	}
	if user.GroupID == "" {
		var msg = "Group wajib diisi!"
		return msg
	}
	if user.LevelID == 0 {
		var msg = "Level wajib diisi!"
		return msg
	}
	if user.Password == "" {
		var msg = "Kata sandi wajib diisi!"
		return msg
	}
	if user.RePassword == "" {
		var msg = "Konfirmasi Kata sandi wajib diisi!"
		return msg
	}
	if user.Password != user.RePassword {
		var msg = "Konfirmasi Kata sandi tidak sama!"
		return msg
	}
	var email models.User
	email.Email = user.Email
	foundEmail, err := u.repo.Find(email)
	if err == nil {
		var msg = "Email sudah terdaftar!"
		foundEmail.ResponseMap()
		return msg
	}
	var nik models.User
	nik.NIK = user.NIK
	foundNIK, err := u.repo.Find(nik)
	if err == nil {
		var msg = "NIK sudah terdaftar!"
		foundNIK.ResponseMap()
		return msg
	}
	var mobile models.User
	mobile.Mobile = user.Mobile
	foundMobile, err := u.repo.Find(mobile)
	if err == nil {
		var msg = "Nomor ponsel sudah terdaftar!"
		foundMobile.ResponseMap()
		return msg
	}
	return nil
}

func (u UserService) AuthValidation(user models.UserLogin) any {
	if user.Email == "" {
		var msg = "Email wajib diisi!"
		return msg
	}
	if user.Password == "" {
		var msg = "Kata sandi wajib diisi!"
		return msg
	}
	var email models.User
	email.Email = user.Email
	foundEmail, err := u.repo.Find(email)
	if err != nil {
		var msg = "Email tidak terdaftar!"
		foundEmail.ResponseMap()
		return msg
	}
	return nil
}

func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

// Login -> Gets validated user
func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}

// Login -> Gets validated user
func (u UserService) Find(user models.User) (models.User, error) {
	return u.repo.Find(user)
}

func (u UserService) FindByID(user models.User) (*models.User, error) {
	return u.repo.FindByID(user)
}
