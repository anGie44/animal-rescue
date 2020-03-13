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
	// Routes include:
	// /status
	// /adopters - will retrieve list of current/to-be pet adopters
	// /adoptees - will retrieve list of current animals up for adoptions
	r.HandleFunc("/status", statusHandler).Methods("GET")
	r.HandleFunc("/adopters", createAdopterHandler).Methods("POST")
	r.HandleFunc("/adopters", getAdoptersHandler).Methods("GET")
	r.HandleFunc("/adopters/{id}", getAdopterHandler).Methods("GET")
	r.HandleFunc("/adoptees", createAdopteeHandler).Methods("POST")
	r.HandleFunc("/adoptees", getAdopteesHandler).Methods("GET")
	r.HandleFunc("/adoptees/{id}", getAdopteeHandler).Methods("GET")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})
