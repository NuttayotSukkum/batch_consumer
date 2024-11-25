package dto

import (
	"github.com/NuttayotSukkum/batch_consumer/internals/pkg/constants"
)

type CommonErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Message     string `json:"message"`
}

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *BaseError) Error() string {
	return e.Message
}

func NewBaseError(code int, desc string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: desc,
	}
}
func ResponseGenericError() CommonErrorResponse {
	return CommonErrorResponse{
		Code:        constants.StatusCodeUnknowError,
		Description: "Internal Server Error",
		Title:       "Internal Server Error",
		Message:     "The unknown error occurred, please try again later.",
	}
}

func ResponseErrorBucketIsEmpty() CommonErrorResponse {
	return CommonErrorResponse{
		Code:        constants.StatusCodeBucketIsEmpty,
		Description: "Bucket is Empty",
		Title:       "Bucket is Empty",
		Message:     "Bucket is Empty, Please try again later.",
	}
}
