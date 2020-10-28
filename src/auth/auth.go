package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/IvanaKrizic/Dashboard/src/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c.Request)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, "Token required.")
			c.Abort()
			return
		}
		if err := repositories.GetAuth(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		token, err := VerifyToken(tokenString)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//fmt.Println(claims["userId"])
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
				repositories.DeleteAuth(tokenString)
			}
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
	}
}

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	return token, err
}
