# go-commons

Shared Go utilities and models used by the Financial Control services.

## Module

```go
module github.com/TB-Systems/go-commons
```

## Installation

```bash
go get github.com/TB-Systems/go-commons@latest
```

If you are using this module inside the same monorepo, prefer `go.work` to link local modules.

## Packages

- `commonsmodels`: generic response and pagination models.
- `errors`: API error contracts and common error factories.
- `utils`: helpers for HTTP requests, query parsing, response writing, date handling, pgx conversions, and generic utilities.
- `constants`: shared constants used by helpers and error builders.
- `validator`: small interface used by request decoding helpers.

## Main Types and Helpers

### commonsmodels

- `PaginatedResponse[T]`
- `ResponseList[T]`
- `ResponseSuccess`
- `PaginatedParams`
- `PaginatedParamsWithDateRange`
- `PaginatedParamsWithMonthYear`
- `NewResponseSuccess()`

### errors

- `ApiError` interface
- `ApiErrorItem`
- `ErrorResponse`
- `NewApiError(status, message)`
- `NewApiErrorWithErrors(status, messages)`
- Common factories:
	- `UserNotFound`
	- `UserIDInvalid`
	- `InvalidDecodeJsonError`
	- `InvalidFieldError`
	- `BadRequestError`
	- `NotFoundError`
	- `InternalServerError`

### utils

- HTTP request/response:
	- `DecodeJson[T]`
	- `DecodeValidJson[T]`
	- `SendResponse[T]`
	- `SendErrorResponse`
- Query and URL parsing:
	- `GetQueryPage`
	- `GetQueryLimit`
	- `GetQueryMonthAndYear`
	- `GetQueryDatesIfHas`
	- `IDFromURLParam`
	- `GetUserIDFromContext`
- Date and pagination:
	- `NormalizeDay`
	- `CreateDateWithNormalizedDay`
	- `CalculateOffset`
- Generic and conversion helpers:
	- `FindIf`
	- `FindIndex`
	- `IsBlank`
	- `StringToInt64`
	- `PgTypeUUIDToUUID`
	- `UUIDToPgTypeUUID`
	- `Float64ToNumeric`
	- `NumericToFloat64`
	- `TimeToPgTimestamptz`

## Quick Usage

### Build a standard paginated response

```go
import "github.com/TB-Systems/go-commons/commonsmodels"

func makeResponse(items []string) commonsmodels.PaginatedResponse[string] {
		return commonsmodels.PaginatedResponse[string]{
				Items:     items,
				PageCount: 1,
				Page:      1,
		}
}
```

### Return API errors in handlers

```go
import (
		"net/http"

		"github.com/TB-Systems/go-commons/errors"
)

func invalidPayload() errors.ApiError {
		return errors.NewApiError(
				http.StatusBadRequest,
				errors.InvalidDecodeJsonError("invalid request body"),
		)
}
```

### Decode and validate requests from Gin context

```go
import (
		"github.com/gin-gonic/gin"

		"github.com/TB-Systems/go-commons/utils"
)

func readRequest(ctx *gin.Context) {
		req, apiErr := utils.DecodeValidJson[MyRequest](ctx)
		if apiErr != nil {
				utils.SendErrorResponse(ctx, apiErr)
				return
		}

		utils.SendResponse(ctx, req, 200)
}
```

## Development

Run tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Versioning and Release

- Use semantic version tags with `v` prefix (`v1.0.0`, `v1.0.1`, ...).
- Push tags to make versions resolvable by `go get` and `go mod tidy`:

```bash
git tag vX.Y.Z
git push origin vX.Y.Z
```
