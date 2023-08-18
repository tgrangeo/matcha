package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tgrangeo/matcha/database"
	"github.com/tgrangeo/matcha/models"
	"github.com/tgrangeo/matcha/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if !CheckToken(w, r) {
		return
	}
	db := database.ConnectDb()
	users := database.GetUsers(db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {
	if !CheckToken(w, r) {
		return
	}
	//TODO: secure if id doesn t exist
	params := mux.Vars(r)
	tofind, _ := strconv.Atoi(params["id"])
	db := database.ConnectDb()
	user := database.GetUsersById(db, tofind)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetWhere(w http.ResponseWriter, r *http.Request) {
	if !CheckToken(w, r) {
		return
	}
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

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("New User Detected")

	// Define a variable to hold the JSON data
	var newUserInput models.NewUserInput

	// Decode the JSON data from the request body into newUserInput
	err := json.NewDecoder(r.Body).Decode(&newUserInput)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture des données JSON", http.StatusBadRequest)
		return
	}

	// Create the new user using the fields from newUserInput
	newUser := models.NewSubUser(nil, newUserInput.Firstname, newUserInput.Lastname, "", newUserInput.Birthdate, 0, 0, newUserInput.Type, newUserInput.Pokeball)
	log.Println(newUser)
	db := database.ConnectDb()
	database.InsertUser(db, *newUser)
	// Return the new user as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// if (!CheckToken(w,r)){
	// 	return
	// }
	decoder := json.NewDecoder(r.Body)
	db := database.ConnectDb()
	usr := models.User{}
	err := decoder.Decode(&usr)
	if err != nil {
		fmt.Println(err)
	}
	ret, str := utils.CheckUser(usr)
	fmt.Println(str)
	if ret == 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	database.InsertUser(db, usr)
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if !CheckToken(w, r) {
		return
	}
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
	if !CheckToken(w, r) {
		return
	}
	params := mux.Vars(r)
	tofind, _ := strconv.Atoi(params["id"])
	db := database.ConnectDb()
	database.DelUserById(db, tofind)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	if !CheckToken(w, r) {
		return
	}
	db := database.ConnectDb()
	database.DelUser(db)
}
