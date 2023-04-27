package controllers

import (
	"amaranth/api/models"
	"amaranth/api/services"
	"encoding/json"
	"fmt"
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
	fmt.Println("\n", "user", user)
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		//TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle error
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle error
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func (c *usersController) Get(w http.ResponseWriter, r *http.Request) {
	return
}
