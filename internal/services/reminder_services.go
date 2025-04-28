package services

import (
	"context"
	"github.com/google/uuid"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/repositories"
	"time"
)

// ReminderService реализует бизнес-логику для напоминаний
type ReminderService struct {
	repo       repositories.ReminderRepository
	intakeRepo repositories.IntakeLogRepository
}

func NewReminderService(repo repositories.ReminderRepository, intakeRepo repositories.IntakeLogRepository) *ReminderService {
	return &ReminderService{
		repo:       repo,
		intakeRepo: intakeRepo,
	}
}
func (s *ReminderService) CreateReminder(ctx context.Context, r data.ReminderDTO) error {
	return s.repo.CreateReminder(ctx, r)
}

func (s *ReminderService) GetReminderByID(ctx context.Context, id string) (*data.ReminderDTO, error) {
	return s.repo.GetReminderByID(ctx, id)
}

func (s *ReminderService) UpdateReminder(ctx context.Context, id string, r data.ReminderDTO) error {
	return s.repo.UpdateReminder(ctx, id, r)
}
func (s *ReminderService) ListRemindersByUser(ctx context.Context, userID string) ([]data.ReminderDTO, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *ReminderService) DeleteReminder(ctx context.Context, id string) error {
	return s.repo.DeleteReminder(ctx, id)
}

func (s *ReminderService) ListRemindersByMedication(ctx context.Context, medicationID string) ([]data.ReminderDTO, error) {
	return s.repo.ListRemindersByMedication(ctx, medicationID)
}

func (s *ReminderService) ConfirmReminder(ctx context.Context, reminderID string) error {
	// 1. Сначала найдём напоминание по ID
	reminder, err := s.repo.GetReminderByID(ctx, reminderID)
	if err != nil {
		return err
	}

	// 2. Генерируем лог
	now := time.Now().Format("2006-01-02 15:04:05") // MySQL формат

	log := data.IntakeLogDTO{
		ID:           uuid.NewString(),
		MedicationID: reminder.MedicationID, // <--- правильный MedicationID
		IntakeTime:   now,
		Taken:        true,
	}

	// 3. Сохраняем лог
	return s.intakeRepo.CreateLog(ctx, log)
}
