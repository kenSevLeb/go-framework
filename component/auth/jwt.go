package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ClaimContextKey = "user_claims"
)

// jwt
type Jwt struct {
	signKey []byte
}

// New return jwtAuth instance
func New(signKey []byte) *Jwt {
	return &Jwt{signKey: signKey}
}

var (
	TokenValidateFailed = errors.New("token validate failed")
)

// 获取Context里的claims
func GetClaimsFromContext(ctx *gin.Context) interface{} {
	claims, _ := ctx.Get(ClaimContextKey)
	return claims
}

// CreateToken 生成token
func (jwtAuth Jwt) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtAuth.signKey)
}

// ParseWithClaims 解析token
func (jwtAuth Jwt) ParseWithClaims(token string, claims jwt.Claims) error {
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtAuth.signKey, nil
	})

	if err != nil {
		return err
	}

	if err := claims.Valid(); err != nil {
		return TokenValidateFailed
	}
	claims = tokenClaims.Claims

	return nil
}

// ParseWithClaims 解析token
func (jwtAuth Jwt) ParseWithClaimsAndKey(token string, claims jwt.Claims, signKey []byte) error {
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	if err != nil {
		return err
	}

	if err := claims.Valid(); err != nil {
		return TokenValidateFailed
	}
	claims = tokenClaims.Claims

	return nil
}
