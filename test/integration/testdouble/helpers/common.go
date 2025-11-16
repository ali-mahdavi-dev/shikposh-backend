package helpers

import (
	apperrors "github.com/shikposh/framework/errors"

	. "github.com/onsi/gomega"
)

func GetErrorType(err error) apperrors.ErrorType {
	appErr, ok := err.(apperrors.Error)
	Expect(ok).To(BeTrue())
	return appErr.Type()
}

