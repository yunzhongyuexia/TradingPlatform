package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/tools"
)

// JWTMiddleware 是 Gin 中间件，用于验证请求中的 JWT
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 尝试从请求的 Header 中获取 Token
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			ctx.Abort()
			return
		}

		// 解析 Token 并获取 Claims
		claims, err := tools.ParseToken(token)
		if err != nil || claims == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// 在这里可以进行进一步的逻辑处理，比如根据 Claims 中的信息
		// 设置上下文、调用其他服务等。

		// 继续处理请求链
		ctx.Next()
	}
}
