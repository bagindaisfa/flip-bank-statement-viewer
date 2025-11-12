package models

import "time"

type TransactionType string
type TransactionStatus string

const (
    TypeDebit  TransactionType = "DEBIT"
    TypeCredit TransactionType = "CREDIT"

    StatusSuccess TransactionStatus = "SUCCESS"
    StatusFailed  TransactionStatus = "FAILED"
    StatusPending TransactionStatus = "PENDING"
)

type Transaction struct {
    Timestamp int64             `json:"timestamp"`
    Time      time.Time         `json:"time"`
    Name      string            `json:"name"`
    Type      TransactionType   `json:"type"`
    Amount    int64             `json:"amount"`
    Status    TransactionStatus `json:"status"`
    Desc      string            `json:"description"`
}
