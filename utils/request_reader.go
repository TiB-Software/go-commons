package utils

import (
	"github.com/TB-Systems/go-commons/errors"
	"github.com/TB-Systems/go-commons/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DecodeJson[T any](ctx *gin.Context) (T, []errors.ApiErrorItem) {
	var data T

	if err := ctx.BindJSON(&data); err != nil {
		return data, []errors.ApiErrorItem{errors.InvalidDecodeJsonError(err.Error())}
	}

	return data, nil
}

func DecodeValidJson[T validator.Validator](ctx *gin.Context) (T, errors.ApiError) {
	data, errs := DecodeJson[T](ctx)

	if len(errs) > 0 {
		return data, errors.NewApiErrorWithErrors(http.StatusBadRequest, errs)
	}

	errs = data.Validate()

	if len(errs) > 0 {
		return data, errors.NewApiErrorWithErrors(http.StatusUnprocessableEntity, errs)
	}

	return data, nil
}
