package main

// AnimalRescue represents an organization
// consisting of Adopter and Adoptee Relationships
type AnimalRescue struct {
	Adopters       map[int]*Adopter
	Adoptees       map[int]*Adoptee
	Adoptions      map[int]*Adoption
	PetPreferences map[int]*PetPreference

	AdopterSeq       func() int
	AdopteeSeq       func() int
	AdoptionSeq      func() int
	PetPreferenceSeq func() int
}

func (ar *AnimalRescue) init() {
	adopters := make(map[int]*Adopter, 0)
	adoptees := make(map[int]*Adoptee, 0)
	adoptions := make(map[int]*Adoption, 0)
	petPreferences := make(map[int]*PetPreference, 0)
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

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
