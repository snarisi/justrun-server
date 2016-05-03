package main

import (
	"net/http"
	"os"
	"github.com/julienschmidt/httprouter"
	"justrun-server/controllers/home"
	"justrun-server/db"
)

func main() {
	db.Connect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()

	router.GET("/", home.SayHello)

	http.ListenAndServe(":"+port, router)
}
