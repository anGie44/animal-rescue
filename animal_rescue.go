package main

// AnimalRescue represents an organization
// consisting of Adopter and Adoptee Relationships
type AnimalRescue struct {
	Adopters       []*Adopter
	Adoptees       []*Adoptee
	Adoptions      []*Adoption
	PetPreferences []*PetPreference

	AdopterSeq       func() int
	AdopteeSeq       func() int
	AdoptionSeq      func() int
	PetPreferenceSeq func() int
}

func (ar *AnimalRescue) init() {
	adopters := make([]*Adopter, 0)
	adoptees := make([]*Adoptee, 0)
	adoptions := make([]*Adoption, 0)
	petPreferences := make([]*PetPreference, 0)
	adopterSeq := intSeq()
	adopteeSeq := intSeq()
	adoptionSeq := intSeq()
	petPrefSeq := intSeq()

	ar.Adopters = adopters
	ar.Adoptees = adoptees
	ar.Adoptions = adoptions
	ar.PetPreferences = petPreferences
	ar.AdopterSeq = adopterSeq
	ar.AdopteeSeq = adopteeSeq
	ar.AdoptionSeq = adoptionSeq
	ar.PetPreferenceSeq = petPrefSeq

}

// Adopter-related Helpers

func (ar *AnimalRescue) GetAdopterByID(id int) (*Adopter, int) {
	var adopter *Adopter
	var index int
	for i, a := range ar.Adopters {
		if a.ID == id {
			adopter = a
			index = i
		}
	}
	return adopter, index
}

func (ar *AnimalRescue) RemoveAdopterByID(index int) {
	var emptyAdopter *Adopter
	ar.Adopters[index] = ar.Adopters[len(ar.Adopters)-1]
	ar.Adopters[len(ar.Adopters)-1] = emptyAdopter
	ar.Adopters = ar.Adopters[:len(ar.Adopters)-1]
}

// Adoptee-related Helpers

func (ar *AnimalRescue) GetAdopteeByID(id int) (*Adoptee, int) {
	var adoptee *Adoptee
	var index int
	for i, a := range ar.Adoptees {
		if a.ID == id {
			adoptee = a
			index = i
		}
	}
	return adoptee, index
}

func (ar *AnimalRescue) RemoveAdopteeByID(index int) {
	var emptyAdoptee *Adoptee
	ar.Adoptees[index] = ar.Adoptees[len(ar.Adoptees)-1]
	ar.Adoptees[len(ar.Adoptees)-1] = emptyAdoptee
	ar.Adoptees = ar.Adoptees[:len(ar.Adoptees)-1]
}

// Adoption-related Helpers

func (ar *AnimalRescue) GetAdoptionByID(id int) (*Adoption, int) {
	var adoption *Adoption
	var index int
	for i, a := range ar.Adoptions {
		if a.ID == id {
			adoption = a
			index = i
		}
	}
	return adoption, index
}

func (ar *AnimalRescue) RemoveAdoptionByID(index int) {
	var emptyAdoption *Adoption
	ar.Adoptions[index] = ar.Adoptions[len(ar.Adoptions)-1]
	ar.Adoptions[len(ar.Adoptions)-1] = emptyAdoption
	ar.Adoptions = ar.Adoptions[:len(ar.Adoptions)-1]
}

// Pet Preference-related Helpers

func (ar *AnimalRescue) GetPetPreferenceByID(id int) (*PetPreference, int) {
	var petPreference *PetPreference
	var index int
	for i, p := range ar.PetPreferences {
		if p.ID == id {
			petPreference = p
			index = i
		}
	}
	return petPreference, index
}

func (ar *AnimalRescue) RemovePetPreferenceByID(index int) {
	var emptyPetPreference *PetPreference
	ar.PetPreferences[index] = ar.PetPreferences[len(ar.PetPreferences)-1]
	ar.PetPreferences[len(ar.PetPreferences)-1] = emptyPetPreference
	ar.PetPreferences = ar.PetPreferences[:len(ar.PetPreferences)-1]
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
