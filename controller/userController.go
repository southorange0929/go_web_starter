package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_starter/dao"
	"go_web_starter/util"
	"net/http"
)

type UserController struct {
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	user := dao.NewPersonDao().GetPerson()
	util.Log.Info(user)
	ctx.JSON(http.StatusOK,user)
}
