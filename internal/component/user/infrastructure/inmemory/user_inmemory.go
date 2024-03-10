package inmemory

import (
	"context"
	"fmt"

	"github.com/aplulu/modular-monolith-example-go/internal/component/user/domain/model"
	"github.com/aplulu/modular-monolith-example-go/internal/component/user/domain/repository"
)

var users = []*model.User{{
	ID:   "1",
	Name: "User 1",
}, {
	ID:   "2",
	Name: "User 2",
}}

type inMemoryUserRepository struct{}

func (r *inMemoryUserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with id %s not found: %w", id, model.ErrNotFound)
}

func NewInMemoryUserRepository() repository.UserRepository {
	return &inMemoryUserRepository{}
}
