package service

import (
    "errors"

    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/models"
    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/repository"
)

type Service struct {
    repo *repository.InMemoryRepo
}

func NewService(r *repository.InMemoryRepo) *Service {
    return &Service{repo: r}
}

func (s *Service) UploadTransactions(txns []models.Transaction) error {
    if len(txns) == 0 {
        return errors.New("no transactions to upload")
    }
    s.repo.SaveAll(txns)
    return nil
}

func (s *Service) ComputeBalance() int64 {
    // balance = sum(credit SUCCESS) - sum(debit SUCCESS)
    txns := s.repo.All()
    var balance int64 = 0
    for _, t := range txns {
        if t.Status != models.StatusSuccess {
            continue
        }
        if t.Type == models.TypeCredit {
            balance += t.Amount
        } else if t.Type == models.TypeDebit {
            balance -= t.Amount
        }
    }
    return balance
}

type IssuesResult struct {
    Items []models.Transaction `json:"items"`
    Total int                 `json:"total"`
    Page  int                 `json:"page"`
    PerPage int               `json:"per_page"`
}

func (s *Service) ListIssues(page, perPage int, sortBy, order string) (IssuesResult, error) {
    items, total, err := s.repo.ListIssues(page, perPage, sortBy, order)
    if err != nil {
        return IssuesResult{}, err
    }
    return IssuesResult{
        Items: items,
        Total: total,
        Page: page,
        PerPage: perPage,
    }, nil
}
