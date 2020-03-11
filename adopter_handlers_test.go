package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetAdoptersHandler(t *testing.T) {
	petPreference := PetPreference{"Australian Shepherd", "puppy", "male"}
	petPreferenceMarshalled, err := json.Marshal(petPreference)
	petPreferenceStr := string(petPreferenceMarshalled)
	var petPreferences []string
	petPreferences = append(petPreferences, petPreferenceStr)

	adopters = []Adopter{
		{"Angie", "Pinilla", "973.971.9690", "littledoglover@gmail.com", "Female", "09/11/1992", "444 Leonard St",
			"USA", "NY", "11222", petPreferences},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getAdopterHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Adopter{"Angie", "Pinilla"}
	a := []Adopter{}
	err = json.NewDecoder(recorder.Body).Decode(&a)

	if err != nil {
		t.Fatal(err)
	}

	actual := a[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

func TestCreateAdoptersHandler(t *testing.T) {
	petPreference := PetPreference{"Australian Shepherd", "puppy", "male"}
	petPreferenceMarshaled, err := json.Marshal(petPreference)
	petPreferenceStr := string(petPreferenceMarshaled)
	var petPreferences []string
	petPreferences = append(petPreferences, petPreference)

	adopters = []Adopter{
		{"Angie", "Pinilla", "973.971.9690", "littledoglover@gmail.com", "Female", "09/11/1992", "444 Leonard St",
			"USA", "NY", "11222", petPreferences},
	}
	form := newCreateAdopterForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createAdopterHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Adopter{"Angie", "Pinilla", "973.971.9690", "littledoglover@gmail.com", "Female", "09/11/1992", "444 Leonard St",
		"USA", "NY", "11222", petPreferences}

	if err != nil {
		t.Fatal(err)
	}

	actual := adopters[1]
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateAdopterForm() *url.Values {
	form := url.Values{}
	form.Set("first_name", "Angie")
	form.Set("last_name", "Pinilla")
	form.Set("phone", "973.971.9690")
	form.Set("email", "littledoglover@gmail.com")
	form.Set("gender", "Female")
	form.Set("birthdate", "09/11/1992")
	form.Set("address", "444 Leonard St")
	form.Set("country", "USA")
	form.Set("state", "NY")
	form.Set("zip_code", "11222")

	petPreference := PetPreference{"Australian Shepherd", "puppy", "male"}
	petPreferenceMarshaled, err := json.Marshal(petPreference)
	petPrefereneStr := string(petPreferenceMarshaled)
	form.Set("pet_preferences", petPreferenceStr)

	return &form

}
