package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/izaakdale/goBank/domain"
	"github.com/izaakdale/goBank/service"
	"github.com/jmoiron/sqlx"
)

func stanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Missing env variable ")
	}
}

func openDbConnectionClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_SCHEMA")

	// client, err := sqlx.Open("mysql", "root:root@/go_bank")
	client, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbSchema))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func Start() {

	stanityCheck()
	router := mux.NewRouter()

	//wiring
	dbClient := openDbConnectionClient()
	customerRepoDb := domain.NewCustomerRepoDb(dbClient)
	accountRepoDb := domain.NewAccountRepoDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerRepoService(customerRepoDb)}
	ah := AccountHandler{service.NewAccountService(accountRepoDb)}

	router.HandleFunc("/customers", ch.getCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.HandleFunc("/customers/{id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.Transaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	authM := AuthMiddleware{domain.NewAuthRepo()}
	router.Use(authM.authorizationHandler())

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
