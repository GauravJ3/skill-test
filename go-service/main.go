package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GauravJ3/go-service/handlers"

	"github.com/GauravJ3/go-service/external"

	"github.com/GauravJ3/go-service/config"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	if err := external.LoginAndSetTokens(); err != nil {
		log.Fatalf("Auth failed: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/students/{id}/report", handlers.ReportHandler).Methods("GET")

	fmt.Println("✅ Go PDF service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
