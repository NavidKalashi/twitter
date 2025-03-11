package repository

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/NavidKalashi/twitter/internal/core/ports"
	"github.com/redis/go-redis/v9"
)

type OTPRepository struct {
	redis *redis.Client
}

func NewOTPRepository(redis *redis.Client) ports.OTP {
	return &OTPRepository{redis: redis}
}

func (or *OTPRepository) Set(email string, code uint) error {
	ctx := context.Background()
	key := fmt.Sprintf("otp_%s", email)
	return or.redis.Set(ctx, key, code, 2 * time.Minute).Err()
}

func (or *OTPRepository) Get(email string) (uint, error) {
	ctx := context.Background()
	key := fmt.Sprintf("otp_%s", email)

	otpStr, err := or.redis.Get(ctx, key).Result()
	if err == redis.Nil {
			return 0, fmt.Errorf("OTP not found")
	} else if err != nil {
			return 0, err
	}

	// Validate OTP format
	if !regexp.MustCompile(`^\d+$`).MatchString(otpStr) {
			return 0, fmt.Errorf("invalid OTP format: %s", otpStr)
	}

	otpUint64, err := strconv.ParseUint(otpStr, 10, 32)
	if err != nil {
			return 0, fmt.Errorf("failed to parse OTP: %w", err)
	}

	return uint(otpUint64), nil
}