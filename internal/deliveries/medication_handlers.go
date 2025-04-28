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

// MedicationHandler — обработчик для лекарств
type MedicationHandler struct {
	service *services.MedicationService
}

func NewMedicationHandler(service *services.MedicationService) *MedicationHandler {
	return &MedicationHandler{service: service}
}

func (h *MedicationHandler) CreateMedication(w http.ResponseWriter, r *http.Request) {
	log.Println("[POST] /medication")
	var med data.MedicationDTO
	if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if med.ID == "" {
		med.ID = generateUUID()
	}

	if err := h.service.CreateMedication(r.Context(), med); err != nil {
		http.Error(w, "Failed to create medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MedicationHandler) GetMedication(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	med, err := h.service.GetMedicationByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Medication not found", http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(med)
}

func (h *MedicationHandler) UpdateMedication(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var med data.MedicationDTO
	if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateMedication(r.Context(), id, med); err != nil {
		http.Error(w, "Failed to update medication", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *MedicationHandler) DeleteMedication(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteMedication(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete medication", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *MedicationHandler) ListMedicationsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}
	meds, err := h.service.ListMedicationsByUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to fetch medications", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(meds)
}

// generateUUID можно вынести в utils
func generateUUID() string {
	// Можно использовать github.com/google/uuid
	return uuid.New().String()
}
