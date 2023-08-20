package main

import (
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tgrangeo/matcha/database"

	// "github.com/tgrangeo/matcha/models"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tgrangeo/matcha/handler"
)

func main() {
	godotenv.Load(".env")
	db := database.ConnectDb()
	defer db.Close()
	database.CreateTable(db)

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			next.ServeHTTP(w, r)
		})
	}
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	//AUTH
	router.HandleFunc("/api/v1/signin", handler.SignIn).Methods("POST")
	router.HandleFunc("/api/v1/logout", handler.Logout).Methods("POST")
	router.HandleFunc("/api/v1/refresh", handler.Refresh).Methods("POST")

	// Route pour la confirmation d'inscription
	router.HandleFunc("/api/v1/validate", handler.ConfirmRegistration).Methods("POST")

	//Reset mdp
	router.HandleFunc("/api/v1/resetpass", handler.ResetPass).Methods("POST")
	router.HandleFunc("/api/v1/newpass", handler.NewPass).Methods("POST")

	//API USER
	router.HandleFunc("/api/v1/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handler.GetUsersById).Methods("GET")
	router.HandleFunc("/api/v1/users/{where}/{value}", handler.GetWhere).Methods("GET")
	// router.HandleFunc("/api/v1/users/?", handler.GetWhere).Methods("GET") URL PARAMS
	router.HandleFunc("/api/v1/users", handler.CreateNewUser).Methods("POST")
	router.HandleFunc("/api/v1/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users", handler.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/api/v1/users/{id}", handler.DeleteUsersById).Methods("DELETE")
	// upload images
	router.HandleFunc("api/v1/upload", handler.UploadHandler)

	//tags
	// TODO: hello
	// router.HandleFunc("/api/v1/tags", handler.GetTags).Methods("GET")

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}
