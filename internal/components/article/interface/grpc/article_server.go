package grpc

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/aplulu/modular-monolith-example-go/internal/components/article/usecase"
	pbArticle "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/article/v1"
)

type articleServer struct {
	logger  *slog.Logger
	usecase usecase.ArticleUsecase
}

func (s *articleServer) ListArticle(ctx context.Context, empty *emptypb.Empty) (*pbArticle.ListArticleResponse, error) {
	articles, err := s.usecase.ListArticle(ctx)
	if err != nil {
		err := fmt.Errorf("failed to list article: %w", err)
		s.logger.ErrorContext(ctx, err.Error())
		return nil, toGRPCError(err)
	}

	return &pbArticle.ListArticleResponse{
		Articles: toPBArticleList(articles),
	}, nil
}

func RegisterArticleServer(gs grpc.ServiceRegistrar, logger *slog.Logger, usecase usecase.ArticleUsecase) {
	s := &articleServer{
		logger:  logger,
		usecase: usecase,
	}

	pbArticle.RegisterArticleServiceServer(gs, s)
}
