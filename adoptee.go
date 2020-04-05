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

func (ar *AnimalRescue) CreateAdoptee(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	adoptee := new(Adoptee)
	err := decoder.Decode(adoptee)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adoptee.ID = ar.AdopteeSeq()
	ar.Adoptees = append(ar.Adoptees, adoptee)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(adoptee)

}

func (ar *AnimalRescue) GetAdoptee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoptee ID not valid"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	adoptee, _ := ar.GetAdopteeByID(id)
	if adoptee != nil {
		payload, _ := json.Marshal(adoptee)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Adoptee Not Found"))
	}

}

func (ar *AnimalRescue) GetAdoptees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ar.Adoptees)
}

func (ar *AnimalRescue) UpdateAdoptee(w http.ResponseWriter, r *http.Request) {
	// GET adoptee and update fields
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoptee ID not valid"))
		return
	}
	adoptee, _ := ar.GetAdopteeByID(id)

	if adoptee == nil {
		w.Write([]byte(fmt.Sprintf("Adoptee with ID %d not found", id)))
		return
	}

	decoder := json.NewDecoder(r.Body)
	updatedAdoptee := new(Adoptee)
	err = decoder.Decode(updatedAdoptee)

	if err == nil {
		updatedAdoptee.ID = id
		*adoptee = *updatedAdoptee
		json.NewEncoder(w).Encode(adoptee)
	}

}

func (ar *AnimalRescue) DeleteAdoptee(w http.ResponseWriter, r *http.Request) {
	// GET adoptee and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoptee ID not valid"))
		return
	}

	adoptee, index := ar.GetAdopteeByID(id)
	if adoptee != nil {
		ar.RemoveAdopteeByID(index)
		fmt.Fprintf(w, "The adoptee with ID %v has been deleted successfully", id)
	} else {
		fmt.Fprintf(w, "The adoptee with ID %v was not found", id)
	}

}
