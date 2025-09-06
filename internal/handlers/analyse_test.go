package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/takhleeq-ai/finlens-go/models"
)

func TestAnalyseHandler(t *testing.T) {
	// Prepare input
	transactions := []models.Transaction{
		{Description: "Freelance", Amount: 1200},
		{Description: "Tesco", Amount: -150},
		{Description: "Rent", Amount: -700},
		{Description: "Uber", Amount: -30},
	}

	body, _ := json.Marshal(transactions)
	req, err := http.NewRequest("POST", "/api/analyse", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a test HTTP recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AnalyseHandler)

	// Run the handler
	handler.ServeHTTP(rr, req)

	// Check HTTP status
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", rr.Code)
	}

	// Check JSON response
	var result map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}

	// Check that keys exist
	if _, ok := result["summary"]; !ok {
		t.Error("Expected 'summary' key in response")
	}
	if _, ok := result["transactions"]; !ok {
		t.Error("Expected 'transactions' key in response")
	}
}
