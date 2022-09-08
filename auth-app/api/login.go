package api

import (
	"errors"
	"net/http"
	"time"

	"auth-app/model"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type LoginResponse struct {
	TokenString string `json:"token_string"`
}

func (server *ApiServer) login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, "Failed binding validation: invalid request message."))
		return
	}

	found, err := server.userStore.Find(req.Phone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Failed to find user."))
		return
	}
	if found == nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("incorrect phone number or password"), "Incorrect phone number or password."))
		return
	}

	if found.Password != req.Password {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("incorrect phone number or password"), "Incorrect phone number or password."))
		return
	}

	// generating JWT access token
	signedToken, err := server.TokenManager.GenerateToken(func(user *model.User) (string, string, string, int64, time.Duration) {
		return user.Name, user.Phone, user.Role, user.CreatedAt, 1440 * time.Minute
	}(found))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Failed to sign the newly created token"))
		return
	}

	res := LoginResponse{
		TokenString: signedToken,
	}

	ctx.JSON(http.StatusOK, res)
}
