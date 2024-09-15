package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/gdwr/centric/internal/auth"
	"github.com/gdwr/centric/internal/database"
)

type UserService struct {
	authService     *auth.AuthService
	databaseService *database.DatabaseService
}

func NewUserService(a *auth.AuthService, d *database.DatabaseService) *UserService {
	return &UserService{
		authService:     a,
		databaseService: d,
	}
}

func (u UserService) CreateUser(username, password string) (*database.User, error) {
	return u.CreateUserWithCtx(context.Background(), username, password)
}

func (u UserService) CreateUserWithCtx(ctx context.Context, username, password string) (*database.User, error) {
	salt, err := u.authService.GenerateSalt()
	if err != nil {
		return nil, err
	}

	hash := u.authService.HashPassword(password, salt)
	user, err := u.databaseService.CreateUser(ctx, database.CreateUserParams{
		ID:       uuid.New().String(),
		Name:     username,
		Password: hash,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
