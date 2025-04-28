package repositories

import (
	"context"
	"database/sql"
	"medtracker/medtracker/internal/data"
)

// MedicationRepository интерфейс для работы с лекарствами
type MedicationRepository interface {
	GetMedicationByID(ctx context.Context, id string) (*data.MedicationDTO, error)
	CreateMedication(ctx context.Context, medication data.MedicationDTO) error
	UpdateMedication(ctx context.Context, id string, medication data.MedicationDTO) error
	DeleteMedication(ctx context.Context, id string) error
	ListMedicationsByUser(ctx context.Context, userID string) ([]data.MedicationDTO, error)
}

// MySQLMedicationRepository реализация MedicationRepository для MySQL
type MySQLMedicationRepository struct {
	db *sql.DB
}

func NewMySQLMedicationRepository(db *sql.DB) *MySQLMedicationRepository {
	return &MySQLMedicationRepository{db: db}
}

func (r *MySQLMedicationRepository) GetMedicationByID(ctx context.Context, id string) (*data.MedicationDTO, error) {
	query := "SELECT id, user_id, name, dosage, frequency, start_date, end_date FROM medications WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var med data.MedicationDTO
	err := row.Scan(&med.ID, &med.UserID, &med.Name, &med.Dosage, &med.Frequency, &med.StartDate, &med.EndDate)
	if err != nil {
		return nil, err
	}
	return &med, nil
}

func (r *MySQLMedicationRepository) CreateMedication(ctx context.Context, med data.MedicationDTO) error {
	query := "INSERT INTO medications (id, user_id, name, dosage, frequency, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, med.ID, med.UserID, med.Name, med.Dosage, med.Frequency, med.StartDate, med.EndDate)
	return err
}

func (r *MySQLMedicationRepository) UpdateMedication(ctx context.Context, id string, med data.MedicationDTO) error {
	query := "UPDATE medications SET name = ?, dosage = ?, frequency = ?, start_date = ?, end_date = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, med.Name, med.Dosage, med.Frequency, med.StartDate, med.EndDate, id)
	return err
}

func (r *MySQLMedicationRepository) DeleteMedication(ctx context.Context, id string) error {
	query := "DELETE FROM medications WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *MySQLMedicationRepository) ListMedicationsByUser(ctx context.Context, userID string) ([]data.MedicationDTO, error) {
	query := "SELECT id, user_id, name, dosage, frequency, start_date, end_date FROM medications WHERE user_id = ?"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medications []data.MedicationDTO
	for rows.Next() {
		var med data.MedicationDTO
		err := rows.Scan(&med.ID, &med.UserID, &med.Name, &med.Dosage, &med.Frequency, &med.StartDate, &med.EndDate)
		if err != nil {
			return nil, err
		}
		medications = append(medications, med)
	}
	return medications, nil
}
