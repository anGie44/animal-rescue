package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Adoptee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
}

var adoptees []*Adoptee

var createAdopteeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	adoptee := Adoptee{}
	autoIncrement := len(adoptees)

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adoptee.ID = autoIncrement + 1
	adoptee.Name = r.Form.Get("name")
	adoptee.Breed = r.Form.Get("breed")
	adoptee.Gender = r.Form.Get("gender")
	adoptee.Age = r.Form.Get("age")

	adoptees = append(adoptees, &adoptee)

	http.Redirect(w, r, "/assets/", http.StatusFound)

})

var getAdopteeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoptee ID not valid"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	adoptee, _ := getAdopteeByID(id)
	if adoptee != nil {
		payload, _ := json.Marshal(adoptee)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Adoptee Not Found"))
	}

})

var getAdopteesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	adopteeListBytes, err := json.Marshal(adoptees)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(adopteeListBytes)
})

var updateAdopteeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// GET adoptee and update fields
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoptee ID not valid"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	adoptee, _ := getAdopteeByID(id)
	if adoptee != nil {
		requestQuery := r.URL.Query()
		for key := range requestQuery {
			newKey := strings.Title(key)
			value := requestQuery.Get(key)
			reflect.ValueOf(adoptee).Elem().FieldByName(newKey).SetString(value)
		}
		payload, _ := json.Marshal(adoptee)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Adoptee Not Found"))
	}

})

var deleteAdopteeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// GET adoptee and remove from list stored in memory
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte("Adoptee ID not valid"))
		return
	}

	adoptee, index := getAdopteeByID(id)
	if adoptee != nil {
		removeAdopteeByID(index)
		w.Write([]byte("Adoptee with ID " + string(index) + " removed"))
	} else {
		w.Write([]byte("Adoptee Not Found"))
	}

})

func getAdopteeByID(id int) (*Adoptee, int) {
	var adoptee *Adoptee
	var index int
	for i, a := range adoptees {
		if a.ID == id {
			adoptee = a
			index = i
		}
	}
	return adoptee, index
}

func removeAdopteeByID(index int) {
	var emptyAdoptee *Adoptee
	adoptees[index] = adoptees[len(adoptees)-1]
	adoptees[len(adoptees)-1] = emptyAdoptee
	adoptees = adoptees[:len(adoptees)-1]
}
