package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Adopter struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	Gender         string    `json:"gender"`
	Birthdate      time.Time `json:"birthdate"`
	Address        string    `json:"address"`
	Country        string    `json:"country"`
	State          string    `json:"state"`
	City           string    `json:"city"`
	ZipCode        string    `json:"zip_code"`
	PetPreferences []string  `json:"pet_preferences"`
}

type PetPreference struct {
	Breed  string `json:"breed"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

var adopters []Adopter

func getAdopterHandler(w http.ResponseWriter, r *http.Request) {
	adopterListBytes, err := json.Marshal(adopters)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(adopterListBytes)
}

func createAdopterHandler(w http.ResponseWriter, r *http.Request) {
	adopter := Adopter{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adopter.FirstName = r.Form.Get("first_name")
	adopter.LastName = r.Form.Get("last_name")

	adopters = append(adopters, adopter)

	// redirect user to original HTML page
	http.Redirect(w, r, "/assets/", http.StatusFound)

}
