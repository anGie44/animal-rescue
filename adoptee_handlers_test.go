package main

import (
	"bytes"
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
	// Create Adoptee
	form := newCreateAdopteeForm()
	status := createAdoptee(form)

	if status := recorder.Code; status != http.StatusFound {
		return err
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	actual := res[0]
	expected := adoptees[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

func TestCreateAdopteesHandler(t *testing.T) {
	adoptees = []Adoptee{
		{1, "Callie", "Australian Shepherd", "female", "puppy"},
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

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	expected := adoptees[1]

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
	form.Set("id", "1")
	form.Set("name", "Callie")
	form.Set("breed", "Australian Shepherd")
	form.Set("gender", "female")
	form.Set("age", "puppy")

	return &form

}

func createAdoptee(form *url.Values) int {

	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createAdopteeHandler)
	hf.ServeHTTP(recorder, req)

	return recorder.Code

}
