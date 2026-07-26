package commonsmodels

import (
	"time"

	"github.com/google/uuid"
)

type PaginatedParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
	Page   int32
}

type PaginatedParamsWithDateRange struct {
	UserID    uuid.UUID
	Limit     int32
	Offset    int32
	Page      int32
	StartDate time.Time
	EndDate   time.Time
}

type PaginatedParamsWithMonthYear struct {
	UserID     uuid.UUID
	Year       int32
	Month      int32
	Page       int32
	PageOffset int32
	PageLimit  int32
}
