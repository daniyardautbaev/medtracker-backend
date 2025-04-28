package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"medtracker/medtracker/internal/data"
)

// IntakeLogRepository интерфейс для работы с логами приёма
type IntakeLogRepository interface {
	CreateLog(ctx context.Context, log data.IntakeLogDTO) error
	GetLogByID(ctx context.Context, id string) (*data.IntakeLogDTO, error)
	DeleteLog(ctx context.Context, id string) error
	ListLogsByMedication(ctx context.Context, medicationID string) ([]data.IntakeLogDTO, error)
	ListLogsByMedicationName(ctx context.Context, medicationName string) ([]data.IntakeLogDTO, error) // ✅ новый метод
}

// MySQLIntakeLogRepository реализация репозитория для MySQL
type MySQLIntakeLogRepository struct {
	db *sql.DB
}

// NewMySQLIntakeLogRepository создает новый репозиторий для MySQL
func NewMySQLIntakeLogRepository(db *sql.DB) *MySQLIntakeLogRepository {
	return &MySQLIntakeLogRepository{db: db}
}

// CreateLog добавляет новый лог приёма в MySQL
func (r *MySQLIntakeLogRepository) CreateLog(ctx context.Context, log data.IntakeLogDTO) error {
	query := `INSERT INTO intake_logs (id, medication_id, taken_at, was_taken) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, log.ID, log.MedicationID, log.IntakeTime, log.Taken)
	if err != nil {
		fmt.Printf("Error executing query: %v, error: %v\n", query, err)
		return err
	}
	return nil
}

// GetLogByID получает лог приёма по ID
func (r *MySQLIntakeLogRepository) GetLogByID(ctx context.Context, id string) (*data.IntakeLogDTO, error) {
	query := `SELECT id, medication_id, taken_at, was_taken FROM intake_logs WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var log data.IntakeLogDTO
	err := row.Scan(&log.ID, &log.MedicationID, &log.IntakeTime, &log.Taken)
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// DeleteLog удаляет лог приёма
func (r *MySQLIntakeLogRepository) DeleteLog(ctx context.Context, id string) error {
	query := `DELETE FROM intake_logs WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// ListLogsByMedication получает список логов по medication_id
func (r *MySQLIntakeLogRepository) ListLogsByMedication(ctx context.Context, medicationID string) ([]data.IntakeLogDTO, error) {
	query := `SELECT id, medication_id, taken_at, was_taken FROM intake_logs WHERE medication_id = ?`
	rows, err := r.db.QueryContext(ctx, query, medicationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []data.IntakeLogDTO
	for rows.Next() {
		var log data.IntakeLogDTO
		if err := rows.Scan(&log.ID, &log.MedicationID, &log.IntakeTime, &log.Taken); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

// ListLogsByMedicationName получает список логов приёма по части названия лекарства без учёта регистра
func (r *MySQLIntakeLogRepository) ListLogsByMedicationName(ctx context.Context, medicationName string) ([]data.IntakeLogDTO, error) {
	query := `
	SELECT il.id, il.medication_id, il.taken_at, il.was_taken
	FROM intake_logs il
	JOIN medications m ON il.medication_id = m.id
	WHERE LOWER(m.name) LIKE LOWER(CONCAT('%', ?, '%'))
	`

	rows, err := r.db.QueryContext(ctx, query, medicationName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []data.IntakeLogDTO
	for rows.Next() {
		var log data.IntakeLogDTO
		if err := rows.Scan(&log.ID, &log.MedicationID, &log.IntakeTime, &log.Taken); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
