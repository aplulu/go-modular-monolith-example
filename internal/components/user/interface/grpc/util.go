package grpc

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aplulu/modular-monolith-example-go/internal/components/user/domain/model"
	pbUser "github.com/aplulu/modular-monolith-example-go/internal/grpc/example/user/v1"
)

// toPBUser ユーザーをgRPCのユーザーに変換
func toPBUser(user *model.User) *pbUser.User {
	return &pbUser.User{
		Id:   user.ID,
		Name: user.Name,
	}
}

// toGRPCError gRPCのステータスコードに変換
func toGRPCError(err error) error {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return status.Error(codes.NotFound, "not found")
	default:
		return err
	}
}
