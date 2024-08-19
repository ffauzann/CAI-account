package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ffauzann/CAI-account/internal/constant"
	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/redis/go-redis/v9"
)

func (r *redisRepository) SetOTPRegister(ctx context.Context, data *model.RedisSetOTPRegisterData) (err error) {
	key := fmt.Sprintf(constant.RedisKeyOTPRegisterFormat, data.PhoneNumber)

	// Parse exp duration from config.
	d, err := time.ParseDuration(r.config.Dependency.Whatsapp.LoginOTP.Exp)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Convert struct into slice of byte.
	b, err := json.Marshal(data)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Set key-value.
	err = r.redis.Set(ctx, key, b, d).Err()
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *redisRepository) GetOTPRegister(ctx context.Context, phoneNumber string) (data *model.RedisSetOTPRegisterData, err error) {
	key := fmt.Sprintf(constant.RedisKeyOTPRegisterFormat, phoneNumber)

	// Cache not found.
	cmd := r.redis.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return
	}

	// Cache found. Continue to scan.
	b := []byte{}
	err = cmd.Scan(&b)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Unmarshal into struct.
	data = new(model.RedisSetOTPRegisterData)
	err = json.Unmarshal(b, &data)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *redisRepository) DeleteOTPRegister(ctx context.Context, phoneNumber string) (err error) {
	key := fmt.Sprintf(constant.RedisKeyOTPRegisterFormat, phoneNumber)

	return r.redis.Del(ctx, key).Err()
}

func (r *redisRepository) SetOTPLogin(ctx context.Context, data *model.RedisSetOTPLoginData) (err error) {
	key := fmt.Sprintf(constant.RedisKeyOTPLoginFormat, data.PhoneNumber)

	// Parse exp duration from config.
	d, err := time.ParseDuration(r.config.Dependency.Whatsapp.LoginOTP.Exp)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Convert struct into slice of byte.
	b, err := json.Marshal(data)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Set key-value.
	err = r.redis.Set(ctx, key, b, d).Err()
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *redisRepository) GetOTPLogin(ctx context.Context, phoneNumber string) (data *model.RedisSetOTPLoginData, err error) {
	key := fmt.Sprintf(constant.RedisKeyOTPLoginFormat, phoneNumber)

	// Cache not found.
	cmd := r.redis.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return
	}

	// Cache found. Continue to scan.
	b := []byte{}
	err = cmd.Scan(&b)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	// Unmarshal into struct.
	data = new(model.RedisSetOTPLoginData)
	err = json.Unmarshal(b, &data)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	return
}

func (r *redisRepository) DeleteOTPLogin(ctx context.Context, phoneNumber string) (err error) {
	key := fmt.Sprintf(constant.RedisKeyOTPLoginFormat, phoneNumber)

	return r.redis.Del(ctx, key).Err()
}
