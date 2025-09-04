package logic

import "github.com/takhleeq-ai/finlens-go/models"

type Summary struct {
	Income      float64 `json:"total_income"`
	Expenses    float64 `json:"total_expenses"`
	Surplus     float64 `json:"surplus"`
	HealthScore float64 `json:"health_score"`
}

func GenerateSummary(txs []models.Transaction) Summary {
	var income, expenses float64
	for _, tx := range txs {
		if tx.Category == "Income" {
			income += tx.Amount
		} else if tx.Amount < 0 {
			expenses += -tx.Amount
		}
	}
	surplus := income - expenses
	score := 0.0
	if income > 0 {
		score = (surplus / income) * 100
	}
	return Summary{
		Income:      income,
		Expenses:    expenses,
		Surplus:     surplus,
		HealthScore: score,
	}
}
