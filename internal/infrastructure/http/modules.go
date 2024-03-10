package http

import (
	"log/slog"
	"net/http"

	"google.golang.org/grpc"

	articleInMemory "github.com/aplulu/modular-monolith-example-go/internal/component/article/infrastructure/inmemory"
	articleConnect "github.com/aplulu/modular-monolith-example-go/internal/component/article/interface/connect"
	articleGrpc "github.com/aplulu/modular-monolith-example-go/internal/component/article/interface/grpc"
	articleUsecase "github.com/aplulu/modular-monolith-example-go/internal/component/article/usecase"
	userInMemory "github.com/aplulu/modular-monolith-example-go/internal/component/user/infrastructure/inmemory"
	userConnect "github.com/aplulu/modular-monolith-example-go/internal/component/user/interface/connect"
	userGrpc "github.com/aplulu/modular-monolith-example-go/internal/component/user/interface/grpc"
	userUsecase "github.com/aplulu/modular-monolith-example-go/internal/component/user/usecase"
	appGRPC "github.com/aplulu/modular-monolith-example-go/internal/grpc"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
	"github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1/userconnect"
)

// registerArticleModule registers the article module
func registerArticleModule(gs grpc.ServiceRegistrar, mux *http.ServeMux, logger *slog.Logger, internalUserClient pbUser.InternalUserServiceClient, connectInternalUserClient userconnect.InternalUserServiceClient) {
	usecase := articleUsecase.NewArticleUsecase(logger, articleInMemory.NewInMemoryArticleRepository(), internalUserClient, connectInternalUserClient)

	// gRPC Server
	articleGrpc.RegisterArticleServer(gs, logger, usecase)

	// Connect Server
	articleConnect.RegisterArticleServer(mux, logger, usecase)
}

// registerUserModule registers the user module
func registerUserModule(gs grpc.ServiceRegistrar, adapter appGRPC.ConnectUserAdapter, logger *slog.Logger) {
	usecase := userUsecase.NewUserUsecase(logger, userInMemory.NewInMemoryUserRepository())

	// gRPC Server
	userGrpc.RegisterInternalServer(gs, logger, usecase)

	// Connect Server
	userConnect.RegisterInternalServer(adapter, logger, usecase)
}
