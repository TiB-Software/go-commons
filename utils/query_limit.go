package utils

import (
	"github.com/TB-Systems/go-commons/constants"

	"github.com/gin-gonic/gin"
)

func GetQueryLimit(ctx *gin.Context) int32 {
	limitString := ctx.DefaultQuery(constants.LimitText, constants.LimitDefaultString)
	limit, err := StringToInt64(limitString)

	if limit > constants.LimitDefault || err != nil {
		limit = constants.LimitDefault
	}

	return int32(limit)
}
