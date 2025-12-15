package repository

import (
	"ainyx-backend/db/sqlc"
	"context"
)


type UserRepository interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUser(ctx context.Context, id int32) (sqlc.User, error)
	ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error)
	UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type SQLUserRepository struct {
	*sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &SQLUserRepository{
		Queries: db,
	}
}
