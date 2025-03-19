package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"policies/config"
	"policies/handlers"
	"policies/storage"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.DBConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	policyStorage := storage.NewDBPolicyStorage(db)
	policyHandler := handlers.NewPolicyHandler(policyStorage)

	r := mux.NewRouter()
	r.HandleFunc("/policy", policyHandler.CreatePolicy).Methods("POST")

	log.Printf("Server started on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
