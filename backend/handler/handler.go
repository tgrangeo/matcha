package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/tgrangeo/matcha/database"
	"github.com/tgrangeo/matcha/models"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDb()
	users := database.GetUsers(db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {
	//TODO: secure if id doesn t exist
	params := mux.Vars(r)
	tofind, _ := strconv.Atoi(params["id"])
	db := database.ConnectDb()
	user := database.GetUsersById(db, tofind)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetWhere(w http.ResponseWriter, r *http.Request){
	//TODO:where exist ???
	params := mux.Vars(r)
	tofind := params["where"]
	value := params["value"]
	fmt.Println(tofind, value)
	db := database.ConnectDb()
	users := database.GetUsersWhere(db, tofind, value)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	db := database.ConnectDb()
	usr := models.User{}
    err := decoder.Decode(&usr)
	if err != nil {
		fmt.Println(err)
	}
	database.InsertUser(db, usr)
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tofind, _ := strconv.Atoi(params["id"])
	decoder := json.NewDecoder(r.Body)
	db := database.ConnectDb()
	usr := models.User{}
    err := decoder.Decode(&usr)
	if err != nil {
		fmt.Println(err)
	}
	database.UpdateUser(db, usr, tofind)
	w.WriteHeader(http.StatusOK)
}

func DeleteUsersById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tofind, _ := strconv.Atoi(params["id"])
	db := database.ConnectDb()
	database.DelUserById(db, tofind)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDb()
	database.DelUser(db)
}
