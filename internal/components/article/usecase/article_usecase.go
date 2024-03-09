package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aplulu/modular-monolith-example-go/internal/components/article/domain/model"
	"github.com/aplulu/modular-monolith-example-go/internal/components/article/domain/repository"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
)

type ArticleUsecase interface {
	ListArticle(ctx context.Context) ([]*model.Article, error)
}

type articleUsecase struct {
	logger             *slog.Logger
	articleRepository  repository.ArticleRepository
	internalUserClient pbUser.InternalUserServiceClient
}

// ListArticle 記事一覧を取得
func (u *articleUsecase) ListArticle(ctx context.Context) ([]*model.Article, error) {
	// 記事リポジトリから記事一覧を取得
	articles, err := u.articleRepository.ListArticle(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list article: %w", err)
	}

	for _, article := range articles {
		// Userモジュールからユーザー情報を取得
		user, err := u.internalUserClient.GetUser(ctx, &pbUser.GetUserRequest{UserId: article.UserID})
		if err != nil {
			return nil, fmt.Errorf("failed to get user: %w", err)
		}

		article.User = &model.ArticleUser{
			ID:   user.User.Id,
			Name: user.User.Name,
		}
	}

	return articles, nil
}

func NewArticleUsecase(logger *slog.Logger, articleRepository repository.ArticleRepository, internalUserClient pbUser.InternalUserServiceClient) ArticleUsecase {
	return &articleUsecase{
		logger:             logger,
		articleRepository:  articleRepository,
		internalUserClient: internalUserClient,
	}
}
