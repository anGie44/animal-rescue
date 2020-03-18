package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Adoption struct {
	ID      int      `json:"id"`
	Adopter *Adopter `json:"adopter"`
	Adoptee *Adoptee `json:"adoptee"`
	Date    string   `json:"date"`
}

var adoptions []*Adoption

var createAdoptionHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adopterID, err := strconv.Atoi(vars["adopter_id"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	adopteeID, err := strconv.Atoi(vars["adoptee_id"])
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adopter, _ := getAdopterByID(adopterID)
	adoptee, _ := getAdopteeByID(adopteeID)
	if adopter == nil || adoptee == nil {
		w.Write([]byte("Adopter ID and/or Adoptee ID not found"))
		return
	}
	adoption := Adoption{}
	adoption.ID = len(adoptions) + 1
	adoption.Adopter = adopter
	adoption.Adoptee = adoptee
	adoptions = append(adoptions, &adoption)

	payload, _ := json.Marshal(adoptions)
	w.Write([]byte(payload))

})

var getAdoptionHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoption ID not valid"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	adoption, _ := getAdoptionByID(id)
	if adoption != nil {
		payload, _ := json.Marshal(adoption)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Adoption Not Found"))
	}

})

var getAdoptionsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	adoptionListBytes, err := json.Marshal(adoptions)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(adoptionListBytes)
})

var deleteAdoptionHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// GET adopter and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoption ID not valid"))
		return
	}

	adoption, index := getAdoptionByID(id)
	if adoption != nil {
		removeAdoptionByID(index)
		w.Write([]byte("Adoption with ID " + string(index) + " removed"))
	} else {
		w.Write([]byte("Adoption Not Found"))
	}
})

// Helper Functions

func getAdoptionByID(id int) (*Adoption, int) {
	var adoption *Adoption
	var index int
	for i, a := range adoptions {
		if a.ID == id {
			adoption = a
			index = i
		}
	}
	return adoption, index
}

func removeAdoptionByID(index int) {
	var emptyAdoption *Adoption
	adoptions[index] = adoptions[len(adoptions)-1]
	adoptions[len(adoptions)-1] = emptyAdoption
	adoptions = adoptions[:len(adoptions)-1]
}
