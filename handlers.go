package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (a *App) getDates(w http.ResponseWriter, r *http.Request) {
	log.Print("getDates handler")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	res, err := json.Marshal(a.days)
	if err != nil {
		log.Print("getDaset json marhsal err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	log.Print("getUsers handler")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	res, err := json.Marshal(a.users)
	if err != nil {
		log.Print("getUsers json marhsal err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (a *App) addUser(w http.ResponseWriter, r *http.Request) {
	log.Print("addUser handler")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query()
	name := q.Get("name")
	u, err := a.CreateUser(name)
	if err != nil {
		log.Print("addUser error:", err)
		return
	}
	log.Printf("addUser handler: user %v add successfully", u)
	fmt.Fprint(w, u)
}
