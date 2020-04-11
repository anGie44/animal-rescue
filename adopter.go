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

func (a *Adopter) Preferences() []*PetPreference {
	var prefs []*PetPreference
	if len(a.PetPreferences) > 0 {
		petPrefBytes := []byte(a.PetPreferences)
		if err := json.Unmarshal(petPrefBytes, &prefs); err != nil {
			panic(err)
		}
	}
	return prefs
}

func (ar *AnimalRescue) CreateAdopter(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	adopter := new(Adopter)
	err := decoder.Decode(adopter)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adopter.ID = ar.AdopterSeq()
	ar.Adopters[adopter.ID] = adopter

	for _, pref := range adopter.Preferences() {
		pref.ID = ar.PetPreferenceSeq()
		ar.PetPreferences[pref.ID] = pref
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(adopter)

}

func (ar *AnimalRescue) GetAdopter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}

	adopter := ar.Adopters[id]
	if adopter != nil {
		json.NewEncoder(w).Encode(adopter)
	}
}

func (ar *AnimalRescue) GetAdopters(w http.ResponseWriter, r *http.Request) {
	adopters := make([]*Adopter, 0, len(ar.Adopters))
	for _, adopter := range ar.Adopters {
		adopters = append(adopters, adopter)
	}
	json.NewEncoder(w).Encode(adopters)
}

func (ar *AnimalRescue) UpdateAdopter(w http.ResponseWriter, r *http.Request) {
	// GET adopter and update fields
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}

	adopter := ar.Adopters[id]
	if adopter == nil {
		w.Write([]byte(fmt.Sprintf("Adopter with ID %d not found", id)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	updatedAdopter := new(Adopter)
	err = decoder.Decode(updatedAdopter)
	if err == nil {
		updatedAdopter.ID = id
		*adopter = *updatedAdopter
		json.NewEncoder(w).Encode(adopter)
	}
}

func (ar *AnimalRescue) DeleteAdopter(w http.ResponseWriter, r *http.Request) {
	// GET adopter and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte("Adopter ID not valid"))
		return
	}

	adopter := ar.Adopters[id]
	if adopter != nil {
		delete(ar.Adopters, id)
		fmt.Fprintf(w, "The adopter with ID %v has been deleted successfully", id)
	} else {
		fmt.Fprintf(w, "The adopter with ID %v was not found", id)
	}
}
