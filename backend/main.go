package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/tgrangeo/matcha/database"
	"github.com/tgrangeo/matcha/models"
    "github.com/tgrangeo/matcha/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db := database.ConnectDb()
	// defer database.DropUser`s(db)
	defer db.Close()
	database.CreateTable(db)
	//database.GetUsers(db)

	fmt.Println(models.Tags[42])
        
	router := mux.NewRouter()
        //read
	router.HandleFunc("/api/v1/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handler.GetUsersById).Methods("GET")
	router.HandleFunc("/api/v1/users/{where}/{value}", handler.GetWhere).Methods("GET")
	// router.HandleFunc("/api/v1/users/?", handler.GetWhere).Methods("GET") URL PARAMS 
        //create
	router.HandleFunc("/api/v1/users", handler.CreateUser).Methods("POST")
        //update
	router.HandleFunc("/api/v1/users/{id}", handler.UpdateUser).Methods("PUT")
        //delete
        router.HandleFunc("/api/v1/users", handler.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/api/v1/users/{id}", handler.DeleteUsersById).Methods("DELETE")
        http.ListenAndServe(":8080", router)
}
