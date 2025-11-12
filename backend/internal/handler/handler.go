package handler

import (
	"log"
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
func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "failed to parse form: "+err.Error())
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "missing file: "+err.Error())
		return
	}
	defer file.Close()

	txns, err := utils.ParseCSV(file)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid CSV format: "+err.Error())
		return
	}
	if len(txns) == 0 {
		utils.RespondError(w, http.StatusBadRequest, "no valid transactions found in CSV")
		return
	}

	if err := h.svc.UploadTransactions(txns); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to store transactions: "+err.Error())
		return
	}

	log.Printf("[UPLOAD] %d transactions uploaded", len(txns))

	utils.RespondSuccess(w, "file uploaded successfully", map[string]interface{}{
		"count": len(txns),
	})
}

// GET /balance
func (h *Handler) Balance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	bal := h.svc.ComputeBalance()
	utils.RespondSuccess(w, "balance retrieved", map[string]interface{}{
		"balance": bal,
	})
}

// GET /issues
func (h *Handler) Issues(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	perPage, _ := strconv.Atoi(q.Get("per_page"))
	sortBy := q.Get("sort_by")
	order := strings.ToLower(q.Get("order"))

	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10
	}
	if order != "asc" && order != "desc" {
		order = "desc"
	}

	res, err := h.svc.ListIssues(page, perPage, sortBy, order)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "failed to list issues: "+err.Error())
		return
	}

	utils.RespondSuccess(w, "issues retrieved", res)
}
