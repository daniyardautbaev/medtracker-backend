package services

import (
	"context"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/repositories"
)

// MedicationService реализует бизнес-логику над лекарствами
type MedicationService struct {
	repo repositories.MedicationRepository
}

func NewMedicationService(repo repositories.MedicationRepository) *MedicationService {
	return &MedicationService{repo: repo}
}

func (s *MedicationService) CreateMedication(ctx context.Context, med data.MedicationDTO) error {
	return s.repo.CreateMedication(ctx, med)
}

func (s *MedicationService) GetMedicationByID(ctx context.Context, id string) (*data.MedicationDTO, error) {
	return s.repo.GetMedicationByID(ctx, id)
}

func (s *MedicationService) UpdateMedication(ctx context.Context, id string, med data.MedicationDTO) error {
	return s.repo.UpdateMedication(ctx, id, med)
}

func (s *MedicationService) DeleteMedication(ctx context.Context, id string) error {
	return s.repo.DeleteMedication(ctx, id)
}

func (s *MedicationService) ListMedicationsByUser(ctx context.Context, userID string) ([]data.MedicationDTO, error) {
	return s.repo.ListMedicationsByUser(ctx, userID)
}
