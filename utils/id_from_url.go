package utils

import (
	"net/http"

	"github.com/TiB-Software/go-commons/constants"
	"github.com/TiB-Software/go-commons/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IDFromURLParam(ctx *gin.Context) (uuid.UUID, errors.ApiError) {
	idString := ctx.Param(constants.ID)

	id, err := uuid.Parse(idString)

	if err != nil {
		return uuid.UUID{}, errors.NewApiError(http.StatusBadRequest, errors.BadRequestError(constants.InvalidID))
	}

	return id, nil
}
