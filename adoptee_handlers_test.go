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
	adoptees := []Adoptee{
		{1, "Callie", "Australian Shepherd", "female", "puppy"},
	}
	// CREATE Adoptee
	form := newCreateAdopteeForm()
	responseRecorder, err := createAdoptee(form)
	if err != nil {
		t.Fatal(err)
	}
	if status := responseRecorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	// GET Adoptee
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(getAdopteesHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	a := []Adoptee{}
	err = json.NewDecoder(recorder.Body).Decode(&a)
	actual := a[0]
	expected := adoptees[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

func TestCreateAdopteesHandler(t *testing.T) {
	form := newCreateAdopteeForm()
	responseRecorder, err := createAdoptee(form)
	if err != nil {
		t.Fatal(err)
	}
	if status := responseRecorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	a := []Adoptee{}
	err = json.NewDecoder(responseRecorder.Body).Decode(&a)

	if len(a) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", a, []Adoptee{})
	}
}

func newCreateAdopteeForm() *url.Values {
	form := url.Values{}
	form.Set("id", "1")
	form.Set("name", "Callie")
	form.Set("breed", "Australian Shepherd")
	form.Set("gender", "female")
	form.Set("age", "puppy")

	return &form

}

func createAdoptee(form *url.Values) (*httptest.ResponseRecorder, error) {

	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createAdopteeHandler)
	hf.ServeHTTP(recorder, req)

	return recorder, nil

}
