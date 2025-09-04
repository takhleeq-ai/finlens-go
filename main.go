package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/takhleeq-ai/finlens-go/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/analyse", handlers.AnalyseHandler).Methods("POST")

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
