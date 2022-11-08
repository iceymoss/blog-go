package middlewares

import (
	"errors"
	"fmt"

	"blog-go/blog_web/user_web/models"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWT struct {
	singer []byte
}

func NewJWT() *JWT {
	return &JWT{
		singer: []byte("ice_moss"),
	}
}

//GenerateToken 生成token
func (j *JWT) GenerateToken(claims *models.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, *claims)
	fmt.Println(token)
	newTokent, err := token.SignedString(j.singer)
	if err != nil {
		zap.S().Info("签名失败", err)
	}
	return newTokent, nil
}

//ParseToken 解析token
func (j *JWT) ParseToken(token string) (*models.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, models.Claims{}, func(key *jwt.Token) (interface{}, error) {
		return j.singer, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if tokenClaims != nil {
		//断言
		if claims, ok := tokenClaims.Claims.(models.Claims); ok && tokenClaims.Valid {
			return &claims, nil
		}

		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}
