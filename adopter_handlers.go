package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Adopter struct {
	ID             int             `json:"id"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Phone          string          `json:"phone"`
	Email          string          `json:"email"`
	Gender         string          `json:"gender"`
	Birthdate      string          `json:"birthdate"`
	Address        string          `json:"address"`
	Country        string          `json:"country"`
	State          string          `json:"state"`
	City           string          `json:"city"`
	ZipCode        string          `json:"zip_code"`
	PetPreferences []PetPreference `json:"pet_preferences"`
}

type PetPreference struct {
	Breed  string `json:"breed"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

var adopters []Adopter

func createAdopterHandler(w http.ResponseWriter, r *http.Request) {
	adopter := Adopter{}
	autoIncrement := len(adopters) + 1

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adopter.ID = autoIncrement
	adopter.FirstName = r.Form.Get("first_name")
	adopter.LastName = r.Form.Get("last_name")

	adopters = append(adopters, adopter)

	// redirect user to original HTML page
	http.Redirect(w, r, "/assets/", http.StatusFound)

}

func getAdopterHandler(w http.ResponseWriter, r *http.Request) {
	var adopter *Adopter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}

	for _, a := range adopters {
		if a.ID == id {
			adopter = &a
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if adopter != nil {
		payload, _ := json.Marshal(adopter)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Adopter Not Found"))
	}

}

func getAdoptersHandler(w http.ResponseWriter, r *http.Request) {
	adopterListBytes, err := json.Marshal(adopters)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(adopterListBytes)
}

func updateAdopterHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func deleteAdopterHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}
