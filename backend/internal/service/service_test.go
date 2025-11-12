package service

import (
    "testing"

    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/models"
    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/repository"
)

func TestComputeBalance(t *testing.T) {
    repo := repository.NewInMemoryRepo()
    svc := NewService(repo)

    txns := []models.Transaction{
        {Timestamp: 1, Name: "A", Type: models.TypeCredit, Amount: 1000, Status: models.StatusSuccess},
        {Timestamp: 2, Name: "B", Type: models.TypeDebit, Amount: 200, Status: models.StatusSuccess},
        {Timestamp: 3, Name: "C", Type: models.TypeDebit, Amount: 100, Status: models.StatusFailed},
    }
    _ = svc.UploadTransactions(txns)

    bal := svc.ComputeBalance()
    if bal != 800 {
        t.Fatalf("expected 800, got %d", bal)
    }
}
