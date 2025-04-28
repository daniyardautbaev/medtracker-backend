package deliveries

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/services"
	"net/http"
)

type IntakeLogHandler struct {
	service *services.IntakeLogService
}

func NewIntakeLogHandler(service *services.IntakeLogService) *IntakeLogHandler {
	return &IntakeLogHandler{service: service}
}

func (h *IntakeLogHandler) CreateLog(w http.ResponseWriter, r *http.Request) {
	var log data.IntakeLogDTO
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	log.ID = uuid.New().String()
	if err := h.service.CreateLog(r.Context(), log); err != nil {
		http.Error(w, "Failed to create log", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *IntakeLogHandler) GetLog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logEntry, err := h.service.GetLogByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Log not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(logEntry)
}

func (h *IntakeLogHandler) DeleteLog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteLog(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete log", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *IntakeLogHandler) ListLogsByMedication(w http.ResponseWriter, r *http.Request) {
	medicationID := r.URL.Query().Get("medication_id")
	medicationName := r.URL.Query().Get("medication_name")

	var (
		logs []data.IntakeLogDTO
		err  error
	)

	if medicationName != "" {
		logs, err = h.service.ListLogsByMedicationName(r.Context(), medicationName)
	} else if medicationID != "" {
		logs, err = h.service.ListLogsByMedication(r.Context(), medicationID)
	} else {
		http.Error(w, "Missing medication_id or medication_name", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to fetch logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(logs)
}
