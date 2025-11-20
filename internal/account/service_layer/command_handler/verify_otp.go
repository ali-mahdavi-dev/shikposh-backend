package command_handler

import (
	"context"
	"fmt"

	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/api/jwt"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"

	"github.com/pkg/errors"
)

// VerifyOtpResult contains the result of OTP verification
type VerifyOtpResult struct {
	Token        string
	RefreshToken string
	User         *entity.User
	UserExists   bool
}

// VerifyOtpHandler handles OTP verification
func (h *OtpHandler) VerifyOtpHandler(ctx context.Context, cmd *commands.VerifyOtp) (*VerifyOtpResult, error) {
	// Verify OTP
	valid, err := h.otpService.VerifyOTP(ctx, cmd.Phone, cmd.OTP, cmd.Type)
	if err != nil {
		return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to verify OTP: %w", err)
	}
	if !valid {
		return nil, apperrors.Unauthorized(accountphrases.OtpInvalid)
	}

	// Check if user exists
	user, err := h.uow.User(ctx).FindByPhone(ctx, cmd.Phone)
	userExists := err == nil && user != nil

	var accessToken string
	var refreshToken string
	var finalUser *entity.User

	if userExists {
		// User exists - generate token and login
		// Check if user has existing token and remove it
		token, err := h.uow.Token(ctx).FindByUserID(ctx, user.ID)
		if err != nil && !errors.Is(err, repository.ErrTokenNotFound) {
			return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to get token by userID: %w", err)
		}

		if token != nil {
			if err := h.uow.Token(ctx).Remove(ctx, token, false); err != nil {
				return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to remove existing token: %w", err)
			}
		}

		// Generate new access token
		accessToken, err = jwt.GenerateToken(h.cfg.JWT.AccessTokenExpireDuration, h.cfg.JWT.Secret, uint64(user.ID))
		if err != nil {
			return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to generate access token: %w", err)
		}

		// Generate new refresh token
		refreshToken, err = jwt.GenerateToken(h.cfg.JWT.RefreshTokenExpireDuration, h.cfg.JWT.Secret, uint64(user.ID))
		if err != nil {
			return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to generate refresh token: %w", err)
		}

		// Save new token
		err = h.uow.Token(ctx).Save(ctx, entity.NewToken(accessToken, refreshToken, user.ID))
		if err != nil {
			return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to save token: %w", err)
		}

		finalUser = user
	} else if cmd.Type == "register" {
		// For register flow after user has registered, generate token
		// Check if user was just created (this happens when verify is called after register)
		user, err = h.uow.User(ctx).FindByPhone(ctx, cmd.Phone)
		if err == nil && user != nil {
			// User now exists, generate token
			// Check if user has existing token and remove it
			token, err := h.uow.Token(ctx).FindByUserID(ctx, user.ID)
			if err != nil && !errors.Is(err, repository.ErrTokenNotFound) {
				return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to get token by userID: %w", err)
			}

			if token != nil {
				if err := h.uow.Token(ctx).Remove(ctx, token, false); err != nil {
					return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to remove existing token: %w", err)
				}
			}

			// Generate new access token
			accessToken, err = jwt.GenerateToken(h.cfg.JWT.AccessTokenExpireDuration, h.cfg.JWT.Secret, uint64(user.ID))
			if err != nil {
				return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to generate access token: %w", err)
			}

			// Generate new refresh token
			refreshToken, err = jwt.GenerateToken(h.cfg.JWT.RefreshTokenExpireDuration, h.cfg.JWT.Secret, uint64(user.ID))
			if err != nil {
				return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to generate refresh token: %w", err)
			}

			// Save new token
			err = h.uow.Token(ctx).Save(ctx, entity.NewToken(accessToken, refreshToken, user.ID))
			if err != nil {
				return nil, fmt.Errorf("OtpHandler.VerifyOtpHandler failed to save token: %w", err)
			}

			finalUser = user
			userExists = true
		}
	}

	return &VerifyOtpResult{
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         finalUser,
		UserExists:   userExists,
	}, nil
}
