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

func main() {
        db := OpenDb()
        list := models.InitType()
        fmt.Println(list)
        usr := &models.User{
                ID:0,
                Name: "thomas",
        } 
        mux := mux.NewRouter()
        mux.HandleFunc("/hello",
                func(w http.ResponseWriter, r *http.Request) {
                        fmt.Fprintf(w, usr.Name)
                })
        mux.HandleFunc("/test", test)
        http.ListenAndServe(":8080", mux)

}