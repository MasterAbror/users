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

type GroupController struct {
	service service.GroupService
	user    service.UserService
}

func NewGroupController(s service.GroupService) GroupController {
	return GroupController{
		service: s,
	}
}

func (u *GroupController) All(r *gin.Context) {
	currentUser := r.MustGet("currentUser").(models.UserRedis)
	var group models.Group

	keyword := r.Query("keyword")

	data, total, err := u.service.FindAll(group, keyword)

	if err != nil {
		util.ErrorJSON(r, http.StatusBadRequest, "Gagal menemukan data!")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}
	groupData := map[string]interface{}{
		"rows":       respArr,
		"total_rows": total,
	}

	token := util.CreateNewToken(currentUser.ID)
	if token == "" {
		util.ErrorJSON(r, http.StatusBadRequest, "Token gagal digenerate!")
		return
	}
	response := &util.Response{
		Status: true,
		Msg:    "Group berhasil dibuat",
		Data:   groupData,
		Token:  token,
	}
	r.JSON(http.StatusOK, response)
}

func (u *GroupController) CreateGroup(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.UserRedis)
	if currentUser.LevelID != "1" {
		util.ErrorJSON(c, http.StatusBadRequest, "Anda dilarang mengakses fitur ini!")
		return
	}
	var group models.GroupCreate
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Body Provided")
		return
	}
	err = json.Unmarshal(body, &group)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}
	if err := u.service.CreateValidation(group); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	save := u.service.CreateGroup(group, currentUser)
	if save != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Gagal menyimpan data!")
		return
	}
	token := util.CreateNewToken(currentUser.ID)
	if token == "" {
		util.ErrorJSON(c, http.StatusBadRequest, "Token gagal digenerate!")
		return
	}
	response := &util.Response{
		Status: true,
		Msg:    "Group berhasil dibuat",
		Data:   "",
		Token:  token,
	}
	c.JSON(http.StatusOK, response)
}

func (u *GroupController) ReadGroup(r *gin.Context) {
	currentUser := r.MustGet("currentUser").(models.UserRedis)
	var group models.Group
	id := r.Param("id")
	group.ID = id
	groupData, err := u.service.ReadGroup(group)
	if err != nil {
		util.ErrorJSON(r, http.StatusBadRequest, "Group tidak ditemukan!")
		return
	}
	data := map[string]interface{}{
		"id":         groupData.ID,
		"name":       groupData.Name,
		"acronym":    groupData.Acronym,
		"parent":     groupData.Parent,
		"created_by": groupData.CreatedBy,
		"created_at": groupData.CreatedAt,
		"updated_by": groupData.UpdatedBy,
		"updated_at": groupData.UpdatedAt,
	}
	token := util.CreateNewToken(currentUser.ID)
	if token == "" {
		util.ErrorJSON(r, http.StatusBadRequest, "Token gagal digenerate!")
		return
	}
	response := &util.Response{
		Status: true,
		Msg:    "Detail Group",
		Data:   data,
		Token:  token,
	}
	r.JSON(http.StatusOK, response)
}

func (u *GroupController) UpdateGroup(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.UserRedis)
	if currentUser.LevelID != "1" {
		util.ErrorJSON(c, http.StatusBadRequest, "Anda dilarang mengakses fitur ini!")
		return
	}
	var update models.GroupUpdate
	var group models.Group
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Body Provided")
		return
	}
	err = json.Unmarshal(body, &update)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}
	if err := u.service.UpdateValidation(update); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	group.ID = update.ID
	group.Name = update.Name
	group.Acronym = update.Acronym
	group.Parent = update.Parent
	save := u.service.UpdateGroup(group, currentUser)
	if save != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Gagal menyimpan data!")
		return
	}
	token := util.CreateNewToken(currentUser.ID)
	response := &util.Response{
		Status: true,
		Msg:    "Group berhasil diubah",
		Data:   "",
		Token:  token,
	}
	c.JSON(http.StatusOK, response)
}

func (u *GroupController) DeleteGroup(r *gin.Context) {
	currentUser := r.MustGet("currentUser").(models.UserRedis)
	if currentUser.LevelID != "1" {
		util.ErrorJSON(r, http.StatusBadRequest, "Anda dilarang mengakses fitur ini!")
		return
	}
	var group models.Group
	id := r.Param("id")
	group.ID = id
	if err := u.service.DeleteGroup(group); err != nil {
		util.ErrorJSON(r, http.StatusBadRequest, "Hapus group gagal!")
		return
	}
	token := util.CreateNewToken(currentUser.ID)
	if token == "" {
		util.ErrorJSON(r, http.StatusBadRequest, "Token gagal digenerate!")
		return
	}
	response := &util.Response{
		Status: true,
		Msg:    "Group berhasil dihapus",
		Data:   "",
		Token:  token,
	}
	r.JSON(http.StatusOK, response)
}
