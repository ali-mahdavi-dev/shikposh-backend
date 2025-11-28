package handler

import (
	"image/png"

	"shikposh-backend/internal/account/adapter"
	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/domain/commands"
	"shikposh-backend/internal/account/service_layer/command_handler"

	httpapi "github.com/ali-mahdavi-dev/shikposh-framework/api/http"
	"github.com/ali-mahdavi-dev/shikposh-framework/errors"
	"github.com/ali-mahdavi-dev/shikposh-framework/service_layer/messagebus"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

type UserController struct {
	bus         messagebus.MessageBus
	ag          *adapter.AvatarGenerator
	userHandler *command_handler.UserHandler
	otpHandler  *command_handler.OtpHandler
}

func NewUserController(bus messagebus.MessageBus, ag *adapter.AvatarGenerator, userHandler *command_handler.UserHandler, otpHandler *command_handler.OtpHandler) *UserController {
	return &UserController{
		bus:         bus,
		ag:          ag,
		userHandler: userHandler,
		otpHandler:  otpHandler,
	}
}

func (u *UserController) RegisterRoutes(r fiber.Router) {
	publicRoute := r.Group("/api/v1/public")
	{
		publicRoute.Post("/avatar/:id", u.GenerateAvatarHandler)
		publicRoute.Post("/register", u.Register)
		publicRoute.Post("/logout", u.Logout)

		// OTP endpoints
		authRoute := publicRoute.Group("/auth")
		{
			authRoute.Post("/send-otp", u.SendOtp)
			authRoute.Post("/verify-otp", u.VerifyOtp)
			authRoute.Post("/refresh", u.RefreshToken)
		}
	}
}

func (u *UserController) GenerateAvatarHandler(c fiber.Ctx) error {
	identifier := c.Params("id")

	img, err := u.ag.Generate(identifier)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	// Set headers for PNG response
	c.Set("Content-Type", "image/png")

	// Encode image directly to response
	if err := png.Encode(c.Response().BodyWriter(), img); err != nil {
		return httpapi.ResError(c, err)
	}

	return nil
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Handles user registration by parsing the request body and invoking the registration command.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		commands.RegisterUser	true	"RegisterUser request"
//	@Success		201		{object}	httpapi.ResponseResult	"Registration successful"
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request body or unknown provider"
//	@Failure		409		{object}	httpapi.ResponseResult	"User already exists"
//	@Failure		422		{object}	httpapi.ResponseResult	"Unprocessable input (validation failed)"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/public/register [post]
func (u *UserController) Register(c fiber.Ctx) error {
	ctx := c.Context()
	cmd := new(commands.RegisterUser)

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	result, err := u.userHandler.RegisterHandler(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	response := map[string]interface{}{
		"success": true,
		"message": "کاربر با موفقیت ثبت نام شد",
		"token":   result.Token,
	}

	if result.RefreshToken != "" {
		response["refresh_token"] = result.RefreshToken
	}

	if result.User != nil {
		response["user"] = map[string]interface{}{
			"id":         result.User.ID,
			"first_name": result.User.FirstName,
			"last_name":  result.User.LastName,
			"email":      result.User.Email,
			"phone":      result.User.Phone,
		}
	}

	return httpapi.ResSuccess(c, response)
}

// Logout godoc
//
//	@Summary		Logout user
//	@Description	Logs out the authenticated user.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		204	"Logout completed successfully"
//	@Failure		404	{object}	httpapi.ResponseResult	"User not found"
//	@Failure		500	{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/public/logout [post]
func (u *UserController) Logout(c fiber.Ctx) error {
	ctx := c.Context()

	userID := c.Get("user_id")
	if userID == "" {
		return httpapi.ResError(c, errors.NotFound(accountphrases.UserNotFound))
	}

	cmd := new(commands.Logout)
	cmd.UserID = cast.ToUint64(userID)

	err := u.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// SendOtp godoc
//
//	@Summary		Send OTP
//	@Description	Sends OTP code to the provided phone number
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		commands.SendOtp		true	"SendOtp request"
//	@Success		200		{object}	httpapi.ResponseResult	"OTP sent successfully"
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request body"
//	@Failure		422		{object}	httpapi.ResponseResult	"Validation failed"
//	@Failure		429		{object}	httpapi.ResponseResult	"Rate limited - too many requests"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/public/auth/send-otp [post]
func (u *UserController) SendOtp(c fiber.Ctx) error {
	ctx := c.Context()
	cmd := new(commands.SendOtp)

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	err := u.bus.Handle(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	return httpapi.ResSuccess(c, map[string]interface{}{
		"success":    true,
		"message":    "کد OTP با موفقیت ارسال شد",
		"expires_in": 120, // seconds
	})
}

// VerifyOtp godoc
//
//	@Summary		Verify OTP
//	@Description	Verifies OTP code and returns token if user exists, or indicates if user needs to register
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		commands.VerifyOtp		true	"VerifyOtp request"
//	@Success		200		{object}	httpapi.ResponseResult	"OTP verified successfully"
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request body"
//	@Failure		401		{object}	httpapi.ResponseResult	"Invalid or expired OTP"
//	@Failure		422		{object}	httpapi.ResponseResult	"Validation failed"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/public/auth/verify-otp [post]
func (u *UserController) VerifyOtp(c fiber.Ctx) error {
	ctx := c.Context()
	cmd := new(commands.VerifyOtp)

	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	result, err := u.otpHandler.VerifyOtpHandler(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	if result == nil {
		return httpapi.ResError(c, errors.Internal("Unexpected error: result is nil"))
	}

	response := map[string]interface{}{
		"success":     true,
		"user_exists": result.UserExists,
	}

	// Return tokens in response body (frontend will handle storage)
	if result.Token != "" {
		response["token"] = result.Token
		if result.RefreshToken != "" {
			response["refresh_token"] = result.RefreshToken
		}
	}

	if result.User != nil {
		response["user"] = map[string]interface{}{
			"id":         result.User.ID,
			"first_name": result.User.FirstName,
			"last_name":  result.User.LastName,
			"email":      result.User.Email,
			"phone":      result.User.Phone,
		}
	}

	return httpapi.ResSuccess(c, response)
}

// RefreshToken godoc
//
//	@Summary		Refresh access token
//	@Description	Refreshes the access token using a valid refresh token from request body
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		map[string]string		true	"Refresh token request"	example({"refresh_token": "string"})
//	@Success		200		{object}	httpapi.ResponseResult	"Token refreshed successfully"
//	@Failure		400		{object}	httpapi.ResponseResult	"Invalid request body"
//	@Failure		401		{object}	httpapi.ResponseResult	"Invalid or expired refresh token"
//	@Failure		404		{object}	httpapi.ResponseResult	"User not found"
//	@Failure		500		{object}	httpapi.ResponseResult	"Internal server error"
//	@Router			/api/v1/public/auth/refresh [post]
func (u *UserController) RefreshToken(c fiber.Ctx) error {
	ctx := c.Context()

	cmd := new(commands.RefreshToken)
	if err := httpapi.ParseJSON(c, cmd); err != nil {
		return httpapi.ResError(c, err)
	}

	if cmd.RefreshToken == "" || cmd.RefreshToken == "null" {
		return httpapi.ResError(c, errors.Unauthorized(accountphrases.UserNotFound))
	}

	result, err := u.userHandler.RefreshTokenHandler(ctx, cmd)
	if err != nil {
		return httpapi.ResError(c, err)
	}

	// Return new tokens in response body (frontend will handle storage)
	if result == nil {
		return httpapi.ResError(c, errors.Unauthorized(accountphrases.UserNotFound))
	}

	if result.AccessToken == "" {
		return httpapi.ResError(c, errors.Unauthorized(accountphrases.UserNotFound))
	}

	response := map[string]interface{}{
		"success": true,
		"message": "توکن با موفقیت بروزرسانی شد",
		"token":   result.AccessToken,
	}

	if result.RefreshToken != "" {
		response["refresh_token"] = result.RefreshToken
	}

	return httpapi.ResSuccess(c, response)
}
