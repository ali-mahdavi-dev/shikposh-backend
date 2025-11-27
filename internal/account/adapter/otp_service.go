package adapter

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"shikposh-backend/config"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/redisx"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	redis redisx.Connection
	cfg   *config.Config
}

func NewOtpService(redis redisx.Connection, cfg *config.Config) *OtpService {
	return &OtpService{
		redis: redis,
		cfg:   cfg,
	}
}

// GenerateOTP generates a random OTP code
func (o *OtpService) GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	min := int64(100000) // 6 digits minimum
	max := int64(999999) // 6 digits maximum
	otp := rand.Int63n(max-min+1) + min
	return fmt.Sprintf("%06d", otp)
}

// SendOTP stores OTP in Redis with expiration
func (o *OtpService) SendOTP(ctx context.Context, phone string) (string, error) {
	// Generate OTP
	otp := o.GenerateOTP()

	// Create Redis key: otp:{phone}
	key := fmt.Sprintf("otp:%s", phone)

	// Store OTP in Redis with expiration
	expiration := o.cfg.Otp.ExpireTime * time.Second
	err := o.redis.SetValue(ctx, key, otp, expiration)
	if err != nil {
		return "", fmt.Errorf("OtpService.SendOTP failed to store OTP: %w", err)
	}

	// Log OTP for development/testing purposes
	// TODO: In production, replace this with actual SMS service integration
	logging.Info("OTP generated and stored").
		WithString("phone", phone).
		WithString("otp", otp).
		WithString("expires_in", expiration.String()).
		Log()

	return otp, nil
}

// VerifyOTP verifies the OTP code
func (o *OtpService) VerifyOTP(ctx context.Context, phone string, otp string) (bool, error) {
	// Create Redis key
	key := fmt.Sprintf("otp:%s", phone)

	// Get OTP from Redis
	storedOTP, err := o.redis.GetValue(ctx, key)
	if err != nil {
		// Check if error is redis.Nil (key doesn't exist) - this is expected when OTP is expired or already used
		if errors.Is(err, redis.Nil) {
			logging.Warn("OTP verification failed: OTP not found or expired").
				WithString("phone", phone).
				Log()
			return false, nil // Return false, nil (not an error, just invalid OTP)
		}
		// For other Redis errors, return error
		logging.Warn("OTP verification failed: Redis error").
			WithString("phone", phone).
			WithError(err).
			Log()
		return false, fmt.Errorf("OtpService.VerifyOTP failed to get OTP: %w", err)
	}

	// Compare OTPs
	if storedOTP != otp {
		logging.Warn("OTP verification failed: invalid OTP").
			WithString("phone", phone).
			Log()
		return false, nil
	}

	// Delete OTP after successful verification
	_ = o.redis.DeleteKey(ctx, key)

	logging.Info("OTP verified successfully").
		WithString("phone", phone).
		Log()

	return true, nil
}

// CheckRateLimit checks if phone number has exceeded rate limit
func (o *OtpService) CheckRateLimit(ctx context.Context, phone string) (bool, error) {
	key := fmt.Sprintf("otp:rate:%s", phone)
	exists, err := o.redis.ExistsKey(ctx, key)
	if err != nil {
		return false, err
	}
	return !exists, nil // Return true if not rate limited
}

// SetRateLimit sets rate limit for phone number
func (o *OtpService) SetRateLimit(ctx context.Context, phone string) error {
	key := fmt.Sprintf("otp:rate:%s", phone)
	expiration := o.cfg.Otp.Limiter * time.Second
	return o.redis.SetValue(ctx, key, "1", expiration)
}
