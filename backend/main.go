package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
        "github.com/tgrangeo/matcha/models"
        "github.com/tgrangeo/matcha/database"
)

func test(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "hello test")
}

func getAllUsers(w http.ResponseWriter, r *http.Request){
        tab := []models.User{}
}

func main() {
        db :=  database.ConnectDb()
        // defer database.DropUsers(db)
        defer db.Close()
        database.CreateTable(db)
        user := models.User{"thomas","grangeon","hello world","thomas.alkpote@valo.fr"}
        database.InsertUser(db, user)

        user1 := models.User{"eliott","depauw","best sage ever","eliott.nekfeu@valo.fr"}
        database.InsertUser(db, user1)

        database.GetUsers(db)



        mux := mux.NewRouter()
        mux.HandleFunc("/test", test)
        mux.HandleFunc("/Users", getAllUsers)
        http.ListenAndServe(":8080", mux)
}