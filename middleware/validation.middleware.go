package middleware

import (
	"FUMIQ_API/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"reflect"
)

func ValidateRequestData[T any](c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": "Failed to read request body"})
		c.Abort()
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var data T
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": err.Error()})
		c.Abort()
		return
	}

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fieldPass := val.FieldByName("Password")
	fieldConPassword := val.FieldByName("ConfirmPassword")
	fieldNewPassword := val.FieldByName("NewPassword")
	fmt.Println(fieldPass)
	fmt.Println(fieldConPassword)
	errorDescription := ""

	if fieldPass.IsValid() {
		if fieldConPassword.IsValid() {
			if err := utils.ArePasswordsSimilar(fieldPass.String(), fieldConPassword.String()); err != nil {
				errorDescription = err.Error()
			}
		}
		if err := utils.RegexCheck(fieldPass.String(), "^[A-Za-z\\d@$!%*?&]+$"); err != nil {
			errorDescription = err.Error()
		}
		if fieldNewPassword.IsValid() {
			if err := utils.RegexCheck(fieldNewPassword.String(), "^[A-Za-z\\d@$!%*?&]{8,}$"); err != nil {
				errorDescription = err.Error()
			}
		}
	}

	if errorDescription != "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": errorDescription})
		c.Abort()
		return
	}

	c.Next()
}
