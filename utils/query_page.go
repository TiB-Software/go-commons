package utils

import (
	"net/http"

	"github.com/TiB-Software/go-commons/constants"
	"github.com/TiB-Software/go-commons/errors"

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
