package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Adopter struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Birthdate      string `json:"birthdate"`
	Address        string `json:"address"`
	Country        string `json:"country"`
	State          string `json:"state"`
	City           string `json:"city"`
	ZipCode        string `json:"zip_code"`
	PetPreferences string `json:"pet_preferences"`
}

func (a *Adopter) Preferences() []PetPreference {
	var prefs []PetPreference
	if len(a.PetPreferences) > 0 {
		petPrefBytes := []byte(a.PetPreferences)
		json.Unmarshal(petPrefBytes, &prefs)
	}
	return prefs
}

var adopters []*Adopter
var adopterSeq = intSeq()
var petPrefSeq = intSeq()

var createAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var adopter Adopter
	err := decoder.Decode(&adopter)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	adopter.ID = adopterSeq()
	adopters = append(adopters, &adopter)
	for _, pref := range adopter.Preferences() {
		petPreferences = append(petPreferences, &pref)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(adopter)

})

var getAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}

	adopter, _ := getAdopterByID(id)
	if adopter != nil {
		json.NewEncoder(w).Encode(adopter)
	}
})

var getAdoptersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(adopters)
})

var updateAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// GET adopter and update fields
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}
	adopter, _ := getAdopterByID(id)

	if adopter == nil {
		w.Write([]byte(fmt.Sprintf("Adopter with ID %s not found", id)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var updatedAdopter Adopter
	err = decoder.Decode(&updatedAdopter)

	if err == nil {
		adopter.FirstName = updatedAdopter.FirstName
		adopter.LastName = updatedAdopter.LastName
		adopter.Phone = updatedAdopter.Phone
		adopter.Email = updatedAdopter.Email
		adopter.Gender = updatedAdopter.Gender
		adopter.Birthdate = updatedAdopter.Birthdate
		adopter.Address = updatedAdopter.Address
		adopter.Country = updatedAdopter.Country
		adopter.State = updatedAdopter.State
		adopter.City = updatedAdopter.City
		adopter.ZipCode = updatedAdopter.ZipCode

		json.NewEncoder(w).Encode(adopter)
	}
})

var deleteAdopterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// GET adopter and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}

	adopter, index := getAdopterByID(id)
	if adopter != nil {
		removeAdopterByID(index)
		fmt.Fprintf(w, "The adopter with ID %v has been deleted successfully", id)
	} else {
		fmt.Fprintf(w, "The adopter with ID %v was not found", id)
	}

})
