// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package article

import (
	"context"

	"gozero/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PublishRequest  = pb.PublishRequest
	PublishResponse = pb.PublishResponse

	Article interface {
		Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	}

	defaultArticle struct {
		cli zrpc.Client
	}
)

func NewArticle(cli zrpc.Client) Article {
	return &defaultArticle{
		cli: cli,
	}
}

func (m *defaultArticle) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	client := pb.NewArticleClient(m.cli.Conn())
	return client.Publish(ctx, in, opts...)
}
