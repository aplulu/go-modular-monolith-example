package connect

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/bufbuild/connect-go"

	"github.com/aplulu/modular-monolith-example-go/internal/component/user/usecase"
	appGRPC "github.com/aplulu/modular-monolith-example-go/internal/grpc"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
)

type internalServer struct {
	logger  *slog.Logger
	usecase usecase.UserUsecase
}

func (s *internalServer) GetUser(ctx context.Context, req *connect.Request[pbUser.GetUserRequest]) (*connect.Response[pbUser.GetUserResponse], error) {
	user, err := s.usecase.GetUserByID(ctx, req.Msg.UserId)
	if err != nil {
		err := fmt.Errorf("failed to get user: %w", err)
		s.logger.ErrorContext(ctx, err.Error())
		return nil, toGRPCError(err)
	}

	return connect.NewResponse(&pbUser.GetUserResponse{
		User: toPBUser(user),
	}), nil
}

func RegisterInternalServer(adapter appGRPC.ConnectUserAdapter, logger *slog.Logger, usecase usecase.UserUsecase) {
	s := &internalServer{
		logger:  logger,
		usecase: usecase,
	}

	adapter.Register(s)
}
