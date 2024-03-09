package grpc

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	"github.com/aplulu/modular-monolith-example-go/internal/components/user/usecase"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
)

type internalServer struct {
	logger  *slog.Logger
	usecase usecase.UserUsecase
}

func (s *internalServer) GetUser(ctx context.Context, pbReq *pbUser.GetUserRequest) (*pbUser.GetUserResponse, error) {
	user, err := s.usecase.GetUserByID(ctx, pbReq.UserId)
	if err != nil {
		err := fmt.Errorf("failed to get user: %w", err)
		s.logger.ErrorContext(ctx, err.Error())
		return nil, toGRPCError(err)
	}

	return &pbUser.GetUserResponse{
		User: toPBUser(user),
	}, nil
}

func RegisterInternalServer(gs grpc.ServiceRegistrar, logger *slog.Logger, usecase usecase.UserUsecase) {
	s := &internalServer{
		logger:  logger,
		usecase: usecase,
	}

	pbUser.RegisterInternalUserServiceServer(gs, s)
}
