package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api/resource", AuthMiddleware(ProtectedHandler)).Methods("GET")

	// Wrap the router with a logging middleware
	http.Handle("/", LoggerMiddleware(r))

	// Start the server
	port := ":8080"
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// HomeHandler handles the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

// ProtectedHandler handles the protected route
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected resource!"))
}

// LoggerMiddleware is a middleware that logs incoming requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.Header.Get("User-Agent"),
			time.Since(start),
		)
	})
}

// AuthMiddleware is a simple authentication middleware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement your authentication logic here
		// For simplicity, we'll just check if a token exists in the header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add your more complex authentication logic here

		// If authenticated, call the next handler
		next.ServeHTTP(w, r)
	}
}
