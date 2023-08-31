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

type RoleController struct {
	service service.RoleService
}

func NewRoleController(s service.RoleService) RoleController {
	return RoleController{
		service: s,
	}
}

func (u *RoleController) CreateRole(c *gin.Context) {
	var role models.RoleCreate
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Body Provided")
		return
	}
	err = json.Unmarshal(body, &role)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}
	if err := u.service.CreateValidation(role); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	save := u.service.CreateRole(role)
	if save != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Gagal menyimpan data!")
		return
	}
	util.SuccessJSON(c, http.StatusOK, "Buat role berhasil")
}
