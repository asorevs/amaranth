package controllers

import (
	"amaranth/api/models"
	"amaranth/api/services"
	"amaranth/api/utils"
	"encoding/json"
	"io"
	"net/http"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type usersController struct{}

func (c *usersController) Create(w http.ResponseWriter, r *http.Request) {

	var user models.User
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		err := utils.NewBadRequestError("invalid json body")
		utils.RespondError(w, *err)
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		utils.RespondError(w, *restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		saveErr := utils.NewBadRequestError("error saving data")
		utils.RespondError(w, *saveErr)
		return
	}

	utils.RespondJson(w, http.StatusCreated, result)

}

func (c *usersController) Get(w http.ResponseWriter, r *http.Request) {
	return
}
