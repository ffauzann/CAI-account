package client

import (
	"context"

	"github.com/ffauzann/CAI-account/proto/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *userClient) CloseUserAccount(ctx context.Context, req *emptypb.Empty) (*gen.CloseUserAccountResponse, error) {
	return c.userClient.CloseUserAccount(ctx, req)
}

func (c *userClient) IsUserExist(ctx context.Context, req *gen.IsUserExistRequest) (*gen.IsUserExistResponse, error) {
	return c.userClient.IsUserExist(ctx, req)
}
