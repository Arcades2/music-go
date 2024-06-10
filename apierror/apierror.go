package apierror

import "github.com/gin-gonic/gin"

func SendErrorResp(statusCode int, c *gin.Context, errors ...error) {
	errorMessages := []string{}

	for _, err := range errors {
		errorMessages = append(errorMessages, err.Error())
	}

	c.JSON(statusCode, gin.H{"errors": errorMessages})
}
