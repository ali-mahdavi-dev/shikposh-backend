package payment

import (
	"context"
	"fmt"

	"shikposh-backend/config"

	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	"github.com/go-resty/resty/v2"
)

type ZarinPalService struct {
	cfg    *config.ZarinPalConfig
	client *resty.Client
}

func NewZarinPalService(cfg *config.ZarinPalConfig) *ZarinPalService {
	baseURL := "https://sandbox.zarinpal.com"
	if !cfg.Sandbox {
		baseURL = "https://www.zarinpal.com"
	}

	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Content-Type", "application/json")

	return &ZarinPalService{
		cfg:    cfg,
		client: client,
	}
}

type PaymentRequest struct {
	MerchantID  string `json:"merchant_id"`
	Amount      int64  `json:"amount"`
	CallbackURL string `json:"callback_url"`
	Description string `json:"description"`
}

type PaymentResponse struct {
	Status    int    `json:"Status"`
	Authority string `json:"Authority"`
}

type VerifyRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"Authority"`
	Amount     int64  `json:"Amount"`
}

type VerifyResponse struct {
	Status int   `json:"Status"`
	RefID  int64 `json:"RefID"`
}

// RequestPayment creates a payment request and returns the payment authority
// Amount should be in Tomans (not Rials). If your system stores prices in Rials, convert by dividing by 10.
func (s *ZarinPalService) RequestPayment(ctx context.Context, amount int64, description string) (string, error) {
	// Validate merchant ID
	if s.cfg.MerchantID == "" || s.cfg.MerchantID == "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" {
		return "", fmt.Errorf("zarinpal merchant ID is not configured. Please set a valid merchant ID in config")
	}

	// Validate amount (minimum 1000 Toman)
	// If amount seems to be in Rials (very large number), convert to Tomans
	amountInTomans := amount
	if amount > 100000 {
		// Likely in Rials, convert to Tomans
		amountInTomans = amount / 10
		logging.Info("ZarinPal RequestPayment: converting amount from Rials to Tomans").
			WithInt64("original_amount", amount).
			WithInt64("converted_amount", amountInTomans).
			Log()
	}

	if amountInTomans < 1000 {
		return "", fmt.Errorf("zarinpal payment amount must be at least 1000 Toman. Current amount: %d", amountInTomans)
	}

	req := PaymentRequest{
		MerchantID:  s.cfg.MerchantID,
		Amount:      amountInTomans,
		CallbackURL: s.cfg.CallbackURL,
		Description: description,
	}

	var resp PaymentResponse
	httpResp, err := s.client.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/pg/v4/payment/request.json")

	if err != nil {
		logging.Error("ZarinPal RequestPayment: failed to send request").
			WithError(err).
			WithInt64("amount", amount).
			WithInt64("amount_in_tomans", amountInTomans).
			Log()
		return "", fmt.Errorf("failed to request payment: %w", err)
	}

	if httpResp.IsError() {
		logging.Error("ZarinPal RequestPayment: API error").
			WithString("status", httpResp.Status()).
			WithString("body", string(httpResp.Body())).
			WithInt64("amount", amount).
			WithInt64("amount_in_tomans", amountInTomans).
			Log()
		return "", fmt.Errorf("zarinpal API error: status %s, body: %s", httpResp.Status(), string(httpResp.Body()))
	}

	if resp.Status != 100 {
		// Log full response for debugging
		logging.Warn("ZarinPal RequestPayment: payment request failed").
			WithInt("status", resp.Status).
			WithInt64("original_amount", amount).
			WithInt64("amount_in_tomans", amountInTomans).
			WithString("merchant_id", s.cfg.MerchantID).
			WithString("callback_url", s.cfg.CallbackURL).
			Log()

		// Provide more helpful error messages based on status code
		switch resp.Status {
		case 0:
			return "", fmt.Errorf("zarinpal payment request failed: invalid merchant ID or amount. Please check your ZarinPal merchant ID in config file. For sandbox testing, use a valid test merchant ID from ZarinPal. Status: 0")
		case -1:
			return "", fmt.Errorf("zarinpal payment request failed: information submitted is incomplete. Status: -1")
		case -2:
			return "", fmt.Errorf("zarinpal payment request failed: merchant ID or IP is not correct. Status: -2")
		case -3:
			return "", fmt.Errorf("zarinpal payment request failed: amount should be at least 1000 Toman. Current amount: %d Toman. Status: -3", amountInTomans)
		case -4:
			return "", fmt.Errorf("zarinpal payment request failed: merchant level is not sufficient. Status: -4")
		default:
			return "", fmt.Errorf("zarinpal payment request failed with status: %d", resp.Status)
		}
	}

	logging.Info("ZarinPal RequestPayment: payment request successful").
		WithString("authority", resp.Authority).
		WithInt64("original_amount", amount).
		WithInt64("amount_in_tomans", amountInTomans).
		Log()

	return resp.Authority, nil
}

// VerifyPayment verifies a payment using authority and amount
// Amount should be in Tomans (same as RequestPayment)
func (s *ZarinPalService) VerifyPayment(ctx context.Context, authority string, amount int64) (int64, error) {
	// Convert amount from Rials to Tomans if needed (same logic as RequestPayment)
	amountInTomans := amount
	if amount > 100000 {
		amountInTomans = amount / 10
	}

	req := VerifyRequest{
		MerchantID: s.cfg.MerchantID,
		Authority:  authority,
		Amount:     amountInTomans,
	}

	var resp VerifyResponse
	httpResp, err := s.client.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/pg/v4/payment/verify.json")

	if err != nil {
		logging.Error("ZarinPal VerifyPayment: failed to send request").
			WithError(err).
			WithString("authority", authority).
			WithInt64("amount", amount).
			Log()
		return 0, fmt.Errorf("failed to verify payment: %w", err)
	}

	if httpResp.IsError() {
		logging.Error("ZarinPal VerifyPayment: API error").
			WithString("status", httpResp.Status()).
			WithString("body", string(httpResp.Body())).
			WithString("authority", authority).
			WithInt64("amount", amount).
			Log()
		return 0, fmt.Errorf("zarinpal API error: status %s, body: %s", httpResp.Status(), string(httpResp.Body()))
	}

	if resp.Status != 100 && resp.Status != 101 {
		logging.Warn("ZarinPal VerifyPayment: payment verification failed").
			WithInt("status", resp.Status).
			WithString("authority", authority).
			WithInt64("amount", amount).
			Log()
		return 0, fmt.Errorf("zarinpal payment verification failed with status: %d", resp.Status)
	}

	logging.Info("ZarinPal VerifyPayment: payment verified successfully").
		WithInt64("ref_id", resp.RefID).
		WithString("authority", authority).
		WithInt64("amount", amount).
		Log()

	return resp.RefID, nil
}

// GetPaymentURL returns the payment gateway URL for redirecting the user
func (s *ZarinPalService) GetPaymentURL(authority string) string {
	baseURL := "https://sandbox.zarinpal.com"
	if !s.cfg.Sandbox {
		baseURL = "https://www.zarinpal.com"
	}
	return fmt.Sprintf("%s/pg/StartPay/%s", baseURL, authority)
}
