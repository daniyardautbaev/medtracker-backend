package services

import (
	"context"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/repositories"
)

// PatientService — сервис для работы с пациентами
type PatientService struct {
	repo repositories.PatientRepository
}

// NewPatientService создает новый сервис для работы с пациентами
func NewPatientService(repo repositories.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

// GetPatientDetails получает информацию о пациенте по ID
func (s *PatientService) GetPatientDetails(ctx context.Context, id string) (*data.PatientDTO, error) {
	return s.repo.GetPatientByID(ctx, id)
}

// CreatePatient создает нового пациента
func (s *PatientService) CreatePatient(ctx context.Context, patient data.PatientDTO) error {
	return s.repo.CreatePatient(ctx, patient)
}

// UpdatePatient обновляет информацию о пациенте
func (s *PatientService) UpdatePatient(ctx context.Context, id string, patient data.PatientDTO) error {
	return s.repo.UpdatePatient(ctx, id, patient)
}

// DeletePatient удаляет пациента по ID
func (s *PatientService) DeletePatient(ctx context.Context, id string) error {
	return s.repo.DeletePatient(ctx, id)
}
