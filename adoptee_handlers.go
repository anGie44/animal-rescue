package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Adoptee struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Breed string `json:"breed"`
	Gender string `json:"gender"`
	Age string `json:"age"`
}

var adoptees []Adoptee

func getAdopteeHandler(w http.ResponseWriter, r *http.Request) {
	adopteeListBytes, err := json.Marshal(adoptees)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(adopteeListBytes)

}

func createAdopteeHandler(w http.ResponseWriter, r *http.Request) {
	adoptee := Adoptee{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adoptee.Name = r.Form.Get("name")
	adoptee.Breed = r.Form.Get("breed")
	adoptee.Gender = r.Form.Get("gender")
	adoptee.Age = r.Form.Get("age")

	adoptees = append(adoptees, adoptee)

	http.Redirect(w, r, "/assets/", http.StatusFound)

}

