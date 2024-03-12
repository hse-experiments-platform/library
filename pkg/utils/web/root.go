package web

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/hse-experiments-platform/library/pkg/utils/web/errors"
)

func writeCodedErr(rw http.ResponseWriter, err *errors.CodedError) {
	rw.WriteHeader(err.Code)
	bytes, errM := json.Marshal(err)
	if errM != nil {
		slog.Error("cannot marshal http error", slog.String("err", errM.Error()))
		_, _ = rw.Write([]byte("\"error\""))
		return
	}

	if _, wErr := rw.Write(bytes); wErr != nil {
		slog.Error(wErr.Error())
	}
}

func writeSuccess(rw http.ResponseWriter, msg []byte) {
	rw.WriteHeader(http.StatusOK)
	if _, wErr := rw.Write(msg); wErr != nil {
		slog.Error(wErr.Error())
	}
}

func WithErrorHandler[T any](f func(ctx context.Context, headers http.Header, r *http.Request) (T, error)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		res, err := f(r.Context(), rw.Header(), r)
		rw.Header().Set("Content-Type", "application/json")
		if err != nil {
			slog.Error("http error", slog.String("err", err.Error()))
			writeCodedErr(rw, errors.GetCodedError(err))
			return
		}

		resBytes, mErr := json.Marshal(res)
		if mErr != nil {
			slog.Error("http response marshal error", slog.String("err", err.Error()))
			writeCodedErr(rw, errors.InternalError(mErr))
			return
		}

		writeSuccess(rw, resBytes)
	}
}
