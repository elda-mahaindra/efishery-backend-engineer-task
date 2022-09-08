package api

import (
	"net/http"

	"auth-app/token"

	"github.com/gin-gonic/gin"
)

type VerifyTokenResponse struct {
	CreatedAt int64  `json:"created_at"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

func (server *ApiServer) verifyToken(ctx *gin.Context) {
	serviceName := "VerifyToken"

	// get token payload from context
	payload := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)

	// authorization
	isAuthorized, err := server.authorize(serviceName, payload.Role)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err, "Authorization failed."))
		return
	}

	if !isAuthorized {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err, "Authorization failed."))
		return
	}

	res := VerifyTokenResponse{
		CreatedAt: payload.CreatedAt,
		Name:      payload.Name,
		Phone:     payload.Phone,
		Role:      payload.Role,
	}

	ctx.JSON(http.StatusOK, res)
}
