package api

import (
	"auth-app/util"

	"github.com/go-playground/validator/v10"
)

var validRole validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if role, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedRole(role)
	}
	return false
}
