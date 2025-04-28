package services

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"medtracker/medtracker/internal/data"
	"medtracker/medtracker/internal/repositories"
)

// UserService реализует бизнес-логику для пользователей
type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// RegisterUser регистрирует нового пользователя с хешированием пароля
func (s *UserService) RegisterUser(ctx context.Context, user data.UserDTO, plainPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)
	return s.repo.CreateUser(ctx, user)
}

// AuthenticateUser проверяет email + пароль
func (s *UserService) AuthenticateUser(ctx context.Context, email string, password string) (*data.UserDTO, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
