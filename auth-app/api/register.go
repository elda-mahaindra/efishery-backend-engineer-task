package api

import (
	"net/http"
	"time"

	"auth-app/model"
	"auth-app/util"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Role  string `json:"role" binding:"required,role"`
}

type RegisterResponse struct {
	User *model.User `json:"user"`
}

func (server *ApiServer) register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, "Failed binding validation: invalid request message."))
		return
	}

	// write to json file implementation
	createdAt := time.Now().UnixNano()
	password := util.RandomPassword()

	user := &model.User{
		CreatedAt: createdAt,
		Name:      req.Name,
		Password:  password,
		Phone:     req.Phone,
		Role:      req.Role,
	}

	server.userStore.Save(user)

	res := RegisterResponse{
		User: user,
	}

	ctx.JSON(http.StatusOK, res)
}
