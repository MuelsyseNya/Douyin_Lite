package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	UserID int64
	jwt.StandardClaims
}

// SignKey JWT签名密钥
const SignKey = "This_is_my_signed_key"

// GenerateToken 生成JWT Token
func GenerateToken(userID int64) (signedToken string, err error) {
	claims := CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			NotBefore: time.Now().Unix() - 300,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对Token进行签名
	secretKey := []byte(SignKey)
	signedToken, err = token.SignedString(secretKey)
	return signedToken, err
}
