package logic

import (
	"testing"

	"github.com/takhleeq-ai/finlens-go/models"
)

func TestGenerateSummary(t *testing.T) {
	transactions := []models.Transaction{
		{Description: "Freelance", Amount: 1200, Category: "Income"},
		{Description: "Rent", Amount: -500, Category: "Housing"},
		{Description: "Groceries", Amount: -200, Category: "Groceries"},
	}

	summary := GenerateSummary(transactions)

	if summary.Income != 1200 {
		t.Errorf("Expected income to be 1200, got %.2f", summary.Income)
	}
	if summary.Expenses != 700 {
		t.Errorf("Expected expenses to be 700, got %.2f", summary.Expenses)
	}
	if summary.Surplus != 500 {
		t.Errorf("Expected surplus to be 500, got %.2f", summary.Surplus)
	}
	if summary.HealthScore != 41.67 {
		t.Errorf("Expected health score to be 41.67, got %.2f", summary.HealthScore)
	}
}
