package main

import (
	"log"
	"net/http"
	"webhookapi/controller"

	"github.com/gorilla/mux"
)

// Middleware to handle CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Register routes
	r := mux.NewRouter()
	r.HandleFunc("/watchresponses", controller.ReminderWebhook).Methods("POST")

	// Wrap the router with CORS middleware
	handlerWithCORS := enableCORS(r)

	// Specify the port to run the server
	port := ":8080"
	log.Printf("Server is running on ports %s\n", port)

	// Start the server
	err := http.ListenAndServe(port, handlerWithCORS)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
