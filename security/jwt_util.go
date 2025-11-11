package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUtil struct {
	Secret       string
	ExpirationMs int64
}

func (j *JwtUtil) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(j.ExpirationMs) * time.Millisecond).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

func (j *JwtUtil) ValidateToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["sub"].(string), nil
	}
	return "", nil
}
