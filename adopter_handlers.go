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

var createAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	adopter := Adopter{}
	petPreferences := []PetPreference{}
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
	adopter.Phone = r.Form.Get("phone")
	adopter.Email = r.Form.Get("email")
	adopter.Gender = r.Form.Get("gender")
	adopter.Birthdate = r.Form.Get("birthdate")
	adopter.Address = r.Form.Get("address")
	adopter.Country = r.Form.Get("country")
	adopter.State = r.Form.Get("state")
	adopter.City = r.Form.Get("city")
	adopter.ZipCode = r.Form.Get("zipcode")
	petPreferenceA := PetPreference{r.Form.Get("pet_preference_a_breed"), r.Form.Get("pet_preference_a_age"), r.Form.Get("pet_preference_a_gender")}
	petPreferenceB := PetPreference{r.Form.Get("pet_preference_b_breed"), r.Form.Get("pet_preference_b_age"), r.Form.Get("pet_preference_b_gender")}
	petPreferenceC := PetPreference{r.Form.Get("pet_preference_c_breed"), r.Form.Get("pet_preference_c_age"), r.Form.Get("pet_preference_c_gender")}
	petPreferences = append(petPreferences, petPreferenceA)
	petPreferences = append(petPreferences, petPreferenceB)
	petPreferences = append(petPreferences, petPreferenceC)
	adopter.PetPreferences = petPreferences

	adopters = append(adopters, adopter)

	// redirect user to original HTML page
	// http.Redirect(w, r, "/", http.StatusFound)
	payload, _ := json.Marshal(adopters)
	w.Write([]byte(payload))

})

var getAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var adopter Adopter
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}

	for _, a := range adopters {
		if a.ID == id {
			fmt.Println(a.ID, id)
			adopter = a
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if &adopter != nil {
		payload, _ := json.Marshal(adopter)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Adopter Not Found"))
	}

})

var getAdoptersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	adopterListBytes, err := json.Marshal(adopters)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(adopterListBytes)
})

var updateAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// TODO
})

var deleteAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// TODO
})
