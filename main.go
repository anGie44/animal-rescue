package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/auth0-community/go-auth0"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	jose "gopkg.in/square/go-jose.v2"
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
	r.Handle("/adopters", authMiddleware(createAdopterHandler)).Methods("POST")
	r.Handle("/adopters", authMiddleware(getAdoptersHandler)).Methods("GET")
	r.Handle("/adopters/{id}", authMiddleware(getAdopterHandler)).Methods("GET")
	r.Handle("/adoptees", authMiddleware(createAdopteeHandler)).Methods("POST")
	r.Handle("/adoptees", authMiddleware(getAdopteesHandler)).Methods("GET")
	r.Handle("/adoptees/{id}", authMiddleware(getAdopteeHandler)).Methods("GET")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

/* Handler for verifying token */
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte(os.Getenv("YOUR-API-CLIENT-SECRET"))
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{os.Getenv("YOUR-AUTH0-API-AUDIENCE")}

		configuration := auth0.NewConfiguration(secretProvider, audience, "https://"+os.Getenv("YOUR-AUTH0-DOMAIN")+".auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration, nil)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
