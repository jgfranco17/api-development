// router.go
package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jgfranco17/api-development/handlers"
	"github.com/jgfranco17/api-development/middleware"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/api/resource", middleware.AuthMiddleware(handlers.ProtectedHandler)).Methods("GET")

	// Wrap the router with a logging middleware
	return middleware.LoggerMiddleware(r)
}

func StartServer() {
	router := SetupRouter()

	// Start the server
	port := ":8080"
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
