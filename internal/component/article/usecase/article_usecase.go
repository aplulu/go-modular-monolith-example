package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/bufbuild/connect-go"

	"github.com/aplulu/modular-monolith-example-go/internal/component/article/domain/model"
	"github.com/aplulu/modular-monolith-example-go/internal/component/article/domain/repository"
	"github.com/aplulu/modular-monolith-example-go/internal/config"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
	"github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1/userconnect"
)

type ArticleUsecase interface {
	ListArticle(ctx context.Context) ([]*model.Article, error)
}

type articleUsecase struct {
	logger            *slog.Logger
	articleRepository repository.ArticleRepository
	// InternalUserServiceClient (gRPC)
	internalUserClient pbUser.InternalUserServiceClient
	// InternalUserServiceClient (Connect)
	connectInternalUserClient userconnect.InternalUserServiceClient
}

// ListArticle 記事一覧を取得
func (u *articleUsecase) ListArticle(ctx context.Context) ([]*model.Article, error) {
	// 記事リポジトリから記事一覧を取得
	articles, err := u.articleRepository.ListArticle(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list article: %w", err)
	}

	for _, article := range articles {
		var user *pbUser.User
		// Userモジュールからユーザー情報を取得
		if config.InternalProtocol() == "connect" { // Connect
			res, err := u.connectInternalUserClient.GetUser(ctx, connect.NewRequest(&pbUser.GetUserRequest{UserId: article.UserID}))
			if err != nil {
				return nil, fmt.Errorf("failed to get user: %w", err)
			}
			user = res.Msg.User
		} else { // gRPC
			res, err := u.internalUserClient.GetUser(ctx, &pbUser.GetUserRequest{UserId: article.UserID})
			if err != nil {
				return nil, fmt.Errorf("failed to get user: %w", err)
			}
			user = res.User
		}

		article.User = &model.ArticleUser{
			ID:   user.Id,
			Name: user.Name,
		}
	}

	return articles, nil
}

func NewArticleUsecase(logger *slog.Logger, articleRepository repository.ArticleRepository, internalUserClient pbUser.InternalUserServiceClient, connectInternalUserClient userconnect.InternalUserServiceClient) ArticleUsecase {
	return &articleUsecase{
		logger:                    logger,
		articleRepository:         articleRepository,
		internalUserClient:        internalUserClient,
		connectInternalUserClient: connectInternalUserClient,
	}
}
