package deliveries

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/services"
	"net/http"
)

// ReminderHandler обрабатывает HTTP-запросы для напоминаний
type ReminderHandler struct {
	service *services.ReminderService
}

func NewReminderHandler(service *services.ReminderService) *ReminderHandler {
	return &ReminderHandler{service: service}
}

func (h *ReminderHandler) CreateReminder(w http.ResponseWriter, r *http.Request) {
	var reminder data.ReminderDTO
	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if reminder.ID == "" {
		reminder.ID = uuid.New().String()
	}
	if err := h.service.CreateReminder(r.Context(), reminder); err != nil {
		http.Error(w, "Failed to create reminder", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ReminderHandler) GetReminder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	rem, err := h.service.GetReminderByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Reminder not found", http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(rem)
}

func (h *ReminderHandler) UpdateReminder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var reminder data.ReminderDTO
	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateReminder(r.Context(), id, reminder); err != nil {
		http.Error(w, "Failed to update reminder", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ReminderHandler) DeleteReminder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteReminder(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete reminder", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *ReminderHandler) ListRemindersByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	log.Println("👉 /reminders called with user_id =", userID)

	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	// Проверка на корректный UUID
	if _, err := uuid.Parse(userID); err != nil {
		log.Println("❌ Invalid UUID:", err)
		http.Error(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}

	reminders, err := h.service.ListRemindersByUser(r.Context(), userID)
	if err != nil {
		log.Println("❌ Failed to get reminders:", err)
		http.Error(w, "Failed to get reminders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(reminders)
}

func (h *ReminderHandler) ListRemindersByMedication(w http.ResponseWriter, r *http.Request) {
	medicationID := r.URL.Query().Get("medication_id")
	if medicationID == "" {
		http.Error(w, "Missing medication_id", http.StatusBadRequest)
		return
	}
	reminders, err := h.service.ListRemindersByMedication(r.Context(), medicationID)
	if err != nil {
		http.Error(w, "Failed to get reminders", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(reminders)
}

func (h *ReminderHandler) ConfirmReminder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.service.ConfirmReminder(r.Context(), id); err != nil {
		http.Error(w, "Failed to confirm reminder", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
