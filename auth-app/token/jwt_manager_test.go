package token

import (
	"testing"
	"time"

	"auth-app/util"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestJWTManager(t *testing.T) {
	manager, err := NewJWTManager(util.RandomString(32))
	require.NoError(t, err)

	name := util.RandomName()
	phone := util.RandomPhone()
	role := "admin"
	createdAt := time.Now().UnixNano()
	duration := 1440 * time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := manager.GenerateToken(name, phone, role, createdAt, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := manager.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.TokenId)
	require.Equal(t, name, payload.Name)
	require.Equal(t, phone, payload.Phone)
	require.Equal(t, role, payload.Role)
	require.Equal(t, createdAt, payload.CreatedAt)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	manager, err := NewJWTManager(util.RandomString(32))
	require.NoError(t, err)

	name := util.RandomName()
	phone := util.RandomPhone()
	role := "admin"
	createdAt := time.Now().UnixNano()

	token, err := manager.GenerateToken(name, phone, role, createdAt, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := manager.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	name := util.RandomName()
	phone := util.RandomPhone()
	role := "admin"
	createdAt := time.Now().UnixNano()

	payload, err := NewPayload(name, phone, role, createdAt, time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	manager, err := NewJWTManager(util.RandomString(32))
	require.NoError(t, err)

	payload, err = manager.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
