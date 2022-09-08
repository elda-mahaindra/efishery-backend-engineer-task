package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message string `json:"message"`
}

func (server *ApiServer) ping(ctx *gin.Context) {
	res := PingResponse{
		Message: "PONG",
	}
	ctx.JSON(http.StatusOK, res)
}
