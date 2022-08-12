package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mux-server/user"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var err error

func main() {
	// load file .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env file ")
	}

	zap.L().Info("ENV Load Completes.")

	//creating a mux router
	router := mux.NewRouter()

	//creating multiple handlers for all requests
	router.HandleFunc("/user/details", user.GetUserDetails).Methods(http.MethodGet)
	router.HandleFunc("/user/service/{service}", user.PostService).Methods(http.MethodPost)

	//starting the server
	zap.L().Info(fmt.Sprintf("Listening & Serving on : %s", os.Getenv("APP_PORT")))

	err = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), router)
	if err != nil {
		zap.L().Error("Listening & Serving, error :", zap.Error(err))
		return
	}
}
