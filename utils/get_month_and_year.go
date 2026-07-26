package utils

import (
	"github.com/TB-Systems/go-commons/constants"
	"github.com/TB-Systems/go-commons/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQueryMonthAndYear(ctx *gin.Context) (int32, int32, errors.ApiError) {
	monthString := ctx.DefaultQuery(constants.MonthText, constants.EmptyString)
	month, err := StringToInt64(monthString)

	if err != nil {
		return 0, 0, errors.NewApiError(
			http.StatusBadRequest,
			errors.BadRequestError(constants.MonthInvalidMsg),
		)
	}

	yearString := ctx.DefaultQuery(constants.YearText, constants.EmptyString)
	year, err := StringToInt64(yearString)

	if err != nil {
		return 0, 0, errors.NewApiError(
			http.StatusBadRequest,
			errors.BadRequestError(constants.YearInvalidMsg),
		)
	}

	if month < 1 || month > 12 {
		return 0, 0, errors.NewApiError(
			http.StatusBadRequest,
			errors.BadRequestError(constants.MonthInvalidMsg),
		)
	}

	if year < 1970 {
		return 0, 0, errors.NewApiError(
			http.StatusBadRequest,
			errors.BadRequestError(constants.YearMustBe1970OrLaterMsg),
		)
	}

	return int32(month), int32(year), nil
}
