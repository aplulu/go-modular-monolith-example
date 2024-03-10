package repository

import (
	"context"

	"github.com/aplulu/modular-monolith-example-go/internal/component/article/domain/model"
)

type ArticleRepository interface {
	ListArticle(ctx context.Context) ([]*model.Article, error)
}
