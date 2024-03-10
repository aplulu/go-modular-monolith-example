package repository

import (
	"context"

	"github.com/aplulu/modular-monolith-example-go/internal/component/user/domain/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}
