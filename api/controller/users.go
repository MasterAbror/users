package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MasterAbror/users/api/service"
	"github.com/MasterAbror/users/models"
	"github.com/MasterAbror/users/util"

	"github.com/gin-gonic/gin"
)

// UserController struct
type UserController struct {
	service service.UserService
}

// NewUserController : NewUserController
func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

// CreateUser ->  calls CreateUser services for validated user
func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Body Provided")
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}
	if err := u.service.CreateValidation(user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	hashPassword, _ := util.HashPassword(user.Password)
	user.Password = hashPassword
	save := u.service.CreateUser(user)
	if save != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Gagal menyimpan data!")
		return
	}
	util.SuccessJSON(c, http.StatusOK, "Buat akun berhasil")
}

// LoginUser : Generates JWT Token for validated user
func (u *UserController) LoginUser(c *gin.Context) {
	var user models.UserLogin

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Body Provided")
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}
	if err := u.service.AuthValidation(user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Kredensial akun tidak dikenal!")
		return
	}
	var userData models.User
	userData.ID = dbUser.ID
	results, err := u.service.Find(userData)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Data tidak ditemukan!")
		return
	}
	ttl, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRED_IN"))
	token, err := util.CreateToken(ttl, results.ID, os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Token gagal digenerate!")
		return
	}
	rau := `{
		"token": "` + token + `",
		"id": "` + results.ID + `",
		"fullname": "` + results.Fullname + `",
		"group_id": "` + results.GroupID + `",
		"level_id": "` + strconv.Itoa(results.LevelID) + `",
		"role_id": "` + results.RoleID + `"
	}`
	if er := util.SetRedis(results.ID, rau); er != "" {
		util.ErrorJSON(c, http.StatusBadRequest, er)
		return
	}

	response := &util.Response{
		Status: true,
		Msg:    "Token berhasil dibuat",
		Data:   token,
	}
	c.JSON(http.StatusOK, response)
}

func (u *UserController) Authorization(r *gin.Context) {
	var access_token string

	authorizationHeader := r.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		access_token = fields[1]
	} else {
		access_token = ""
	}

	if access_token == "" {
		util.ErrorJSON(r, http.StatusUnauthorized, "Anda belum login!")
		return
	}

	id, err := util.ValidateToken(access_token, os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"))
	if err != nil {
		util.ErrorJSON(r, http.StatusUnauthorized, "Sesi Anda telah habis! Silahkan login kembali!")
		return
	}

	key := id.(string)
	raw := util.GetRedis(key)
	if raw == "err" {
		util.ErrorJSON(r, http.StatusUnauthorized, "Anda belum login!")
		return
	}

	user := raw.(models.UserRedis)
	if user.Token != access_token {
		util.ErrorJSON(r, http.StatusUnauthorized, "Sesi Anda telah habis! Silahkan login kembali.")
		return
	}

	r.Set("currentUser", user)
	r.Next()
}
