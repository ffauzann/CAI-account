package grpc

import (
	"context"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/CAI-account/proto/gen"
	"github.com/ffauzann/common/util/str"
)

func (s *srv) VerifyOTP(ctx context.Context, req *gen.VerifyOTPRequest) (res *gen.VerifyOTPResponse, err error) {
	param := util.CastStruct[model.VerifyOTPRequest](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	param.PhoneNumber = str.PhoneWithCountryCode(param.PhoneNumber, constant.DefaultCountryCode, true)
	result, err := s.service.VerifyOTP(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.VerifyOTPResponse](result)

	res.Data, _ = util.CastToAnyMap(result.Data)

	return
}
