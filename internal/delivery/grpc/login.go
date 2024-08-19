package grpc

import (
	"context"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/CAI-account/proto/gen"
	"github.com/ffauzann/common/util/str"
)

func (s *srv) Login(ctx context.Context, req *gen.LoginRequest) (res *gen.LoginResponse, err error) {
	param := util.CastStruct[model.LoginRequest](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	result, err := s.service.Login(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = &gen.LoginResponse{
		AccessToken:  result.Token.AccessToken,
		RefreshToken: result.Token.RefreshToken,
	}

	return
}

func (s *srv) LoginV2(ctx context.Context, req *gen.LoginV2Request) (res *gen.LoginV2Response, err error) {
	param := util.CastStruct[model.LoginV2Request](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	param.PhoneNumber = str.PhoneWithCountryCode(param.PhoneNumber, constant.DefaultCountryCode, true)

	result, err := s.service.LoginV2(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.LoginV2Response](result)

	return
}
