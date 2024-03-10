package connect

import (
	"github.com/aplulu/modular-monolith-example-go/internal/component/article/domain/model"
	pbArticle "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/article/v1"
)

// toPBArticleList 記事リストをgRPCの記事リストに変換
func toPBArticleList(articles []*model.Article) []*pbArticle.Article {
	var pbArticles []*pbArticle.Article
	for _, article := range articles {
		pbArticles = append(pbArticles, toPBArticle(article))
	}

	return pbArticles
}

// toPBArticle 記事をgRPCの記事に変換
func toPBArticle(article *model.Article) *pbArticle.Article {
	return &pbArticle.Article{
		Id:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		User:    toPBArticleUser(article.User),
	}
}

// toPBArticleUser ユーザーをgRPCのユーザーに変換
func toPBArticleUser(user *model.ArticleUser) *pbArticle.ArticleUser {
	return &pbArticle.ArticleUser{
		Id:   user.ID,
		Name: user.Name,
	}
}

// toGRPCError gRPCのステータスコードに変換
func toGRPCError(err error) error {
	switch {
	default:
		return err
	}
}
