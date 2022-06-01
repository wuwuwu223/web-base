package utils

//jwt

import (
	"github.com/gin-gonic/gin"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

// CreateToken creates a new token
func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SecretKey))
}

// ParseToken parses token
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid checks if token is valid
func TokenValid(tokenString string) error {
	token, err := ParseToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

//jwt auth middleware
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "please login",
			})
			c.Abort()
			return
		}
		if err := TokenValid(token); err != nil {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "token is not valid",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
