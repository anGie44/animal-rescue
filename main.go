package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Defining the API
	r.Handle("/status", statusHandler).Methods("GET")

	r.Handle("/adopter", createAdopterHandler).Methods("POST")
	r.Handle("/adopters", getAdoptersHandler).Methods("GET")
	r.Handle("/adopter/{id}", getAdopterHandler).Methods("GET")
	r.Handle("/adopter/{id}/update", updateAdopterHandler).Methods("GET")
	r.Handle("/adopter/{id}/delete", deleteAdopterHandler).Methods("GET")

	r.Handle("/adoptee", createAdopteeHandler).Methods("GET")
	r.Handle("/adoptees", getAdopteesHandler).Methods("GET")
	r.Handle("/adoptee/{id}", getAdopteeHandler).Methods("GET")
	r.Handle("/adoptee/{id}/update", updateAdopteeHandler).Methods("GET")
	r.Handle("/adoptees/{id}/delete", deleteAdopteeHandler).Methods("GET")

	r.Handle("/petpref", createPetPreferenceHandler).Methods("GET")
	r.Handle("/petprefs", getPetPreferencesHandler).Methods("GET")
	r.Handle("/petpref/{id}/update", updatePetPrefenceHandler).Methods("GET")
	r.Handle("/petpref/{id}/delete", deletePetPreferenceHandler).Methods("GET")

	r.Handle("/adoption/adopter/{adopter_id}/adoptee/{adoptee_id}", createAdoptionHandler).Methods("GET")
	r.Handle("/adoptions", getAdoptionsHandler).Methods("GET")
	r.Handle("/adoption/{id}", getAdoptionHandler).Methods("GET")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})
