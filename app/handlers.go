package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Dbaker1298/banking/service"
	"github.com/gorilla/mux"
)

// REST Handlers
type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "John Doe", City: "Chicago", Zipcode: "12345"},
	// 	{Name: "Jane Doe", City: "Chicago", Zipcode: "12345"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
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
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
		return
	}
}
