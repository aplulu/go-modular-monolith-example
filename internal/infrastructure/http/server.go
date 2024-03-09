package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/aplulu/modular-monolith-example-go/internal/config"
	appGRPC "github.com/aplulu/modular-monolith-example-go/internal/grpc"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
)

var server http.Server

// StartServer starts the server
func StartServer(logger *slog.Logger) error {
	httpMux := chi.NewMux()

	grpcServer := grpc.NewServer()
	if config.GRPCReflectionService() {
		reflection.Register(grpcServer)
	}

	// Register Modules
	internalUserServiceAdapter := appGRPC.NewServiceAdapter()
	internalUserClient := pbUser.NewInternalUserServiceClient(internalUserServiceAdapter)

	registerUserModule(internalUserServiceAdapter, logger)

	registerArticleModule(grpcServer, logger, internalUserClient)

	server = http.Server{
		Addr: net.JoinHostPort(config.Listen(), config.Port()),
		Handler: h2c.NewHandler(
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// gRPC
				if req.ProtoMajor == 2 && strings.Contains(req.Header.Get("Content-Type"), "application/grpc") {
					grpcServer.ServeHTTP(w, req)
				} else {
					httpMux.ServeHTTP(w, req)
				}
			}),
			&http2.Server{},
		),
	}

	listenHost := config.Listen()
	if listenHost == "" {
		listenHost = "localhost"
	}
	logger.Info(fmt.Sprintf("Server started at http://%s:%s", listenHost, config.Port()))
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// StopServer stops the server gracefully
func StopServer(ctx context.Context) error {
	return server.Shutdown(ctx)
}
