package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type maker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("key length must be %v", chacha20poly1305.KeySize)
	}

	return &maker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (m *maker) CreateToken(userID int64, roles []string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, roles, duration)
	if err != nil {
		return "", err
	}

	return m.paseto.Encrypt(m.symmetricKey, payload, nil)
}

func (m *maker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
