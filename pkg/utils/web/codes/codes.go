package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var CodesMap = [...]int{
	codes.Unknown:            http.StatusInternalServerError,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.Canceled:           499,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.Internal:           http.StatusInternalServerError,
	codes.OutOfRange:         http.StatusRequestedRangeNotSatisfiable,
	codes.Aborted:            http.StatusConflict,
	codes.Unavailable:        http.StatusServiceUnavailable,
}
