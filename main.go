package main

import (
	"log"
	"net/http"
)

func main() {
	app := NewApp()
	app.HandleFunc("/dates", app.getDates)
	app.HandleFunc("/users", app.getUsers)
	app.HandleFunc("/add", app.addUser)

	log.Println("Starting server :4000")
	log.Fatal(http.ListenAndServe(":4000", app))
}
