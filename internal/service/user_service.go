package service

import (
	"ainyx-backend/db/sqlc"
	"ainyx-backend/internal/models"
	"context"
	"time"
)

type UserService struct {
	queries *sqlc.Queries
}

func NewUserService(db *sqlc.Queries) *UserService {
	return &UserService{queries: db}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	user, err := s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  parsedDob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  models.CalculateAge(user.Dob),
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.queries.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  models.CalculateAge(user.Dob),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, page, limit int32) ([]models.UserResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	users, err := s.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var response []models.UserResponse
	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Format("2006-01-02"),
			Age:  models.CalculateAge(u.Dob),
		})
	}
	return response, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	user, err := s.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  parsedDob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  models.CalculateAge(user.Dob),
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.queries.DeleteUser(ctx, id)
}