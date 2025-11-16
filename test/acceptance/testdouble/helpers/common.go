package helpers

import (
	apperrors "github.com/ali-mahdavi-dev/framework/errors"

	. "github.com/onsi/gomega"
)

// GetErrorType extracts error type from error
func GetErrorType(err error) apperrors.ErrorType {
	appErr, ok := err.(apperrors.Error)
	Expect(ok).To(BeTrue())
	return appErr.Type()
}

