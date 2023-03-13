package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/tgrangeo/go_serv/database"
	"github.com/tgrangeo/go_serv/models"
)

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

func main(){

	database.InitDb()
	database.InsertDB()
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/todos", getTodos)
	router.DELETE("/todos", deleteTodos)
	router.DELETE("/todos/name", deleteByName)
	router.POST("/todos", newTodo)

	http.ListenAndServe(":3000", router)
}

func printAll(todos []models.Todolist) {
	for x := range todos {
		fmt.Println(todos[x].Name)
	}
}

func deleteByName(w http.ResponseWriter, r *http.Request, data httprouter.Params){
	//parse request
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    var t string
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
    }
	//del in db
	database.DeleteByName(t)
}

func newTodo(w http.ResponseWriter, r *http.Request, data httprouter.Params){
	//parse request
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    var t models.Todolist
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
    }
	//create in db or return error
	if (database.Create(t.Name,t.Description) == 1){
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func deleteTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.WriteHeader(http.StatusOK)
	database.DeleteAll()
}

func getTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	todos := database.FindAll()
	//printAll(todos)
	response := &JsonResponse{Data: &todos}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
	print(response)
}


func openFile(file string)(string,error){
	ret, err := ioutil.ReadFile(file) 
	if err != nil {
		return "", err
	}
	return string(ret),err
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, _ := openFile("index.html")
	fmt.Fprint(w, index)
}