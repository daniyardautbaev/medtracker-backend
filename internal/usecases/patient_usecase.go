package usecases

import (
	"context"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/services"
)

// PatientUseCase — сценарий работы с пациентами
type PatientUseCase struct {
	service *services.PatientService
}

// NewPatientUseCase создает новый UseCase для работы с пациентами
func NewPatientUseCase(service *services.PatientService) *PatientUseCase {
	return &PatientUseCase{service: service}
}

// Execute получает информацию о пациенте по ID
func (u *PatientUseCase) Execute(ctx context.Context, id string) (*data.PatientDTO, error) {
	return u.service.GetPatientDetails(ctx, id)
}

// CreatePatient создает нового пациента
func (u *PatientUseCase) CreatePatient(ctx context.Context, patient data.PatientDTO) error {
	return u.service.CreatePatient(ctx, patient)
}

// UpdatePatient обновляет информацию о пациенте
func (u *PatientUseCase) UpdatePatient(ctx context.Context, id string, patient data.PatientDTO) error {
	return u.service.UpdatePatient(ctx, id, patient)
}

// DeletePatient удаляет пациента
func (u *PatientUseCase) DeletePatient(ctx context.Context, id string) error {
	return u.service.DeletePatient(ctx, id)
}
