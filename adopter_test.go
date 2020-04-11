package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func TestCreateAdopter(t *testing.T) {
	ar := new(AnimalRescue)
	ar.init()
	tests := []struct {
		FirstName string
		LastName  string
		Email     string
	}{
		{
			FirstName: "test_user",
			LastName:  "example",
			Email:     "test_user.example@gmail.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.FirstName, func(t *testing.T) {
			resp, err := newAdopterRequest(ar, tt.FirstName, tt.LastName, tt.Email)
			if err != nil {
				t.Error(err)
			}
			adopter, err := parseAdopter(resp)
			assertStatus(t, resp.StatusCode, http.StatusCreated)
			assertResponseBody(t, strconv.Itoa(adopter.ID), "1")
			assertResponseBody(t, adopter.FirstName, tt.FirstName)
			assertResponseBody(t, adopter.LastName, tt.LastName)
			assertResponseBody(t, adopter.Email, tt.Email)
		})
	}

}

func TestAdopterPetPreferences(t *testing.T) {
	animalRescue := new(AnimalRescue)
	animalRescue.init()
	adopter := new(Adopter)
	adopter.PetPreferences = `[{"id":1, "breed":"Daschound","age":"Puppy","gender":"Male"}]`
	expected := make([]*PetPreference, 0)
	petPref := &PetPreference{
		ID:     1,
		Breed:  "Daschound",
		Age:    "Puppy",
		Gender: "Male",
	}
	expected = append(expected, petPref)
	expectedPref := expected[0]
	actualPref := adopter.Preferences()[0]
	if !reflect.DeepEqual(actualPref, expectedPref) {
		t.Errorf("Adopter's PetPreference = %v; expected %v", actualPref, expectedPref)
	}
}

func newAdopterRequest(ar *AnimalRescue, fn, ln, email string) (*http.Response, error) {
	data, err := json.Marshal(map[string]string{
		"first_name": fn,
		"last_name":  ln,
		"email":      email,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/adopters", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ar.CreateAdopter)
	handler.ServeHTTP(rr, req)

	return rr.Result(), nil
}

func parseAdopter(resp *http.Response) (*Adopter, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	adopter := new(Adopter)
	err = json.Unmarshal([]byte(body), adopter)
	return adopter, err
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %s want %s", got, want)
	}
}
