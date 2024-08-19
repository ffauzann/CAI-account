package grpc

import (
	"context"
	"regexp"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/CAI-account/proto/gen"
	"github.com/ffauzann/common/util/str"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *srv) IsUserExist(ctx context.Context, req *gen.IsUserExistRequest) (res *gen.IsUserExistResponse, err error) {
	if err = validateIsUserExist(req); err != nil {
		return
	}

	param := util.CastStruct[model.IsUserExistRequest](req)
	param.PhoneNumber = str.PhoneWithCountryCode(param.PhoneNumber, constant.DefaultCountryCode, true)
	result, err := s.service.IsUserExist(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.IsUserExistResponse](result)
	return
}

func validateIsUserExist(req *gen.IsUserExistRequest) error {
	if req.Email != "" {
		regexEmail := regexp.MustCompile(constant.RegexEmail)
		if !regexEmail.MatchString(req.GetEmail()) {
			return constant.ErrMalformedEmail
		}
	}

	if len(req.GetPhoneNumber()) < 4 { //nolint
		req.PhoneNumber = ""
	}

	if req.Email == "" && req.PhoneNumber == "" {
		return constant.ErrNoArg
	}

	return nil
}

func (s *srv) CloseUserAccount(ctx context.Context, req *emptypb.Empty) (res *gen.CloseUserAccountResponse, err error) {
	claims, ok := util.ClaimsFromContext(ctx)
	if !ok {
		err = constant.ErrUnauthenticated
		util.LogContext(ctx).Warn(err.Error())
		return
	}

	result, err := s.service.CloseUserAccount(ctx, &model.CloseUserAccountRequest{
		UserId: claims.UserId,
	})
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.CloseUserAccountResponse](result)

	return
}
