package token

import "time"

type Manager interface {
	GenerateToken(name, phone, role string, createdAt int64, duration time.Duration) (string, error)
	VerifyToken(tokenString string) (*Payload, error)
}
