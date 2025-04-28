package repositories

import (
	"context"
	"database/sql"
	"medtracker/medtracker/internal/data"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(ctx context.Context, user data.UserDTO) error
	GetUserByEmail(ctx context.Context, email string) (*data.UserDTO, error)
	GetUserByID(ctx context.Context, id string) (*data.UserDTO, error)
}

// MySQLUserRepository — реализация UserRepository для MySQL
type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) CreateUser(ctx context.Context, user data.UserDTO) error {
	query := `INSERT INTO users (id, name, email, password_hash, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt)
	return err
}

func (r *MySQLUserRepository) GetUserByEmail(ctx context.Context, email string) (*data.UserDTO, error) {
	query := `SELECT id, name, email, password_hash, created_at FROM users WHERE email = ?`
	row := r.db.QueryRowContext(ctx, query, email)

	var user data.UserDTO
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) GetUserByID(ctx context.Context, id string) (*data.UserDTO, error) {
	query := `SELECT id, name, email, password_hash, created_at FROM users WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var user data.UserDTO
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
