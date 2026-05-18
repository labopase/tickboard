package repositories

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/tickboard/products/users/user/models"
)

type UserPostgreRepository interface {
	GetUserByID(ctx context.Context, id string) (*models.UserModel, error)
	CreateUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error)
	UpdateUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error)
	DeleteUser(ctx context.Context, id string) (*models.UserModel, error)
}
