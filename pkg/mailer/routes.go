package mailer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContactUsRouter(context *gin.Context) {
	var payload ContactUsPayload

	err := context.BindJSON(&payload)
	if err != nil {
		context.Error(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request schema"})
		return
	}

	SendEmail(payload)

	context.Status(http.StatusNoContent)
}
