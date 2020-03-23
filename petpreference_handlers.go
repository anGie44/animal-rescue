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

var petPreferences []*PetPreference
var petPreferenceSeq = intSeq()

var createPetPreferenceHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	petPreference := PetPreference{}
	petPreference.ID = petPreferenceSeq()
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
	petPreferences = append(petPreferences, &petPreference)
	payload, _ := json.Marshal(petPreference)
	w.Write([]byte(payload))
})

var getPetPreferencesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	petPrefsListBytes, err := json.Marshal(petPreferences)
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
	// GET PetPreference and remove from list stored in memory
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
