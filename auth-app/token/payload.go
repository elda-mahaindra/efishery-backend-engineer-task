package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("failed to verify token: token has expired")
	ErrInvalidToken = errors.New("failed to verify token: token is invalid")
)

type Payload struct {
	// User Info
	CreatedAt int64  `json:"created_at"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	// Token Info
	ExpiredAt time.Time `json:"expired_at"`
	IssuedAt  time.Time `json:"issued_at"`
	TokenId   uuid.UUID `json:"token_id"`
}

func NewPayload(name, phone, role string, createdAt int64, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		// User Info
		CreatedAt: createdAt,
		Name:      name,
		Phone:     phone,
		Role:      role,
		// Token Info
		ExpiredAt: time.Now().Add(duration),
		IssuedAt:  time.Now(),
		TokenId:   tokenId,
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
