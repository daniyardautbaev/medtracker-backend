package repositories

import (
	"context"
	"database/sql"
	"log"
	"medtracker/medtracker/internal/data"
)

// ReminderRepository интерфейс для работы с напоминаниями
type ReminderRepository interface {
	CreateReminder(ctx context.Context, r data.ReminderDTO) error
	GetReminderByID(ctx context.Context, id string) (*data.ReminderDTO, error)
	UpdateReminder(ctx context.Context, id string, r data.ReminderDTO) error
	DeleteReminder(ctx context.Context, id string) error
	ListRemindersByMedication(ctx context.Context, medicationID string) ([]data.ReminderDTO, error)
	FindByUserID(ctx context.Context, userID string) ([]data.ReminderDTO, error)
}

// MySQLReminderRepository реализация репозитория для MySQL
type MySQLReminderRepository struct {
	db *sql.DB
}

func NewMySQLReminderRepository(db *sql.DB) *MySQLReminderRepository {
	return &MySQLReminderRepository{db: db}
}

// CreateReminder добавляет новое напоминание в MySQL
func (r *MySQLReminderRepository) CreateReminder(ctx context.Context, rem data.ReminderDTO) error {
	queryGetName := `SELECT name FROM medications WHERE id = ?`
	row := r.db.QueryRowContext(ctx, queryGetName, rem.MedicationID)
	if err := row.Scan(&rem.MedicationName); err != nil {
		log.Printf("❌ Ошибка при получении названия лекарства для ID %s: %v", rem.MedicationID, err)
		return err
	}

	query := `INSERT INTO reminders (id, medication_id,time,frequency, medication_name, user_id)
	          VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, rem.ID, rem.MedicationID, rem.Time, rem.Frequency, rem.MedicationName, rem.UserID)
	if err != nil {
		log.Printf("❌ Ошибка вставки напоминания: %v", err)
		return err
	}
	log.Printf("✅ Напоминание добавлено для пользователя %s: %s", rem.UserID, rem.MedicationName)
	return nil
}

// GetReminderByID получает напоминание по ID из MySQL
func (r *MySQLReminderRepository) GetReminderByID(ctx context.Context, id string) (*data.ReminderDTO, error) {
	query := `SELECT id, medication_id, time, frequency, medication_name FROM reminders WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	var rem data.ReminderDTO
	if err := row.Scan(&rem.ID, &rem.MedicationID, &rem.Time, &rem.Frequency, &rem.MedicationName); err != nil {
		log.Printf("❌ Ошибка при получении напоминания по ID %s: %v", id, err)
		return nil, err
	}
	log.Printf("✅ Напоминание получено для ID %s: %v", id, rem)
	return &rem, nil
}

// UpdateReminder обновляет напоминание в MySQL
func (r *MySQLReminderRepository) UpdateReminder(ctx context.Context, id string, rem data.ReminderDTO) error {
	query := `UPDATE reminders SET medication_id = ?, time = ?, frequency = ?, medication_name = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, rem.MedicationID, rem.Time, rem.Frequency, rem.MedicationName, id)
	if err != nil {
		log.Printf("❌ Ошибка обновления напоминания по ID %s: %v", id, err)
		return err
	}
	log.Printf("✅ Напоминание обновлено для ID %s", id)
	return nil
}

// FindByUserID возвращает напоминания по user_id
func (r *MySQLReminderRepository) FindByUserID(ctx context.Context, userID string) ([]data.ReminderDTO, error) {
	log.Println("✅ Запрос на получение напоминаний для user_id:", userID)

	query := `SELECT id, medication_id, time, frequency, medication_name FROM reminders WHERE user_id = ?`
	log.Println("✅ Выполняется запрос:", query)

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("❌ Ошибка при запросе для user_id %s: %v", userID, err)
		return nil, err
	}
	defer rows.Close()

	var reminders []data.ReminderDTO
	for rows.Next() {
		var rem data.ReminderDTO
		if err := rows.Scan(&rem.ID, &rem.MedicationID, &rem.Time, &rem.Frequency, &rem.MedicationName); err != nil {
			log.Printf("❌ Ошибка при сканировании данных для user_id %s: %v", userID, err)
			return nil, err
		}
		reminders = append(reminders, rem)
	}

	log.Printf("✅ Напоминания получены для user_id %s: %v", userID, reminders)
	return reminders, nil
}

// DeleteReminder удаляет напоминание из MySQL
func (r *MySQLReminderRepository) DeleteReminder(ctx context.Context, id string) error {
	query := `DELETE FROM reminders WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("❌ Ошибка при удалении напоминания по ID %s: %v", id, err)
		return err
	}
	log.Printf("✅ Напоминание удалено для ID %s", id)
	return nil
}

// ListRemindersByMedication получает список напоминаний по medication_id
func (r *MySQLReminderRepository) ListRemindersByMedication(ctx context.Context, medicationID string) ([]data.ReminderDTO, error) {
	query := `SELECT id, medication_id, time, frequency, medication_name FROM reminders WHERE medication_id = ?`
	rows, err := r.db.QueryContext(ctx, query, medicationID)
	if err != nil {
		log.Printf("❌ Ошибка при запросе по medication_id %s: %v", medicationID, err)
		return nil, err
	}
	defer rows.Close()

	var reminders []data.ReminderDTO
	for rows.Next() {
		var rem data.ReminderDTO
		if err := rows.Scan(&rem.ID, &rem.MedicationID, &rem.Time, &rem.Frequency, &rem.MedicationName); err != nil {
			log.Printf("❌ Ошибка при сканировании данных по medication_id %s: %v", medicationID, err)
			return nil, err
		}
		reminders = append(reminders, rem)
	}

	log.Printf("✅ Напоминания по medication_id %s получены: %v", medicationID, reminders)
	return reminders, nil
}
