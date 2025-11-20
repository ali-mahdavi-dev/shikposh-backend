package command_handler

import (
	"context"
	"errors"
	"fmt"

	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/adapter/repository"
	"shikposh-backend/internal/account/domain/entity"

	"github.com/ali-mahdavi-dev/shikposh-framework/api/jwt"
	apperrors "github.com/ali-mahdavi-dev/shikposh-framework/errors"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

type RefreshTokenResult struct {
	AccessToken  string
	RefreshToken string
}

func (h *UserHandler) RefreshTokenHandler(ctx context.Context, refreshToken string) (*RefreshTokenResult, error) {
	var result *RefreshTokenResult

	err := h.uow.Do(ctx, func(ctx context.Context) error {
		// Parse and validate refresh token
		token, err := jwtlib.Parse(refreshToken, func(token *jwtlib.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
				return nil, jwtlib.ErrSignatureInvalid
			}
			return h.cfg.JWT.Secret, nil
		})

		if err != nil || !token.Valid {
			return apperrors.Unauthorized(accountphrases.UserNotFound)
		}

		// Extract user ID from token
		claims, ok := token.Claims.(jwtlib.MapClaims)
		if !ok {
			return apperrors.Unauthorized(accountphrases.UserNotFound)
		}

		userID := cast.ToUint64(claims["user_id"])
		if userID == 0 {
			return apperrors.Unauthorized(accountphrases.UserNotFound)
		}

		// Find token in database
		dbToken, err := h.uow.Token(ctx).FindByUserID(ctx, entity.UserID(userID))
		if err != nil {
			if errors.Is(err, repository.ErrTokenNotFound) {
				return apperrors.Unauthorized(accountphrases.UserNotFound)
			}
			return fmt.Errorf("RefreshTokenHandler failed to get token by userID: %w", err)
		}

		// Verify refresh token matches (handle null case)
		if dbToken == nil || dbToken.RefreshToken == "" || dbToken.RefreshToken != refreshToken {
			return apperrors.Unauthorized(accountphrases.UserNotFound)
		}

		// Generate new access token
		newAccessToken, err := jwt.GenerateToken(h.cfg.JWT.AccessTokenExpireDuration, h.cfg.JWT.Secret, userID)
		if err != nil {
			return fmt.Errorf("RefreshTokenHandler failed to generate access token: %w", err)
		}

		// Generate new refresh token
		newRefreshToken, err := jwt.GenerateToken(h.cfg.JWT.RefreshTokenExpireDuration, h.cfg.JWT.Secret, userID)
		if err != nil {
			return fmt.Errorf("RefreshTokenHandler failed to generate refresh token: %w", err)
		}

		// Update token in database
		dbToken.Token = newAccessToken
		dbToken.RefreshToken = newRefreshToken
		err = h.uow.Token(ctx).Save(ctx, dbToken)
		if err != nil {
			return fmt.Errorf("RefreshTokenHandler failed to save token: %w", err)
		}

		result = &RefreshTokenResult{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
