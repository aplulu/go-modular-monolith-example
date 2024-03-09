package grpc

import (
	"context"
	"reflect"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceAdapter struct {
	service interface{}
}

func (s *ServiceAdapter) RegisterService(_ *grpc.ServiceDesc, impl interface{}) {
	s.service = impl
}

func (s *ServiceAdapter) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error {
	parts := strings.Split(method, "/")

	req := []reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(args),
	}

	result := reflect.ValueOf(s.service).MethodByName(parts[2]).Call(req)
	if !result[0].IsZero() {
		reflect.ValueOf(reply).Elem().Set(result[0].Elem())
	}

	err, _ := result[1].Interface().(error)
	return err
}

func (s *ServiceAdapter) NewStream(_ context.Context, _ *grpc.StreamDesc, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, status.Error(codes.Unimplemented, "streaming is not supported")
}

func NewServiceAdapter() *ServiceAdapter {
	return &ServiceAdapter{}
}
