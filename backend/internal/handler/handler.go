package handler

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strconv"
	"strings"

    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/service"
    "github.com/bagindaisfa/flip-bank-statement-viewer/internal/utils"
)

type Handler struct {
    svc *service.Service
}

func NewHandler(s *service.Service) *Handler {
    return &Handler{svc: s}
}

// POST /upload
// form file field: file
func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20) // 10MB
    if err != nil {
        http.Error(w, "failed parse multipart form: "+err.Error(), http.StatusBadRequest)
        return
    }
    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "failed read file: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()
    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "failed read file: "+err.Error(), http.StatusInternalServerError)
        return
    }
    txns, err := utils.ParseCSV(strings.NewReader(string(data)))
    if err != nil {
        http.Error(w, "failed parse csv: "+err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.svc.UploadTransactions(txns); err != nil {
        http.Error(w, "upload failed: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, `{"status":"ok","count":`+strconv.Itoa(len(txns))+`}`)
}

func (h *Handler) Balance(w http.ResponseWriter, r *http.Request) {
    bal := h.svc.ComputeBalance()
    resp := map[string]interface{}{
        "balance": bal,
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Issues(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    page := 1
    perPage := 10
    sortBy := q.Get("sort_by")
    order := q.Get("order")
    if order != "asc" && order != "desc" {
        order = "desc"
    }
    if s := q.Get("page"); s != "" {
        if v, err := strconv.Atoi(s); err == nil && v > 0 {
            page = v
        }
    }
    if s := q.Get("per_page"); s != "" {
        if v, err := strconv.Atoi(s); err == nil && v > 0 {
            perPage = v
        }
    }

    res, err := h.svc.ListIssues(page, perPage, sortBy, order)
    if err != nil {
        http.Error(w, "failed list issues: "+err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(res)
}
