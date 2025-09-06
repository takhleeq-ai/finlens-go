package logic

import (
	"testing"

	"github.com/takhleeq-ai/finlens-go/models"
)

func TestCategoriseTransactions(t *testing.T) {
	tests := []struct {
		description string
		amount      float64
		expectedCat string
	}{
		{"Freelance Payment", 1200, "Income"},
		{"Tesco Store", -150, "Groceries"},
		{"Rent for March", -700, "Housing"},
		{"Uber trip", -25, "Transport"},
		{"Mystery vendor", -50, "Other"},
	}

	var txs []models.Transaction
	for _, test := range tests {
		txs = append(txs, models.Transaction{
			Description: test.description,
			Amount:      test.amount,
		})
	}

	categorised := CategoriseTransactions(txs)

	for i, tx := range categorised {
		if tx.Category != tests[i].expectedCat {
			t.Errorf("Expected '%s' for '%s', got '%s'",
				tests[i].expectedCat, tx.Description, tx.Category)
		}
	}
}
