package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var animalRescue *AnimalRescue

func main() {
	animalRescue = new(AnimalRescue)
	animalRescue.init()
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// Defining the API Routes
	r.Handle("/status", statusHandler).Methods("GET")

	r.HandleFunc("/adopter", animalRescue.CreateAdopter).Methods("POST")
	r.HandleFunc("/adopters", animalRescue.GetAdopters).Methods("GET")
	r.HandleFunc("/adopter/{id}", animalRescue.GetAdopter).Methods("GET")
	r.HandleFunc("/adopter/{id}", animalRescue.UpdateAdopter).Methods("PATCH")
	r.HandleFunc("/adopter/{id}", animalRescue.DeleteAdopter).Methods("DELETE")

	r.HandleFunc("/adoptee", animalRescue.CreateAdoptee).Methods("POST")
	r.HandleFunc("/adoptees", animalRescue.GetAdoptees).Methods("GET")
	r.HandleFunc("/adoptee/{id}", animalRescue.GetAdoptee).Methods("GET")
	r.HandleFunc("/adoptee/{id}", animalRescue.UpdateAdoptee).Methods("PATCH")
	r.HandleFunc("/adoptees/{id}", animalRescue.DeleteAdoptee).Methods("DELETE")

	r.HandleFunc("/petpref", animalRescue.CreatePetPreference).Methods("POST")
	r.HandleFunc("/petprefs", animalRescue.GetPetPreferences).Methods("GET")
	r.HandleFunc("/petpref/{id}", animalRescue.UpdatePetPreference).Methods("PATCH")
	r.HandleFunc("/petpref/{id}", animalRescue.DeletePetPreference).Methods("DELETE")

	r.HandleFunc("/adoption", animalRescue.CreateAdoption).Methods("POST")
	r.HandleFunc("/adoptions", animalRescue.GetAdoptions).Methods("GET")
	r.HandleFunc("/adoption/{id}", animalRescue.GetAdoption).Methods("GET")
	r.HandleFunc("/adoption/{id}", animalRescue.DeleteAdoption).Methods("DELETE")

	rocketEmoji := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")

	fmt.Printf("Server running on port 8080 %s\n", rocketEmoji)
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))

}

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})
