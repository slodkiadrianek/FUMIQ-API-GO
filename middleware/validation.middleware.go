package middleware

import (
	"FUMIQ_API/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func ValidateRequestData[T interface{}](c *gin.Context) {
	var data T
	if err := c.ShouldBind(&data); err != nil {
		errorDescription := err.Error()
		val := reflect.ValueOf(data)
		fieldPass := val.FieldByName("Password")
		fieldConPassword := val.FieldByName("ConfirmPassword")
		fieldNewPassword := val.FieldByName("NewPassword")
		if fieldPass.IsValid() {
			if fieldConPassword.IsValid() {
				err := utils.ArePasswordsSimilar(fieldPass.String(), fieldConPassword.String())
				if err != nil {
					errorDescription = err.Error()
				}
			}
			err := utils.RegexCheck(fieldPass.String(), "/^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$/")
			if err != nil {
				errorDescription = err.Error()
			}
			err = utils.RegexCheck(fieldNewPassword.String(), "/^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$/")
			if err != nil {
				errorDescription = err.Error()
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": errorDescription})
		return
	}
	c.Next()
}
