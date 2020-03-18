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
	// Routes with token verification include:
	// /status
	// /adopters - will retrieve list of current/to-be pet adopters
	// /adoptees - will retrieve list of current animals up for adoptions
	r.Handle("/status", statusHandler).Methods("GET")

	r.Handle("/adopter", createAdopterHandler).Methods("POST")
	r.Handle("/adopters", getAdoptersHandler).Methods("GET")
	r.Handle("/adopter/{id}", getAdopterHandler).Methods("GET")
	r.Handle("/adopter/{id}/update", updateAdopterHandler).Methods("GET")
	r.Handle("/adopter/{id}/delete", deleteAdopterHandler).Methods("GET")

	r.Handle("/adoptee", createAdopteeHandler).Methods("POST")
	r.Handle("/adoptees", getAdopteesHandler).Methods("GET")
	r.Handle("/adoptee/{id}", getAdopteeHandler).Methods("GET")
	r.Handle("/adoptee/{id}/update", updateAdopteeHandler).Methods("GET")
	r.Handle("/adoptees/{id}/delete", deleteAdopteeHandler).Methods("GET")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})
