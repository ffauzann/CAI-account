package grpc

import (
	"context"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/CAI-account/proto/gen"
)

func (s *srv) GetListAccount(ctx context.Context, req *gen.GetListAccountRequest) (res *gen.GetListAccountResponse, err error) {
	param := util.CastStruct[model.GetListAccountRequest](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	claims, ok := util.ClaimsFromContext(ctx)
	if !ok {
		err = constant.ErrUnauthenticated
		util.LogContext(ctx).Warn(err.Error())
		return
	}
	param.UserId = claims.UserId

	result, err := s.service.GetListAccount(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.GetListAccountResponse](result)
	for i := range result.Accounts {
		res.Accounts[i].Category = constant.CategoryInternalToGenMap[result.Accounts[i].Category]
	}

	return
}

func (s *srv) UpdateBalance(ctx context.Context, req *gen.UpdateBalanceRequest) (res *gen.UpdateBalanceResponse, err error) {
	param := util.CastStruct[model.UpdateBalanceRequest](req)
	if err = util.ValidateStruct(param); err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	if param.DestinationAccountId != 0 && param.Amount > 0 {
		err = constant.ErrInvalidTransferAmount
		util.LogContext(ctx).Error(err.Error())
		return
	}

	claims, ok := util.ClaimsFromContext(ctx)
	if !ok {
		err = constant.ErrUnauthenticated
		util.LogContext(ctx).Warn(err.Error())
		return
	}
	param.UserId = claims.UserId

	result, err := s.service.UpdateBalance(ctx, param)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	res = util.CastStruct[gen.UpdateBalanceResponse](result)

	return
}
