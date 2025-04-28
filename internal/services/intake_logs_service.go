package services

import (
	"context"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/repositories"
)

type IntakeLogService struct {
	repo repositories.IntakeLogRepository
}

func NewIntakeLogService(repo repositories.IntakeLogRepository) *IntakeLogService {
	return &IntakeLogService{repo: repo}
}

func (s *IntakeLogService) CreateLog(ctx context.Context, log data.IntakeLogDTO) error {
	return s.repo.CreateLog(ctx, log)
}

func (s *IntakeLogService) GetLogByID(ctx context.Context, id string) (*data.IntakeLogDTO, error) {
	return s.repo.GetLogByID(ctx, id)
}

func (s *IntakeLogService) DeleteLog(ctx context.Context, id string) error {
	return s.repo.DeleteLog(ctx, id)
}

func (s *IntakeLogService) ListLogsByMedication(ctx context.Context, medicationID string) ([]data.IntakeLogDTO, error) {
	return s.repo.ListLogsByMedication(ctx, medicationID)
}

func (s *IntakeLogService) ListLogsByMedicationName(ctx context.Context, medicationName string) ([]data.IntakeLogDTO, error) {
	return s.repo.ListLogsByMedicationName(ctx, medicationName)
}
