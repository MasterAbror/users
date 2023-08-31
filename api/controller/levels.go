package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/MasterAbror/users/api/service"
	"github.com/MasterAbror/users/models"
	"github.com/MasterAbror/users/util"

	"github.com/gin-gonic/gin"
)

type LevelController struct {
	service service.LevelService
}

func NewLevelController(s service.LevelService) LevelController {
	return LevelController{
		service: s,
	}
}

func (u *LevelController) CreateLevel(c *gin.Context) {
	var level models.LevelCreate
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Body Provided")
		return
	}
	err = json.Unmarshal(body, &level)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}
	if err := u.service.CreateValidation(level); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	save := u.service.CreateLevel(level)
	if save != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Gagal menyimpan data!")
		return
	}
	util.SuccessJSON(c, http.StatusOK, "Buat level berhasil")
}
