package main

import (
	"reflect"
	"testing"
)

func TestCreateAdopter(t *testing.T) {

}

func TestGetAdopter(t *testing.T) {

}

func TestGetAdopters(t *testing.T) {

}

func TestUpdateAdopter(t *testing.T) {

}

func TestDeleteAdopter(t *testing.T) {

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
