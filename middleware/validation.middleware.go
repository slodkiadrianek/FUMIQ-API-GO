package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	z "github.com/Oudwins/zog"
	"github.com/gin-gonic/gin"
)

type validate interface {
	Validate() (z.ZogIssueMap, error)
}

func ValidateRequestData[T validate](source string) gin.HandlerFunc {
	return func(c *gin.Context) {
		processBody := func() {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": "Failed to read request body"})
				c.Abort()
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			var data T
			if err = json.Unmarshal(bodyBytes, &data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": err.Error()})
				c.Abort()
				return
			}
			errMaps, err := data.Validate()
			if errMaps != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "Category": "Validation", "description": errMaps["$root"]})
			}
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "Category": "Validation", "description": err})
			}
		}
		switch source {
		case "params":
			var data T
			err := c.ShouldBindUri(&data)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Validation Error", "description": "Failed to read request body"})
				c.Abort()
				return
			}
		case "body":
			processBody()
		default:
			processBody()
		}
		c.Next()
	}
}
