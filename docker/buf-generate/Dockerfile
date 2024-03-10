FROM golang:1.22.1-alpine

WORKDIR /workspace

RUN go install github.com/bufbuild/buf/cmd/buf@v1.30.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 \
    && go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@v1.10.0

ENTRYPOINT ["/go/bin/buf"]
