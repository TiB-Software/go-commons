package utils

import (
	"github.com/TB-Systems/go-commons/constants"
	"github.com/TB-Systems/go-commons/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQueryPage(ctx *gin.Context) (int32, errors.ApiError) {
	pageString := ctx.DefaultQuery(constants.PageText, constants.PageDefaultString)
	page, err := StringToInt64(pageString)

	if err != nil {
		return 0, errors.NewApiError(http.StatusBadRequest, errors.BadRequestError(constants.InvalidPageParam))
	}

	if page == 0 {
		page = 1
	}

	return int32(page), nil
}
