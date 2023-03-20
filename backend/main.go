package main

import (
	// "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/tgrangeo/matcha/database"
	"github.com/tgrangeo/matcha/models"
        "github.com/tgrangeo/matcha/handler"
)

func main() {
	db := database.ConnectDb()
	// defer database.DropUsers(db)
	defer db.Close()
	database.CreateTable(db)
	//database.GetUsers(db)

        
	router := mux.NewRouter()
        //read
	router.HandleFunc("/api/v1/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handler.GetUsersById).Methods("GET")
        //create
	router.HandleFunc("/api/v1/users", handler.CreateUser).Methods("POST")
        //update
	router.HandleFunc("/api/v1/users/{id}", handler.UpdateUser).Methods("PUT")
        //delete
        router.HandleFunc("/api/v1/users", handler.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/api/v1/users/{id}", handler.DeleteUserById).Methods("DELETE")
        http.ListenAndServe(":8080", router)
}
