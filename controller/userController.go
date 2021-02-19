package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "user")
}
