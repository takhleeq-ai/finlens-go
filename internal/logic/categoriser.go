package logic

import (
	"strings"

	"github.com/takhleeq-ai/finlens-go/models"
)

func CategoriseTransactions(txs []models.Transaction) []models.Transaction {
	for i, tx := range txs {
		desc := strings.ToLower(tx.Description)
		switch {
		case strings.Contains(desc, "freelance"), strings.Contains(desc, "salary"):
			txs[i].Category = "Income"
		case strings.Contains(desc, "tesco"), strings.Contains(desc, "grocery"):
			txs[i].Category = "Groceries"
		case strings.Contains(desc, "rent"):
			txs[i].Category = "Housing"
		case strings.Contains(desc, "uber"):
			txs[i].Category = "Transport"
		default:
			txs[i].Category = "Other"
		}
	}
	return txs
}
