package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kenSevLeb/go-framework/component/auth"
	"github.com/kenSevLeb/go-framework/http/response"
	"strings"
)

var (
	TokenNotExist = fmt.Errorf("token not exist")
)

// validateToken 验证token
func validateToken(jwtAuth *auth.Jwt, ctx *gin.Context) error {
	// 获取token
	tokenStr := ctx.GetHeader("Authorization")
	if strings.TrimSpace(tokenStr) == "" {
		return TokenNotExist
	}

	// parse claims
	claims := new(jwt.StandardClaims)
	if err := jwtAuth.ParseWithClaims(tokenStr, claims); err != nil {
		return err
	}

	// token存入context
	ctx.Set(auth.ClaimContextKey, claims)
	return nil
}

// JWT gin的jwt中间件
func JWT(jwtAuth *auth.Jwt) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 解析token
		if err := validateToken(jwtAuth, ctx); err != nil {
			response.Wrap(ctx).Error(401, fmt.Sprintf("InvalidToken:%v", err))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
