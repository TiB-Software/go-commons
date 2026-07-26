package utils

import (
	"net/http"
	"time"

	"github.com/TiB-Software/go-commons/constants"
	"github.com/TiB-Software/go-commons/errors"

	"github.com/gin-gonic/gin"
)

func GetQueryDatesIfHas(ctx *gin.Context) (time.Time, time.Time, bool, errors.ApiError) {
	startDateString := ctx.DefaultQuery(constants.StartDateText, constants.EmptyString)
	endDateString := ctx.DefaultQuery(constants.EndDateText, constants.EmptyString)

	if IsBlank(startDateString) && IsBlank(endDateString) {
		return time.Time{}, time.Time{}, false, nil
	}

	startDate, err := time.Parse(time.DateOnly, startDateString)

	if err != nil {
		return time.Time{}, time.Time{}, false, errors.NewApiError(http.StatusBadRequest, errors.BadRequestError(constants.InvalidStartDate))
	}

	endDate, err := time.Parse(time.DateOnly, endDateString)

	if err != nil {
		return time.Time{}, time.Time{}, false, errors.NewApiError(http.StatusBadRequest, errors.BadRequestError(constants.InvalidEndDate))
	}

	return startDate, endDate, true, nil
}
