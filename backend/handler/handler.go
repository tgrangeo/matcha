package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tgrangeo/matcha/database"
	"github.com/tgrangeo/matcha/models"
	"github.com/tgrangeo/matcha/utils"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	claims, valid := CheckToken(w, r)
	if !valid {
		return
	}
	fmt.Println(claims.mail)
	db := database.ConnectDb()
	users := database.GetUsers(db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {
	claims, valid := CheckToken(w, r)
	if !valid {
		return
	}
	fmt.Println(claims.mail)
	//TODO: secure if id doesn t exist
	params := mux.Vars(r)
	tofind, _ := strconv.Atoi(params["id"])
	db := database.ConnectDb()
	user := database.GetUsersById(db, tofind)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetWhere(w http.ResponseWriter, r *http.Request) {
	claims, valid := CheckToken(w, r)
	if !valid {
		return
	}
	fmt.Println(claims.mail)
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

type PassRequest struct {
	Token string `json:"token"`
	Pass  string `json:"pass"`
}

func NewPass(w http.ResponseWriter, r *http.Request) {
	var req PassRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	db := database.ConnectDb()
	users := database.GetUsersWhere(db, "temp_token", req.Token)
	crypted, _ := utils.HashPassword(req.Pass)
	users[0].Pass = crypted
	database.UpdateUser(db, users[0])
	w.WriteHeader(http.StatusOK)
}

type TokenRequest struct {
	Token string `json:"token"`
}

func ResetPass(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	//create token
	tok, err := utils.RandomToken()
	if err != nil {
		panic("alert")
	}
	// token in db
	db := database.ConnectDb()
	user := database.GetUsersWhere(db, "email", req.Token)
	user[0].Token = tok
	database.UpdateUser(db, user[0])
	//send mail
	if err := utils.SendResetPassEmail(req.Token, tok); err != nil {
		http.Error(w, "Erreur lors de l'envoi de l'e-mail de confirmation", http.StatusInternalServerError)
		return
	}

}

func ConfirmRegistration(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	db := database.ConnectDb()
	users := database.GetUsersWhere(db, "email", "thomas.grangeon9@gmail.com")
	if users[0].Token == req.Token {
		fmt.Println("token ok")
		w.WriteHeader(http.StatusOK)

	} else {
		fmt.Println("token false")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("Inscription confirmée !"))
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

	//generer un token random pour la confirmation
	tok, err := utils.RandomToken()
	if err != nil {
		panic("alert")
	}

	// Create the new user using the fields from newUserInput
	log.Println("sub " + tok)
	newUser := models.NewSubUser(nil, newUserInput.Firstname, newUserInput.Lastname, newUserInput.Email, newUserInput.Birthdate, newUserInput.Pass, tok, 0, 0, newUserInput.Type, newUserInput.Pokeball)
	log.Println(newUser)

	//checking
	log.Println("check")
	// e, msg := utils.CheckUser(*newUser)
	// if e > 0 {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := ErrorResponse{
	// 		Code:    e,
	// 		Message: msg,
	// 	}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// Sending confirmation email
	log.Println("mailing")
	if err := utils.SendConfirmationEmail(newUserInput.Email, tok); err != nil {
		http.Error(w, "Erreur lors de l'envoi de l'e-mail de confirmation", http.StatusInternalServerError)
		return
	}

	//register to db
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
	claims, valid := CheckToken(w, r)
	if !valid {
		return
	}
	fmt.Println(claims.mail)
	decoder := json.NewDecoder(r.Body)
	db := database.ConnectDb()
	usr := models.User{}
	err := decoder.Decode(&usr)
	if err != nil {
		fmt.Println(err)
	}
	database.UpdateUser(db, usr)
	w.WriteHeader(http.StatusOK)
}

func DeleteUsersById(w http.ResponseWriter, r *http.Request) {
	claims, valid := CheckToken(w, r)
	if !valid {
		return
	}
	fmt.Println(claims.mail)
	params := mux.Vars(r)
	tofind, _ := strconv.Atoi(params["id"])
	db := database.ConnectDb()
	database.DelUserById(db, tofind)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	// if !CheckToken(w, r) {
	// 	return
	// }
	db := database.ConnectDb()
	database.DelUsers(db)
	fmt.Println("USERS DROPPED")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	claims, valid := CheckToken(w, r)
	if !valid {
		return
	}
	fmt.Println(claims.mail)

	//todo sortir toute la partie sve image dans une fonction a part + gerer les 5 images max
	r.ParseMultipartForm(10 << 20) // Max file size 10MB

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to read the uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageIndex := r.FormValue("index")

	// You can adjust the path according to your server's directory structure
	uploadDir := "./uploads/"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
		return
	}

	filePath := uploadDir + handler.Filename
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create the file on the server", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Failed to save the file on the server", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully")

	// now insert in db
	db := database.ConnectDb()
	user := database.GetUsersByEmail(db, claims.mail)
	// Convert imageIndex to an integer
	index, err := strconv.Atoi(imageIndex)
	if err != nil {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	// Check if index is within the valid range
	if index < 0 || index >= 4 {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	// Update user.Imageurl at the specific index
	user.Imageurl[index] = filePath
	database.UpdateUser(db, user)

}
