package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Listening and serving HTTP on port :8080")

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
