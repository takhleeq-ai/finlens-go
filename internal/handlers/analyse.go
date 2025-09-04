package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/takhleeq-ai/finlens-go/internal/logic"
	"github.com/takhleeq-ai/finlens-go/models"
)

func AnalyseHandler(w http.ResponseWriter, r *http.Request) {
	var txs []models.Transaction
	err := json.NewDecoder(r.Body).Decode(&txs)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	txs = logic.CategoriseTransactions(txs)
	summary := logic.GenerateSummary(txs)

	output := map[string]interface{}{
		"summary":      summary,
		"transactions": txs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
