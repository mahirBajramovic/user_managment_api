package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mahirB/user_managment_api/models"
	"github.com/mahirB/user_managment_api/services"
)

type Permissions struct {
}

func (p Permissions) List(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	perms, err := models.PermissionsList()
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"permissions": perms})
}

func (p Permissions) Get(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID, _ := strconv.Atoi(params.ByName("id"))

	perms, err := models.UserPermissions(userID)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"user_permissions": perms})
}

func (p Permissions) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID, _ := strconv.Atoi(params.ByName("id"))
	var userPerms []models.UserPermission

	err := json.NewDecoder(req.Body).Decode(&userPerms)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	err = models.CreateUserPermissions(userID, userPerms)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"message": "User permissions edited"})
}
