package middleware

import (
	"FUMIQ_API/config"
	"FUMIQ_API/models"
	"FUMIQ_API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type AuthMiddleware struct {
	Secret  string
	Logger  utils.Logger
	Caching *config.CacheService
}

func (auth *AuthMiddleware) Sign(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    user.ID.Hex(),
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
	fmt.Println("Token", tokenString)
	return tokenString, nil
}

func (auth *AuthMiddleware) Verify(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		auth.Logger.Error("token is missing during verification of request  with data", c.Request.URL)
		err := models.NewError(401, "Authorization", "token is missing")
		c.Error(err)
	}
	result, err := auth.Caching.ExistData(c, "blacklist:"+authHeader)
	if err != nil {
		auth.Logger.Error("error checking blacklist", authHeader)
		err := models.NewError(401, "Authorization", "error occurred during cache checking ")
		c.Error(err)
	}
	if result > 0 {
		auth.Logger.Info("token is blacklisted", authHeader)
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": gin.H{
			"category":    "Authorization",
			"description": "Token is blacklisted",
		}})
	}

	token, err := jwt.Parse(strings.Split(authHeader, " ")[1], func(token *jwt.Token) (interface{}, error) {
		return auth.Secret, nil
	})
	if err != nil {
		auth.Logger.Error("token parsing error", err)
		err := models.NewError(401, "Authorization", "token parsing error ")
		c.Error(err)
	}
	if !token.Valid {
		auth.Logger.Error("token is invalid", token)
		err := models.NewError(401, "Authorization", "token is invalid")
		c.Error(err)
	}

	c.Next()
}

func (auth *AuthMiddleware) BlackList(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		auth.Logger.Error("token is missing during verification of request  with data", c.Request.URL)
		err := models.NewError(401, "Authorization", "token is missing")
		c.Error(err)
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(strings.Split(authHeader, " ")[1], claims, func(token *jwt.Token) (interface{},
		error) {
		return auth.Secret, nil
	})
	if err != nil {
		auth.Logger.Error("token parsing error", err)
		err := models.NewError(401, "Authorization", "token parsing error ")
		c.Error(err)
	}
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	c.Next()
}
