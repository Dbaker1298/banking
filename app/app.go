package app

import (
	"log"
	"net/http"
)

func Start() {
	// define a route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// start the server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
