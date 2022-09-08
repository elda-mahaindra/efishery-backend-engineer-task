package sample

import (
	"auth-app/model"
	"auth-app/util"
)

// NewUser returns a new sample user
func NewUser(role string) *model.User {
	user := &model.User{
		Name:  util.RandomName(),
		Phone: util.RandomPhone(),
		Role:  role,
	}

	if !util.IsSupportedRole(role) {
		user.Role = "admin"
	}

	return user
}

// NewInvalidUser returns a new invalid sample user
func NewInvalidUser() *model.User {
	user := &model.User{
		Name:  util.RandomName(),
		Phone: util.RandomPhone(),
		Role:  "invalid-role-?",
	}

	return user
}
