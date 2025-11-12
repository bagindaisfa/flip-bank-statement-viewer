package utils

import (
    "encoding/csv"
    "io"
    "strconv"
    "strings"
    "time"

    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/models"
)

func ParseCSV(r io.Reader) ([]models.Transaction, error) {
    reader := csv.NewReader(r)
    reader.FieldsPerRecord = -1
    out := make([]models.Transaction, 0)

    for {
        rec, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        // skip empty lines
        if len(rec) == 0 {
            continue
        }
        // accept exactly 6 fields; trim spaces
        if len(rec) < 6 {
            // try to be permissive: skip invalid rows
            continue
        }
        tsStr := strings.TrimSpace(rec[0])
        name := strings.TrimSpace(rec[1])
        ttype := strings.ToUpper(strings.TrimSpace(rec[2]))
        amountStr := strings.TrimSpace(rec[3])
        status := strings.ToUpper(strings.TrimSpace(rec[4]))
        desc := strings.TrimSpace(rec[5])

        ts, err := strconv.ParseInt(tsStr, 10, 64)
        if err != nil {
            // skip invalid timestamp row
            continue
        }
        amt, err := strconv.ParseInt(amountStr, 10, 64)
        if err != nil {
            continue
        }
        txn := models.Transaction{
            Timestamp: ts,
            Time: time.Unix(ts, 0),
            Name: name,
            Type: models.TransactionType(ttype),
            Amount: amt,
            Status: models.TransactionStatus(status),
            Desc: desc,
        }
        out = append(out, txn)
    }
    return out, nil
}
