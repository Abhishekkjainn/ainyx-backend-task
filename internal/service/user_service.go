package service

import (
	"ainyx-backend/db/sqlc"
	"ainyx-backend/internal/models"
	"ainyx-backend/internal/repository"
	"context"
	"errors"
	"time"
)

// Service Interface (Decoupling)
type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error)
	GetUser(ctx context.Context, id int32) (models.UserResponse, error)
	ListUsers(ctx context.Context, page, limit int32) ([]models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	// FIX: Do not ignore the error!
	parsedDob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	user, err := s.repo.CreateUser(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  parsedDob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return s.mapToResponse(user), nil
}

func (s *userService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}
	return s.mapToResponse(user), nil
}

func (s *userService) ListUsers(ctx context.Context, page, limit int32) ([]models.UserResponse, error) {
	if page < 1 { page = 1 }
	if limit < 1 { limit = 10 }
	offset := (page - 1) * limit

	// Pagination (Bonus retained from Abhishek)
	users, err := s.repo.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var response []models.UserResponse
	for _, u := range users {
		response = append(response, s.mapToResponse(u))
	}
	return response, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	parsedDob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	user, err := s.repo.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  parsedDob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return s.mapToResponse(user), nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

// Helper to avoid repetition
func (s *userService) mapToResponse(user sqlc.User) models.UserResponse {
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  models.CalculateAge(user.Dob),
	}
}