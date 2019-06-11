package main

import (
	"github.com/gorilla/mux"
	"github.com/Go_Rest_Api/controllers"
	"github.com/Go_Rest_Api/app"
	"os"
	"fmt"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	

	// attach JWT auth middleware
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println(port)
	err := http.ListenAndServe(":" + port, router) //Launch the app
	if err != nil {
		fmt.Print(err)
	} 
}