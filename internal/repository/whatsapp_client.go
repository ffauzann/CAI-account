package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"go.uber.org/zap"
)

func (r *whatsappClientRepository) Send(ctx context.Context, req *model.WhatsappClientSendTextRequest) (err error) {
	reqPayload := model.WhatsappClientSendTextRequestBody{
		InstanceID: r.config.Dependency.Whatsapp.InstanceID,
		Content:    req.Content,
		To:         req.PhoneNumber,
	}

	bytePayload, err := json.Marshal(&reqPayload)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.config.Dependency.Whatsapp.SenderURL, bytes.NewBuffer(bytePayload))
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-access-key", r.config.Dependency.Whatsapp.XAccessKey)

	httpRes, err := r.client.Do(httpReq)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}
	defer httpRes.Body.Close()

	res := new(model.WhatsappSendTextResponse)
	json.NewDecoder(httpRes.Body).Decode(res)

	util.LogContext(ctx).Info("", zap.Any("res", res))
	return
}
