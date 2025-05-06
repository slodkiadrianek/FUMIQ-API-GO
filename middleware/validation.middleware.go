package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateRequestData[T any](c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": "Failed to read request body"})
		c.Abort()
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	fmt.Println(string(bodyBytes))

	var data T
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": err.Error()})
		c.Abort()
		return
	}
	fmt.Println(data)
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	var errorStack []string

	for i := 0; i < typ.NumField(); i++ {
		options := strings.Split(string(typ.Field(i).Tag.Get("opts")), ",")
		field := val.Field(i)
		var value string
		fieldName := typ.Field(i).Name
		if field.Kind() == reflect.String {
			value = field.String()
		} else {
			continue
		}
		fmt.Println(options)
		for _, option := range options {
			var optionValueInt int
			var optionValueStr string

			if strings.Contains(option, "min") {
				optionValueInt, err = strconv.Atoi(strings.Split(option, "=")[1])
				option = strings.Split(option, "=")[0]
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": err.Error()})
					c.Abort()
					return
				}
			}
			if strings.Contains(option, "max") {
				fmt.Println(option)
				optionValueInt, err = strconv.Atoi(strings.Split(option, "=")[1])
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": err.Error()})
					c.Abort()
					return
				}
				option = strings.Split(option, "=")[0]
			}
			if strings.Contains(option, "confirm") {
				optionValueStr = strings.Split(option, "=")[1]
				option = strings.Split(option, "=")[0]
			}
			if strings.Contains(option, "regex") {
				optionValueStr = strings.Split(option, "=")[1]
				option = strings.Split(option, "=")[0]
			}

			switch option {
			case "required":
				if val.Field(i).IsZero() {
					errorStack = append(errorStack, fmt.Sprintf("%s is required", fieldName))
				}
				fmt.Println(typ.Field(i).Name)
			case "min":
				if len(value) < optionValueInt {
					errorStack = append(errorStack, fmt.Sprintf("%s must be at least %d characters", fieldName, optionValueInt))
				}
			case "max":
				if len(value) > optionValueInt {
					errorStack = append(errorStack, fmt.Sprintf("%s can have maximum %d characters", fieldName, optionValueInt))
				}
			case "email":
				formattedMail := strings.Trim(strings.ToLower(value), " ")
				r, _ := regexp.Compile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
				if !r.MatchString(formattedMail) {
					errorStack = append(errorStack, fmt.Sprintf("%s must be a valid email", fieldName))
				}

			case "confirm":
				fmt.Println(optionValueStr)
				confirmValue := val.FieldByName(optionValueStr).String()
				if value != confirmValue {
					errorStack = append(errorStack, fmt.Sprintf("%s value must match with %s", fieldName, optionValueStr))
				}
			case "regex":
				r, _ := regexp.Compile(optionValueStr)
				if !r.MatchString(value) {
					errorStack = append(errorStack, fmt.Sprintf("%s must pass regex check", fieldName))
				}
			}
		}
	}
	if len(errorStack) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": errorStack})
		c.Abort()
		return
	}
	//}
	c.Next()
}
