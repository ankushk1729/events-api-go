package routes

import (
	"ankumar/events-api/common"
	"ankumar/events-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while registering for the event."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Failed to find the event."})
		return
	}

	user, err := common.GetRequester(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to register for the event."})
		return
	}
	err = event.Register(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while registering for the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for the event."})
}

func cancelRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while registering for the event."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Failed to find the event."})
		return
	}

	user, err := common.GetRequester(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to register for the event."})
	}
	err = event.CancelRegistration(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while cancelling registration for the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration for the event."})
}
