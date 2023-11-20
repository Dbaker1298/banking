package app

import (
	"encoding/json"
	"net/http"

	"github.com/Dbaker1298/banking/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "John Doe", City: "Chicago", Zipcode: "12345"},
	// 	{Name: "Jane Doe", City: "Chicago", Zipcode: "12345"},
	// }

	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	} else {
		writeResponse(w, http.StatusOK, customers)
		return
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	// customer := Customer{Name: "John Doe", City: "Chicago", Zipcode: "12345"}
	// fmt.Fprintf(w, "Customer ID: %s", customerId)
	// json.NewEncoder(w).Encode(customer)

	customer, err := ch.service.GetCustomer(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	} else {
		writeResponse(w, http.StatusOK, customer)
		return
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
