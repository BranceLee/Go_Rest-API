package main

import (
	"fmt"
	"github.com/Go_Rest_Api/app"
	"github.com/Go_Rest_Api/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContacts).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContacts).Methods("GET")

	// attach JWT auth middleware
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("> Server runs on ", port)
	err := http.ListenAndServe(":"+port, router) //Launch the app
	if err != nil {
		fmt.Print(err)
	}
}
