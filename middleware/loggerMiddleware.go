package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web_starter/util"
	"time"
)

func CostTime() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Next()
		end := time.Since(start)
		msg := fmt.Sprintf("Method: %s Url: %s Cost: %v StatusCode: %d", context.Request.Method, context.Request.URL, end, context.Writer.Status())
		util.Log.Info(msg)
	}
}
