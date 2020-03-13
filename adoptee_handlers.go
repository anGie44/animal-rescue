package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Adoptee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

var adoptees []Adoptee

func createAdopteeHandler(w http.ResponseWriter, r *http.Request) {
	adoptee := Adoptee{}
	autoIncrement := len(adoptees)

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adoptee.ID = autoIncrement + 1
	adoptee.Name = r.Form.Get("name")
	adoptee.Breed = r.Form.Get("breed")
	adoptee.Gender = r.Form.Get("gender")
	adoptee.Age = r.Form.Get("age")

	adoptees = append(adoptees, adoptee)

	http.Redirect(w, r, "/assets/", http.StatusFound)

}

func getAdopteeHandler(w http.ResponseWriter, r *http.Request) {
	var adoptee *Adoptee
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoptee ID not valid"))
		return
	}

	for _, a := range adoptees {
		if a.ID == id {
			adoptee = &a
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if adoptee != nil {
		payload, _ := json.Marshal(adoptee)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Adoptee Not Found"))
	}

}

func getAdopteesHandler(w http.ResponseWriter, r *http.Request) {
	adopteeListBytes, err := json.Marshal(adoptees)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(adopteeListBytes)
}

func updateAdopteeHandler(w http.ResponseWriter) {
	// TODO
}

func deleteAdopteeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}
