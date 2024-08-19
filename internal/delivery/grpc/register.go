package grpc

import (
	"context"
	"slices"
	"strings"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/CAI-account/proto/gen"
	"github.com/ffauzann/common/util/str"
)

func (s *srv) Register(ctx context.Context, req *gen.RegisterRequest) (res *gen.RegisterResponse, err error) {
	param := util.CastStruct[model.RegisterRequest](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Set to default Role ID if claims is empty. Otherwise, get from payload.
	claims, ok := util.ClaimsFromContext(ctx)
	if !ok {
		param.RoleId = constant.DefaultRoleId
	} else {
		userRoleId := claims.RoleId
		if !slices.Contains(constant.AllowedRolesRegisterMap[userRoleId], param.RoleId) {
			err = constant.ErrPermissionDenied
			util.LogContext(ctx).Warn(err.Error())
			return
		}
	}

	// if !slices.Contains(constant.AllowedRolesRegister, userRoleId) {
	// 	err = constant.ErrPermissionDenied
	// 	util.LogContext(ctx).Warn(err.Error())
	// 	return
	// }

	// if !slices.Contains(constant.AllowedRolesRegisterMap[userRoleId], param.RoleId) {
	// 	err = constant.ErrPermissionDenied
	// 	util.LogContext(ctx).Warn(err.Error())
	// 	return
	// }

	result, err := s.service.Register(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = &gen.RegisterResponse{
		Code:    gen.RegisterStatusCode(result.StatusCode),
		Reasons: result.Reasons,
	}

	return
}

func (s *srv) RegisterV2(ctx context.Context, req *gen.RegisterV2Request) (res *gen.RegisterV2Response, err error) {
	param := util.CastStruct[model.RegisterV2Request](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	param.PhoneNumber = strings.ReplaceAll(str.PhoneWithCountryCode(param.PhoneNumber, constant.DefaultCountryCode, true), "+", "")

	result, err := s.service.RegisterV2(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.RegisterV2Response](result)

	return
}
