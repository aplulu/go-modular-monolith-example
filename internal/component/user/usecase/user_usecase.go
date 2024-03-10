package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aplulu/modular-monolith-example-go/internal/component/user/domain/model"
	"github.com/aplulu/modular-monolith-example-go/internal/component/user/domain/repository"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

type userUsecase struct {
	logger         *slog.Logger
	userRepository repository.UserRepository
}

func (u *userUsecase) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read user: %w", err)
	}

	return user, nil
}

func NewUserUsecase(logger *slog.Logger, userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		logger:         logger,
		userRepository: userRepository,
	}
}
