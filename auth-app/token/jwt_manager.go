package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTManager is a JSON web token manager
type JWTManager struct {
	signatureKey  string
	signingMethod jwt.SigningMethod
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(secret string) (Manager, error) {
	manager := &JWTManager{
		signatureKey:  secret,
		signingMethod: jwt.SigningMethodHS256,
	}

	return manager, nil
}

// Generate generates and signs a new token for a user
func (manager *JWTManager) GenerateToken(name, phone, role string, createdAt int64, duration time.Duration) (string, error) {
	payload, err := NewPayload(name, phone, role, createdAt, duration)
	if err != nil {
		return "", err
	}

	// creating new token
	token := jwt.NewWithClaims(manager.signingMethod, payload)

	// signing the token using secret key
	return token.SignedString([]byte(manager.signatureKey))
}

// Verify verifies the token string and return the jwt token
func (manager *JWTManager) VerifyToken(tokenString string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(manager.signatureKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return payload, err
}
