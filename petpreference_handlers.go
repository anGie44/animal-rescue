package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type PetPreference struct {
	ID     int    `json:"id"`
	Breed  string `json:"breed"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

var petPrefs []*PetPreference

var createPetPreferenceHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	petPreference := PetPreference{}
	petPreference.ID = len(petPrefs) + 1
	requestQuery := r.URL.Query()

	for key := range requestQuery {
		value := requestQuery.Get(key)
		switch key {
		case "breed":
			petPreference.Breed = value
		case "age":
			petPreference.Age = value
		case "gender":
			petPreference.Gender = value
		}
	}
	petPrefs = append(petPrefs, &petPreference)
	payload, _ := json.Marshal(petPreference)
	w.Write([]byte(payload))
})

var getPetPreferencesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	petPrefsListBytes, err := json.Marshal(petPrefs)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(petPrefsListBytes)
})

var updatePetPrefenceHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// GET PetPreference and update fields
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Pet Preference ID not valid"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	petPreference, _ := getPetPreferenceByID(id)
	if petPreference != nil {
		requestQuery := r.URL.Query()
		for key := range requestQuery {
			value := requestQuery.Get(key)
			reflect.ValueOf(petPreference).Elem().FieldByName(key).SetString(value)
		}
		payload, _ := json.Marshal(petPreference)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Pet Preference Not Found"))
	}
})

var deletePetPreferenceHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// GET adopter and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Pet Preference ID not valid"))
		return
	}

	petPreference, index := getPetPreferenceByID(id)
	if petPreference != nil {
		removePetPreferenceByID(index)
		w.Write([]byte("Pet Preference with ID " + string(index) + " removed"))
	} else {
		w.Write([]byte("Pet Preference Not Found"))
	}
})

// Helper Functions

func getPetPreferenceByID(id int) (*PetPreference, int) {
	var petPreference *PetPreference
	var index int
	for i, p := range petPrefs {
		if p.ID == id {
			petPreference = p
			index = i
		}
	}
	return petPreference, index
}

func removePetPreferenceByID(index int) {
	var emptyPetPreference *PetPreference
	petPrefs[index] = petPrefs[len(petPrefs)-1]
	petPrefs[len(petPrefs)-1] = emptyPetPreference
	petPrefs = petPrefs[:len(petPrefs)-1]
}