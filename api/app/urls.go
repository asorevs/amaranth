package app

import (
	"amaranth/api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", controllers.UsersController.Get).Methods(http.MethodGet)
	router.HandleFunc("/users", controllers.UsersController.Create).Methods(http.MethodPost)
	router.HandleFunc("/users/{user_id}", controllers.UsersController.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{user_id}", controllers.UsersController.Update).Methods(http.MethodPatch)

}
