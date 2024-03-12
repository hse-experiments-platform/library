package errors

import "errors"

type CodedError struct {
	Code int   `json:"code"`
	Err  error `json:"err"`
}

func NewCodedError(code int, err error) *CodedError {
	return &CodedError{
		Code: code,
		Err:  err,
	}
}

func (e *CodedError) Error() string {
	return e.Err.Error()
}

func (e *CodedError) Unwrap() error {
	return e.Err
}

func GetCodedError(err error) *CodedError {
	if err == nil {
		return nil
	}

	var codedErr *CodedError

	if errors.As(err, &codedErr) {
		return codedErr
	}

	return InternalError(err)
}
