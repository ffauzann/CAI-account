package grpc

import (
	"context"

	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/CAI-account/proto/gen"
)

func (s *srv) RefreshToken(ctx context.Context, req *gen.RefreshTokenRequest) (res *gen.RefreshTokenResponse, err error) {
	param := util.CastStruct[model.RefreshTokenRequest](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	result, err := s.service.RefreshToken(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.RefreshTokenResponse](result)
	return
}
