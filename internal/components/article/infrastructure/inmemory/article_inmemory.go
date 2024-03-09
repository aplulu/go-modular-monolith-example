package inmemory

import (
	"context"

	"github.com/aplulu/modular-monolith-example-go/internal/components/article/domain/model"
	"github.com/aplulu/modular-monolith-example-go/internal/components/article/domain/repository"
)

var articles = []*model.Article{{
	ID:      "1",
	Title:   "Title 1",
	Content: "Content 1",
	UserID:  "1",
}, {
	ID:      "2",
	Title:   "Title 2",
	Content: "Content 2",
	UserID:  "2",
}}

type inMemoryArticleRepository struct{}

func (r *inMemoryArticleRepository) ListArticle(ctx context.Context) ([]*model.Article, error) {
	return articles, nil
}

func NewInMemoryArticleRepository() repository.ArticleRepository {
	return &inMemoryArticleRepository{}
}
