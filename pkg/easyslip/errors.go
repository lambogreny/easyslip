package easyslip

import (
	"fmt"
)

// Error represents an API error
type Error struct {
	Status  int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("easyslip: %s (status: %d)", e.Message, e.Status)
}

// NewError creates a new Error
func NewError(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}

// Predefined error messages
var (
	ErrInvalidImage           = NewError(400, "invalid_image")
	ErrImageSizeTooLarge      = NewError(400, "image_size_too_large")
	ErrInvalidTransactionID   = NewError(400, "invalid_transaction_id")
	ErrInvalidAmount          = NewError(400, "invalid_amount")
	ErrInvalidSenderName      = NewError(400, "invalid_sender_name")
	ErrInvalidReceiverName    = NewError(400, "invalid_receiver_name")
	ErrInvalidReceiverPhone   = NewError(400, "invalid_receiver_phone")
	ErrUnauthorized           = NewError(401, "unauthorized")
	ErrAccessDenied           = NewError(403, "access_denied")
	ErrAccountNotVerified     = NewError(403, "account_not_verified")
	ErrApplicationExpired     = NewError(403, "application_expired")
	ErrApplicationDeactivated = NewError(403, "application_deactivated")
	ErrQuotaExceeded          = NewError(403, "quota_exceeded")
	ErrSlipNotFound           = NewError(404, "slip_not_found")
	ErrQRCodeNotFound         = NewError(404, "qrcode_not_found")
	ErrServerError            = NewError(500, "server_error")
	ErrAPIServerError         = NewError(500, "api_server_error")
	ErrInvalidPayload         = NewError(400, "invalid_payload")
)

// Map of error messages to predefined errors
var errorMap = map[string]*Error{
	"invalid_image":           ErrInvalidImage,
	"image_size_too_large":    ErrImageSizeTooLarge,
	"invalid_transaction_id":  ErrInvalidTransactionID,
	"invalid_amount":          ErrInvalidAmount,
	"invalid_sender_name":     ErrInvalidSenderName,
	"invalid_receiver_name":   ErrInvalidReceiverName,
	"invalid_receiver_phone":  ErrInvalidReceiverPhone,
	"unauthorized":            ErrUnauthorized,
	"access_denied":           ErrAccessDenied,
	"account_not_verified":    ErrAccountNotVerified,
	"application_expired":     ErrApplicationExpired,
	"application_deactivated": ErrApplicationDeactivated,
	"quota_exceeded":          ErrQuotaExceeded,
	"slip_not_found":          ErrSlipNotFound,
	"qrcode_not_found":        ErrQRCodeNotFound,
	"server_error":            ErrServerError,
	"api_server_error":        ErrAPIServerError,
	"invalid_payload":         ErrInvalidPayload,
}

// NewErrorFromMessage creates a new Error from an error message
func NewErrorFromMessage(status int, message string) *Error {
	if err, exists := errorMap[message]; exists {
		return err
	}
	return NewError(status, message)
}
