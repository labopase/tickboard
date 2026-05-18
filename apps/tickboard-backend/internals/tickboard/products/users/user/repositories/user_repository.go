package repositories

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/pgsql"
	"github.com/halimdotnet/tickboard-backend/internals/tickboard/products/users/user/models"
)

type userPostgreRepository struct {
	pool pgsql.Client
	log  logger.Logger
}

func NewUserRepository(client pgsql.Client, log logger.Logger) UserPostgreRepository {
	return &userPostgreRepository{pool: client, log: log}
}

func (r *userPostgreRepository) GetUserByID(ctx context.Context, id string) (*models.UserModel, error) {
	return &models.UserModel{
		ID:          "AB120",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "[EMAIL_ADDRESS]",
		PhoneArea:   "ID",
		PhoneNumber: "81234567890",
	}, nil
}

func (r *userPostgreRepository) CreateUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error) {
	return nil, nil
}

func (r *userPostgreRepository) UpdateUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error) {
	return nil, nil
}

func (r *userPostgreRepository) DeleteUser(ctx context.Context, id string) (*models.UserModel, error) {
	return nil, nil
}
