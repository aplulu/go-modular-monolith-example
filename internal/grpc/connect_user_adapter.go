package grpc

import (
	"context"

	"github.com/bufbuild/connect-go"

	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
	"github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1/userconnect"
)

type ConnectUserAdapter interface {
	Register(handler userconnect.InternalUserServiceHandler)

	GetUser(ctx context.Context, req *connect.Request[pbUser.GetUserRequest]) (*connect.Response[pbUser.GetUserResponse], error)
}

type connectUserAdapter struct {
	handler userconnect.InternalUserServiceHandler
}

func (a *connectUserAdapter) Register(handler userconnect.InternalUserServiceHandler) {
	a.handler = handler
}

func (a *connectUserAdapter) GetUser(ctx context.Context, req *connect.Request[pbUser.GetUserRequest]) (*connect.Response[pbUser.GetUserResponse], error) {
	return a.handler.GetUser(ctx, req)
}

func NewInternalUserAdapter() ConnectUserAdapter {
	return &connectUserAdapter{}
}
