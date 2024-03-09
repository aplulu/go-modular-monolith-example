package http

import (
	"log/slog"

	"google.golang.org/grpc"

	articleInMemory "github.com/aplulu/modular-monolith-example-go/internal/components/article/infrastructure/inmemory"
	articleGrpc "github.com/aplulu/modular-monolith-example-go/internal/components/article/interface/grpc"
	articleUsecase "github.com/aplulu/modular-monolith-example-go/internal/components/article/usecase"
	userInMemory "github.com/aplulu/modular-monolith-example-go/internal/components/user/infrastructure/inmemory"
	userGrpc "github.com/aplulu/modular-monolith-example-go/internal/components/user/interface/grpc"
	userUsecase "github.com/aplulu/modular-monolith-example-go/internal/components/user/usecase"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
)

// registerArticleModule registers the article module
func registerArticleModule(gs grpc.ServiceRegistrar, logger *slog.Logger, internalUserClient pbUser.InternalUserServiceClient) {
	usecase := articleUsecase.NewArticleUsecase(logger, articleInMemory.NewInMemoryArticleRepository(), internalUserClient)

	articleGrpc.RegisterArticleServer(gs, logger, usecase)
}

func registerUserModule(gs grpc.ServiceRegistrar, logger *slog.Logger) {
	usecase := userUsecase.NewUserUsecase(logger, userInMemory.NewInMemoryUserRepository())

	userGrpc.RegisterInternalServer(gs, logger, usecase)
}
