package main

import (
	"net/http"
	"os"
	"github.com/julienschmidt/httprouter"
	"justrun-server/controllers/home"
	"justrun-server/controllers/users"
	"justrun-server/db"
)

func main() {
	db.Connect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}

	router := httprouter.New()

	router.GET("/", home.SayHello)

	router.GET("/users", users.FindAll)
	router.GET("/users/:email", users.FindByEmail)
	router.POST("/users", users.Create)

	http.ListenAndServe(":"+port, router)
}
