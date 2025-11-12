package repository

import (
    "errors"
    "sort"
    "sync"

    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/models"
)

type InMemoryRepo struct {
    mu           sync.RWMutex
    transactions []models.Transaction
}

func NewInMemoryRepo() *InMemoryRepo {
    return &InMemoryRepo{}
}

func (r *InMemoryRepo) SaveAll(txns []models.Transaction) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.transactions = append(r.transactions, txns...)
}

func (r *InMemoryRepo) All() []models.Transaction {
    r.mu.RLock()
    defer r.mu.RUnlock()
    out := make([]models.Transaction, len(r.transactions))
    copy(out, r.transactions)
    return out
}

func (r *InMemoryRepo) Clear() {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.transactions = nil
}

// List issues (non-success) with pagination + sorting
// sortBy: "timestamp"|"amount"|"name"; order: "asc"|"desc"
func (r *InMemoryRepo) ListIssues(page, perPage int, sortBy, order string) ([]models.Transaction, int, error) {
    if perPage <= 0 || page <= 0 {
        return nil, 0, errors.New("invalid pagination params")
    }

    r.mu.RLock()
    filtered := make([]models.Transaction, 0)
    for _, t := range r.transactions {
        if t.Status != models.StatusSuccess {
            filtered = append(filtered, t)
        }
    }
    r.mu.RUnlock()

    // sorting
    cmp := func(i, j int) bool { return false }
    switch sortBy {
    case "timestamp":
        if order == "asc" {
            cmp = func(i, j int) bool { return filtered[i].Timestamp < filtered[j].Timestamp }
        } else {
            cmp = func(i, j int) bool { return filtered[i].Timestamp > filtered[j].Timestamp }
        }
    case "amount":
        if order == "asc" {
            cmp = func(i, j int) bool { return filtered[i].Amount < filtered[j].Amount }
        } else {
            cmp = func(i, j int) bool { return filtered[i].Amount > filtered[j].Amount }
        }
    case "name":
        if order == "asc" {
            cmp = func(i, j int) bool { return filtered[i].Name < filtered[j].Name }
        } else {
            cmp = func(i, j int) bool { return filtered[i].Name > filtered[j].Name }
        }
    default:
        // default sort by timestamp desc
        if order == "asc" {
            cmp = func(i, j int) bool { return filtered[i].Timestamp < filtered[j].Timestamp }
        } else {
            cmp = func(i, j int) bool { return filtered[i].Timestamp > filtered[j].Timestamp }
        }
    }

    sort.Slice(filtered, cmp)

    total := len(filtered)
    start := (page - 1) * perPage
    if start >= total {
        return []models.Transaction{}, total, nil
    }
    end := start + perPage
    if end > total {
        end = total
    }
    return filtered[start:end], total, nil
}
