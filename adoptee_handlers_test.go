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

func TestGetAdopteesHandler(t *testing.T) {
	adoptees = []Adoptee{
		{"PS555048", "Callie", "Australian Shepherd", "female", "puppy"},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getAdopteeHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Adoptee{"PS555048", "Callie", "Australian Shepherd", "female", "puppy"}
	a := []Adoptee{}
	err = json.NewDecoder(recorder.Body).Decode(&a)

	if err != nil {
		t.Fatal(err)
	}

	actual := a[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

func TestCreateAdopteesHandler(t *testing.T) {
	adoptees = []Adoptee{
		{"PS555048", "Callie", "Australian Shepherd", "female", "puppy"},
	}
	form := newCreateAdopteeForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createAdopteeHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Adoptee{"PS555048", "Callie", "Australian Shepherd", "female", "puppy"}

	if err != nil {
		t.Fatal(err)
	}

	actual := adoptees[1]
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateAdopteeForm() *url.Values {
	form := url.Values{}
	form.Set("id", "PS555048")
	form.Set("name", "Callie")
	form.Set("breed", "Australian Shepherd")
	form.Set("gender", "female")
	form.Set("age", "puppy")

	return &form

}
