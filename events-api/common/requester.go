package common

import (
	"ankumar/events-api/models"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetRequester(context *gin.Context) (*models.User, error) {
	userData, exists := context.Get("user")
	if !exists {
		return nil, errors.New("Unable to get the requester")
	}
	user := userData.(*models.User)
	return user, nil
}
