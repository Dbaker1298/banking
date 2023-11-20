package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	// define a route
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// start the server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
