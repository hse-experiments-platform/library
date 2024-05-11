package token

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type Maker interface {
	CreateToken(userID int64, roles []string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
	TokenExtractorUnaryInterceptor() grpc.UnaryServerInterceptor
	TransferToAnotherContext(src, dst context.Context) (context.Context, error)
}
