package connect

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/aplulu/modular-monolith-example-go/internal/component/article/usecase"
	pbArticle "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/article/v1"
	"github.com/aplulu/modular-monolith-example-go/internal/grpc/example/article/v1/articleconnect"
)

type articleServer struct {
	logger  *slog.Logger
	usecase usecase.ArticleUsecase
}

func (s *articleServer) ListArticle(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[pbArticle.ListArticleResponse], error) {
	articles, err := s.usecase.ListArticle(ctx)
	if err != nil {
		err := fmt.Errorf("failed to list article: %w", err)
		s.logger.ErrorContext(ctx, err.Error())
		return nil, toGRPCError(err)
	}

	res := connect.NewResponse(&pbArticle.ListArticleResponse{
		Articles: toPBArticleList(articles),
	})

	// Getリクエストの場合はキャッシュを有効にする
	if req.HTTPMethod() == http.MethodGet {
		res.Header().Set("Cache-Control", "public, max-age=60")
	}

	return res, nil
}

func RegisterArticleServer(mux *http.ServeMux, logger *slog.Logger, usecase usecase.ArticleUsecase) {
	s := &articleServer{
		logger:  logger,
		usecase: usecase,
	}

	path, handler := articleconnect.NewArticleServiceHandler(s)
	mux.Handle(path, handler)
}
