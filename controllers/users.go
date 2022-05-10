package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mahirB/user_managment_api/models"
	"github.com/mahirB/user_managment_api/services"
)

type Users struct {
	User models.User
}

func (u Users) List(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	queryValues := req.URL.Query()

	var filter models.UserFilter

	// Uncomment to start using serverside pagination
	/* 	filter.Limit, _ = strconv.Atoi(queryValues.Get("limit"))
	   	if filter.Limit == 0 {
	   		filter.Limit = 10
	   	}

	   	filter.Offset, _ = strconv.Atoi(queryValues.Get("offset"))
	   	if filter.Offset != 0 {
	   		filter.Offset = filter.Limit * filter.Offset
	   	}
	*/

	filter.Status = queryValues.Get("status")

	users, err := models.UsersList(filter)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	// Uncomment to start using serverside pagination
	/* 	count, err := models.UserCount(filter)
	   	if err != nil {
	   		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	   		return
	   	} */

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"users": users /* "count": count */})
}

func (u Users) Get(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	u.User.ID, _ = strconv.Atoi(params.ByName("id"))

	err := u.User.Get()
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"user": u.User})
}

func (u Users) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := json.NewDecoder(req.Body).Decode(&u.User)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	errMsg := checkUserData(u.User)
	if errMsg != "" {
		services.Renderer.Render(res, http.StatusNotAcceptable, map[string]interface{}{"error": errMsg})
		return
	}
	if len(u.User.Password) < 3 || len(u.User.Password) > 25 {
		services.Renderer.Render(res, http.StatusNotAcceptable, map[string]interface{}{"error": "Password length incompatible"})
		return
	}

	err = u.User.Create()
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"user": u.User})
}

func (u Users) Update(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := json.NewDecoder(req.Body).Decode(&u.User)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	errMsg := checkUserData(u.User)
	if errMsg != "" {
		services.Renderer.Render(res, http.StatusNotAcceptable, map[string]interface{}{"error": errMsg})
		return
	}

	err = u.User.Update()
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"user": u.User})
}

func (u Users) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	uid, _ := strconv.Atoi(params.ByName("id"))

	tx, err := models.CreateTransaction()
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		tx.Rollback()
		return
	}
	err = models.DeleteUserPermissions(uid, tx)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		tx.Rollback()
		return
	}

	err = models.UserDelete(uid, tx)
	if err != nil {
		services.Renderer.Render(res, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()

	services.Renderer.Render(res, http.StatusOK, map[string]interface{}{"message": "User with id: " + strconv.Itoa(uid) + " has been deleted"})
}

func checkUserData(user models.User) string {
	if len(user.FirstName) < 3 {
		return "Firstname length incompatible"
	}
	if len(user.LastName) < 3 {
		return "Lastname length incompatible"
	}

	if len(user.Email) < 5 {
		return "Email length incompatible"
	}
	if len(user.Username) < 5 {
		return "Username length incompatible"
	}

	if !(user.Status == 0 || user.Status == 1) {
		return "Bad status value"
	}

	return ""
}
