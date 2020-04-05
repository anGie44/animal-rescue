package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PetPreference struct {
	ID     int    `json:"id"`
	Breed  string `json:"breed"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func (ar *AnimalRescue) CreatePetPreference(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	petPreference := new(PetPreference)
	err := decoder.Decode(petPreference)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	petPreference.ID = ar.PetPreferenceSeq()
	ar.PetPreferences = append(ar.PetPreferences, petPreference)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(petPreference)
}

func (ar *AnimalRescue) GetPetPreferences(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ar.PetPreferences)
}

func (ar *AnimalRescue) UpdatePetPreference(w http.ResponseWriter, r *http.Request) {
	// GET pet_preference and update fields
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Pet Preference ID not valid"))
		return
	}
	petPreference, _ := ar.GetPetPreferenceByID(id)

	if petPreference == nil {
		w.Write([]byte(fmt.Sprintf("Pet Preference with ID %d not found", id)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	updatedPetPreference := new(PetPreference)
	err = decoder.Decode(updatedPetPreference)

	if err == nil {
		updatedPetPreference.ID = id
		*petPreference = *updatedPetPreference
		json.NewEncoder(w).Encode(petPreference)
	}
}

func (ar *AnimalRescue) DeletePetPreference(w http.ResponseWriter, r *http.Request) {
	// GET pet_preference and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Pet Preference ID not valid"))
		return
	}

	petPreference, index := ar.GetPetPreferenceByID(id)
	if petPreference != nil {
		ar.RemovePetPreferenceByID(index)
		fmt.Fprintf(w, "The pet_preference with ID %v has been deleted successfully", id)
	} else {
		fmt.Fprintf(w, "The pet_preference with ID %v was not found", id)
	}
}
