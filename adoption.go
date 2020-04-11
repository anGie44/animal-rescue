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

func (ar *AnimalRescue) CreateAdoption(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	adoption := new(Adoption)
	err := decoder.Decode(adoption)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	adoption.ID = ar.AdoptionSeq()
	ar.Adoptions[adoption.ID] = adoption

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(adoption)

}

func (ar *AnimalRescue) GetAdoption(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoption ID not valid"))
		return
	}

	adoption := ar.Adoptions[id]
	if adoption != nil {
		json.NewEncoder(w).Encode(adoption)
	}
}

func (ar *AnimalRescue) GetAdoptions(w http.ResponseWriter, r *http.Request) {
	adoptions := make([]*Adoption, 0, len(ar.Adoptions))
	for _, adoption := range ar.Adoptions {
		adoptions = append(adoptions, adoption)
	}
	json.NewEncoder(w).Encode(adoptions)
}

func (ar *AnimalRescue) DeleteAdoption(w http.ResponseWriter, r *http.Request) {
	// GET adoption and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoption ID not valid"))
		return
	}

	adoption := ar.Adoptions[id]
	if adoption != nil {
		delete(ar.Adoptions, id)
		fmt.Fprintf(w, "The adoption with ID %v has been deleted successfully", id)
	} else {
		fmt.Fprintf(w, "The adoption with ID %v was not found", id)
	}
}
