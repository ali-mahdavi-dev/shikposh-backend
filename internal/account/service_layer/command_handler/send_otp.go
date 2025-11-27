package command_handler

import (
	"context"
	"fmt"

	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/domain/commands"

	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"
)

// SendOtpHandler handles sending OTP
func (h *OtpHandler) SendOtpHandler(ctx context.Context, cmd *commands.SendOtp) error {
	// Check rate limit
	allowed, err := h.otpService.CheckRateLimit(ctx, cmd.Phone)
	if err != nil {
		return fmt.Errorf("OtpHandler.SendOtpHandler failed to check rate limit: %w", err)
	}
	if !allowed {
		return apperrors.RateLimit(accountphrases.OtpRateLimited)
	}

	// Send OTP
	_, err = h.otpService.SendOTP(ctx, cmd.Phone)
	if err != nil {
		return fmt.Errorf("OtpHandler.SendOtpHandler failed to send OTP: %w", err)
	}

	// Set rate limit
	err = h.otpService.SetRateLimit(ctx, cmd.Phone)
	if err != nil {
		// Log error but don't fail the request
		fmt.Printf("Failed to set rate limit: %v\n", err)
	}

	return nil
}
