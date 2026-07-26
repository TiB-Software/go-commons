package errors

import (
	"github.com/TB-Systems/go-commons/constants"
	"testing"
)

func TestUserNotFound(t *testing.T) {
	detail := "cookie not found"
	result := UserNotFound(detail)

	if result.UserMessage != constants.UserUnauthorized {
		t.Errorf("Expected UserMessage '%s', got '%s'", constants.UserUnauthorized, result.UserMessage)
	}
	if result.SystemMessage != constants.UserIDNotFound {
		t.Errorf("Expected SystemMessage '%s', got '%s'", constants.UserIDNotFound, result.SystemMessage)
	}
	if result.SystemDetail != detail {
		t.Errorf("Expected SystemDetail '%s', got '%s'", detail, result.SystemDetail)
	}
}

func TestUserIDInvalid(t *testing.T) {
	detail := "invalid uuid format"
	result := UserIDInvalid(detail)

	if result.UserMessage != constants.UserUnauthorized {
		t.Errorf("Expected UserMessage '%s', got '%s'", constants.UserUnauthorized, result.UserMessage)
	}
	if result.SystemMessage != constants.UserIDInvalid {
		t.Errorf("Expected SystemMessage '%s', got '%s'", constants.UserIDInvalid, result.SystemMessage)
	}
	if result.SystemDetail != detail {
		t.Errorf("Expected SystemDetail '%s', got '%s'", detail, result.SystemDetail)
	}
}

func TestInvalidDecodeJsonError(t *testing.T) {
	detail := "unexpected end of JSON input"
	result := InvalidDecodeJsonError(detail)

	if result.UserMessage != constants.InvalidData {
		t.Errorf("Expected UserMessage '%s', got '%s'", constants.InvalidData, result.UserMessage)
	}
	if result.SystemMessage != constants.DecodeJsonError {
		t.Errorf("Expected SystemMessage '%s', got '%s'", constants.DecodeJsonError, result.SystemMessage)
	}
	if result.SystemDetail != detail {
		t.Errorf("Expected SystemDetail '%s', got '%s'", detail, result.SystemDetail)
	}
}

func TestInvalidFieldError(t *testing.T) {
	message := "name is required"
	result := InvalidFieldError(message)

	if result.UserMessage != message {
		t.Errorf("Expected UserMessage '%s', got '%s'", message, result.UserMessage)
	}
	if result.SystemMessage != constants.InvalidFieldError {
		t.Errorf("Expected SystemMessage '%s', got '%s'", constants.InvalidFieldError, result.SystemMessage)
	}
	if result.SystemDetail != message {
		t.Errorf("Expected SystemDetail '%s', got '%s'", message, result.SystemDetail)
	}
}

func TestBadRequestError(t *testing.T) {
	detail := "invalid request parameters"
	result := BadRequestError(detail)

	if result.UserMessage != detail {
		t.Errorf("Expected UserMessage '%s', got '%s'", detail, result.UserMessage)
	}
	if result.SystemMessage != constants.BadRequestError {
		t.Errorf("Expected SystemMessage '%s', got '%s'", constants.BadRequestError, result.SystemMessage)
	}
	if result.SystemDetail != detail {
		t.Errorf("Expected SystemDetail '%s', got '%s'", detail, result.SystemDetail)
	}
}

func TestNotFoundError(t *testing.T) {
	message := "category not found"
	result := NotFoundError(message)

	if result.UserMessage != message {
		t.Errorf("Expected UserMessage '%s', got '%s'", message, result.UserMessage)
	}
	if result.SystemMessage != constants.NotFoundError {
		t.Errorf("Expected SystemMessage '%s', got '%s'", constants.NotFoundError, result.SystemMessage)
	}
	if result.SystemDetail != constants.StoreErrorNoRowsMsg {
		t.Errorf("Expected SystemDetail '%s', got '%s'", constants.StoreErrorNoRowsMsg, result.SystemDetail)
	}
}

func TestInternalServerError(t *testing.T) {
	detail := "database connection failed"
	result := InternalServerError(detail)

	if result.UserMessage != constants.InternalServerError {
		t.Errorf("Expected UserMessage '%s', got '%s'", constants.InternalServerError, result.UserMessage)
	}
	if result.SystemMessage != constants.InternalServerError {
		t.Errorf("Expected SystemMessage '%s', got '%s'", constants.InternalServerError, result.SystemMessage)
	}
	if result.SystemDetail != detail {
		t.Errorf("Expected SystemDetail '%s', got '%s'", detail, result.SystemDetail)
	}
}

func TestErrorFunctionsReturnApiErrorItem(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(string) ApiErrorItem
		input    string
		expected ApiErrorItem
	}{
		{
			name:  "UserNotFound",
			fn:    UserNotFound,
			input: "test detail",
			expected: ApiErrorItem{
				UserMessage:   constants.UserUnauthorized,
				SystemMessage: constants.UserIDNotFound,
				SystemDetail:  "test detail",
			},
		},
		{
			name:  "UserIDInvalid",
			fn:    UserIDInvalid,
			input: "test detail",
			expected: ApiErrorItem{
				UserMessage:   constants.UserUnauthorized,
				SystemMessage: constants.UserIDInvalid,
				SystemDetail:  "test detail",
			},
		},
		{
			name:  "InvalidDecodeJsonError",
			fn:    InvalidDecodeJsonError,
			input: "test detail",
			expected: ApiErrorItem{
				UserMessage:   constants.InvalidData,
				SystemMessage: constants.DecodeJsonError,
				SystemDetail:  "test detail",
			},
		},
		{
			name:  "InvalidFieldError",
			fn:    InvalidFieldError,
			input: "test message",
			expected: ApiErrorItem{
				UserMessage:   "test message",
				SystemMessage: constants.InvalidFieldError,
				SystemDetail:  "test message",
			},
		},
		{
			name:  "BadRequestError",
			fn:    BadRequestError,
			input: "test detail",
			expected: ApiErrorItem{
				UserMessage:   "test detail",
				SystemMessage: constants.BadRequestError,
				SystemDetail:  "test detail",
			},
		},
		{
			name:  "InternalServerError",
			fn:    InternalServerError,
			input: "test detail",
			expected: ApiErrorItem{
				UserMessage:   constants.InternalServerError,
				SystemMessage: constants.InternalServerError,
				SystemDetail:  "test detail",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %+v, got %+v", tc.expected, result)
			}
		})
	}
}
