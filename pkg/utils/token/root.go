package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = fmt.Errorf("token is expired")
	ErrInvalidToken = fmt.Errorf("token is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserID    int64     `json:"username,omitempty"`
	Roles     []string  `json:"roles,omitempty"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID int64, roles []string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        id,
		UserID:    userID,
		Roles:     roles,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, err
}

func (t *Payload) Valid() error {
	if time.Now().After(t.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
