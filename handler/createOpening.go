package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtheusvalle/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Error("Request validation error: ", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Error("Error creating opening", err)
		sendError(ctx, http.StatusInternalServerError, "Error creating opening")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
