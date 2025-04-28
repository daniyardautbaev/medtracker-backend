package repositories

import (
	"context"
	"database/sql"
	"medtracker/medtracker/internal/data"
)

// PatientRepository интерфейс для работы с пациентами
type PatientRepository interface {
	GetPatientByID(ctx context.Context, id string) (*data.PatientDTO, error)
	CreatePatient(ctx context.Context, patient data.PatientDTO) error
	UpdatePatient(ctx context.Context, id string, patient data.PatientDTO) error
	DeletePatient(ctx context.Context, id string) error
}

// MySQLPatientRepository репозиторий для MySQL
type MySQLPatientRepository struct {
	db *sql.DB
}

// NewMySQLPatientRepository создает новый репозиторий для MySQL
func NewMySQLPatientRepository(db *sql.DB) *MySQLPatientRepository {
	return &MySQLPatientRepository{db: db}
}

// GetPatientByID получает пациента по ID из MySQL
func (r *MySQLPatientRepository) GetPatientByID(ctx context.Context, id string) (*data.PatientDTO, error) {
	query := "SELECT id, name, age, `condition` FROM patients WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var patient data.PatientDTO
	err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.Condition)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

// CreatePatient добавляет нового пациента в MySQL
func (r *MySQLPatientRepository) CreatePatient(ctx context.Context, patient data.PatientDTO) error {
	query := "INSERT INTO patients (id, name, age, `condition`) VALUES (?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, patient.ID, patient.Name, patient.Age, patient.Condition)
	return err
}

// UpdatePatient обновляет информацию о пациенте в MySQL
func (r *MySQLPatientRepository) UpdatePatient(ctx context.Context, id string, patient data.PatientDTO) error {
	query := "UPDATE patients SET name = ?, age = ?, `condition` = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, patient.Name, patient.Age, patient.Condition, id)
	return err
}

// DeletePatient удаляет пациента из MySQL
func (r *MySQLPatientRepository) DeletePatient(ctx context.Context, id string) error {
	query := `DELETE FROM patients WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
