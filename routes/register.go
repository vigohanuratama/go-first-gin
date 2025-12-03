package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not find event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not register event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event registered"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not cancel registration"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event cancelled"})
}
