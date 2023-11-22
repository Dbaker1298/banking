package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Dbaker1298/banking/domain"
	"github.com/Dbaker1298/banking/logger"
	"github.com/Dbaker1298/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWD") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("Environment variable(s) not defined. Check your .env file to set the environment variables")
	}
}

func Start() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sanityCheck()

	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDB)}

	// define the routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// start the server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info((fmt.Sprintf("Starting the application on %s:%s", address, port)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	client.SetConnMaxLifetime(time.Minute * 3)
	return client
}
