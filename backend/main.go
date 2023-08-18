package main

import (
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tgrangeo/matcha/database"

	// "github.com/tgrangeo/matcha/models"
	"github.com/joho/godotenv"
	"github.com/tgrangeo/matcha/handler"
)

func main() {
	godotenv.Load(".env")
	db := database.ConnectDb()
	defer db.Close()
	database.CreateTable(db)

	router := mux.NewRouter()

	//AUTH
	router.HandleFunc("/api/v1/signin", handler.SignIn).Methods("POST")
	router.HandleFunc("/api/v1/logout", handler.Logout).Methods("POST")
	router.HandleFunc("/api/v1/refresh", handler.Refresh).Methods("POST")

	//API USER
	router.HandleFunc("/api/v1/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handler.GetUsersById).Methods("GET")
	router.HandleFunc("/api/v1/users/{where}/{value}", handler.GetWhere).Methods("GET")
	// router.HandleFunc("/api/v1/users/?", handler.GetWhere).Methods("GET") URL PARAMS
	router.HandleFunc("/api/v1/users", handler.CreateNewUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users", handler.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/api/v1/users/{id}", handler.DeleteUsersById).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
