# FinLens 🧠💸

An AI-powered income & expense verifier built with Golang — built to parse transaction data, categorise it, and summarise financial health. Designed for modern Open Banking use cases.

## 🔧 Features

- Categorises bank transactions
- Calculates income, expenses, surplus
- Health score: how well you're managing cash flow
- Optional: AI-generated financial summaries

## 🚀 Run Locally

```bash
go run main.go. 
```

## Example API Usage  

```bash
curl -X POST http://localhost:8080/api/analyse \
  -H "Content-Type: application/json" \
  --data-binary @data/transactions.json
```

## Structure

finlens-go/
├── main.go
├── internal/
│   ├── handlers/
│   └── logic/
├── models/
├── data/
