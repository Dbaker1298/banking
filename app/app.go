package app

import (
	"log"
	"net/http"

	"github.com/Dbaker1298/banking/domain"
	"github.com/Dbaker1298/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define the routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
