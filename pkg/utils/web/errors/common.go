package errors

import (
	"fmt"
	"net/http"
)

func InternalError(err error) *CodedError {
	return &CodedError{
		Code: http.StatusInternalServerError,
		Err:  fmt.Errorf("internal error: %w", err),
	}
}
