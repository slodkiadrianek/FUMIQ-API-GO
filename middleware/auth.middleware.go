package middleware

import (
	"FUMIQ_API/config"
	"FUMIQ_API/models"
	"FUMIQ_API/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Secret  string
	Logger  utils.Logger
	Caching *config.CacheService
}
type userClaims struct {
	models.User
	exp int64
	jwt.RegisteredClaims
}

func (auth *AuthMiddleware) Sign(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID.Hex(),
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"exp":       time.Now().Add(2 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(auth.Secret))
	if err != nil {
		auth.Logger.Error("error signing token", err)
		return "", err
	}
	return tokenString, nil
}

func (auth *AuthMiddleware) Verify(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		auth.Logger.Error("token is missing during verification of request  with data", c.Request.URL)
		err := models.NewError(401, "Authorization", "token is missing")
		c.Error(err)
		c.Abort()
		return
	}
	tokenString := strings.Split(authHeader, " ")[1]
	result, err := auth.Caching.ExistData(c, "blacklist:"+tokenString)
	if err != nil {
		auth.Logger.Error("error checking blacklist", authHeader)
		err := models.NewError(401, "Authorization", "error occurred during cache checking ")
		c.Error(err)
		c.Abort()
		return
	}
	if result > 0 {
		auth.Logger.Info("token is blacklisted", authHeader)
		err := models.NewError(401, "Authorization", "Token is blacklisted")
		c.Error(err)
		c.Abort()
		return
	}

	var claims userClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.Secret), nil
	})

	if err != nil {
		auth.Logger.Error("token parsing error", err.Error())
		fmt.Println("Token parsing error details:", err.Error())
		err := models.NewError(401, "Authorization", "token parsing error")
		c.Error(err)
		c.Abort()
		return
	}
	if !token.Valid {
		auth.Logger.Error("token is invalid", token)
		err := models.NewError(401, "Authorization", "token is invalid")
		c.Error(err)
		c.Abort()
		return
	}
	c.Set("userId", claims.User.ID.Hex())
	c.Set("firstName", claims.FirstName)
	c.Set("lastName", claims.LastName)
	c.Next()
}

func (auth *AuthMiddleware) BlackList(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		auth.Logger.Error("token is missing during verification of request  with data", c.Request.URL)
		err := models.NewError(401, "Authorization", "token is missing")
		c.Error(err)
		c.Abort()
		return
	}
	tokenString := strings.Split(authHeader, " ")[1]
	var claims userClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.Secret), nil
	})
	if err != nil {
		auth.Logger.Error("token parsing error", err.Error())
		fmt.Println("Token parsing error details:", err.Error())
		err := models.NewError(401, "Authorization", "token parsing error")
		c.Error(err)
		c.Abort()
		return
	}
	if !token.Valid {
		auth.Logger.Error("token is invalid", token)
		err := models.NewError(401, "Authorization", "token is invalid")
		c.Error(err)
		c.Abort()
		return
	}
	err = auth.Caching.SetData(c, "blacklist:"+tokenString, "true", time.Duration(claims.exp))
	if err != nil {
		auth.Logger.Error("error adding token to blacklist", err.Error())
		err := models.NewError(401, "Authorization", "error occurred during cache adding")
		c.Error(err)
	}
	auth.Logger.Info("Token added to blacklist", tokenString)
	c.Next()
}
