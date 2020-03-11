package main

import (
	"net/http"

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
	r.HandleFunc("/status", NotImplemented).Methods("GET")
	r.HandleFunc("/adopters", getAdopterHandler).Methods("GET")
	r.HandleFunc("/adopters", createAdopterHandler).Methods("POST")
	r.HandleFunc("/adoptees", getAdopteeHandler).Methods("GET")
	r.HandleFunc("/adoptees", createAdopteeHandler).Methods("POST")

	http.ListenAndServe(":3000", r)

}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
