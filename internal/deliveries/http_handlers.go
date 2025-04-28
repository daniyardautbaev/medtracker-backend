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

// PatientHandler — обработчик для работы с пациентами
type PatientHandler struct {
	services *services.PatientService
}

// NewPatientHandler создает новый PatientHandler
func NewPatientHandler(services *services.PatientService) *PatientHandler {
	return &PatientHandler{services: services}
}

// GetPatientHandler обрабатывает запросы для получения информации о пациенте по ID
func (h *PatientHandler) GetPatientHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	patient, err := h.services.GetPatientDetails(r.Context(), id)
	if err != nil {
		http.Error(w, "Patient Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(patient)
}

// CreatePatientHandler обрабатывает запросы для создания нового пациента
func (h *PatientHandler) CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[POST] /patient")

	var patientData data.PatientDTO
	if err := json.NewDecoder(r.Body).Decode(&patientData); err != nil {
		log.Printf("Bad input: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if patientData.ID == "" {
		patientData.ID = uuid.New().String()
	}

	log.Printf("Saving patient: %+v", patientData)

	if err := h.services.CreatePatient(r.Context(), patientData); err != nil {
		log.Printf("❌ Failed to create patient: %v", err)
		http.Error(w, "Failed to create patient", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Patient created successfully")
	w.WriteHeader(http.StatusCreated)
}

// UpdatePatientHandler обрабатывает запросы для обновления информации о пациенте
func (h *PatientHandler) UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var patientData data.PatientDTO
	if err := json.NewDecoder(r.Body).Decode(&patientData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.services.UpdatePatient(r.Context(), id, patientData); err != nil {
		http.Error(w, "Failed to update patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeletePatientHandler обрабатывает запросы для удаления пациента
func (h *PatientHandler) DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.services.DeletePatient(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
