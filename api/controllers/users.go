package controllers

import (
	"amaranth/api/models"
	"amaranth/api/services"
	"amaranth/api/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
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
		utils.RespondError(w, *saveErr)
		return
	}

	utils.RespondJson(w, http.StatusCreated, result)

}

func (c *usersController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, userErr := primitive.ObjectIDFromHex(vars["user_id"])
	if userErr != nil {
		err := utils.NewBadRequestError("not a valid ObjectID")
		utils.RespondError(w, *err)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		utils.RespondError(w, *getErr)
	}
	utils.RespondJson(w, http.StatusOK, user)
}

func (c *usersController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, userErr := primitive.ObjectIDFromHex(vars["user_id"])
	if userErr != nil {
		err := utils.NewBadRequestError("not a valid ObjectID")
		utils.RespondError(w, *err)
		return
	}

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
	user.Id = userId

	isPartial := r.Method == http.MethodPatch
	result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		utils.RespondError(w, *updateErr)
		return
	}

	utils.RespondJson(w, http.StatusOK, result)
}
