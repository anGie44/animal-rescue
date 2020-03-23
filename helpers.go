package main

import "strings"

func getAdopterByID(id int) (*Adopter, int) {
	var adopter *Adopter
	var index int
	for i, a := range adopters {
		if a.ID == id {
			adopter = a
			index = i
		}
	}
	return adopter, index
}

func removeAdopterByID(index int) {
	var emptyAdopter *Adopter
	adopters[index] = adopters[len(adopters)-1]
	adopters[len(adopters)-1] = emptyAdopter
	adopters = adopters[:len(adopters)-1]
}

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

func getAdoptionByID(id int) (*Adoption, int) {
	var adoption *Adoption
	var index int
	for i, a := range adoptions {
		if a.ID == id {
			adoption = a
			index = i
		}
	}
	return adoption, index
}

func removeAdoptionByID(index int) {
	var emptyAdoption *Adoption
	adoptions[index] = adoptions[len(adoptions)-1]
	adoptions[len(adoptions)-1] = emptyAdoption
	adoptions = adoptions[:len(adoptions)-1]
}

func getPetPreferenceByID(id int) (*PetPreference, int) {
	var petPreference *PetPreference
	var index int
	for i, p := range petPreferences {
		if p.ID == id {
			petPreference = p
			index = i
		}
	}
	return petPreference, index
}

func removePetPreferenceByID(index int) {
	var emptyPetPreference *PetPreference
	petPreferences[index] = petPreferences[len(petPreferences)-1]
	petPreferences[len(petPreferences)-1] = emptyPetPreference
	petPreferences = petPreferences[:len(petPreferences)-1]
}

func formatKey(key string) string {
	var str string
	if strings.Contains(key, "_") {
		str = strings.Replace(key, "_", " ", -1)
		str = strings.Title(str)
		str = strings.Replace(str, " ", "", -1)

	} else {
		str = strings.Title(key)
	}
	return str
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
