package constants

// COMMONS
const (
	EmptyString        = ""
	Success            = "success"
	ID                 = "id"
	LimitText          = "limit"
	LimitDefaultString = "10"
	LimitDefault       = 10
	PageText           = "page"
	PageDefaultString  = "1"
	PageDefault        = 1
	StartDateText      = "start_date"
	EndDateText        = "end_date"
	MonthText          = "month"
	YearText           = "year"
	UserID             = "user_id"
)

// STORE ERRORS
const (
	StoreErrorNoRowsMsg = "no rows in result set"
)

// ERRORS_USER_MESSAGES
const (
	UserUnauthorized                    = "USER_UNAUTHORIZED"
	InvalidStartDate                    = "INVALID_START_DATE"
	InvalidEndDate                      = "INVALID_END_DATE"
	EndDateBeforeStartDate              = "END_DATE_BEFORE_START_DATE"
	TransactionTypeMsg                  = "TRANSACTION_TYPE_INVALID"
	TransactionTypeEmptyMsg             = "TRANSACTION_TYPE_EMPTY"
	NameEmptyMsg                        = "NAME_EMPTY"
	IconEmptyMsg                        = "ICON_EMPTY"
	NameInvalidCharsCountMsg            = "NAME_INVALID_CHARS_COUNT"
	IconInvalidCharsCountMsg            = "ICON_INVALID_CHARS_COUNT"
	LimitReachedMsg                     = "LIMIT_REACHED"
	CannotBeDeletedMsg                  = "CANNOT_BE_DELETED_BECAUSE_IT_HAS_ASSOCIATED_TRANSACTIONS"
	ValueInvalidMsg                     = "VALUE_INVALID"
	DateEmptyMsg                        = "DATE_EMPTY_OR_INVALID"
	DateInvalidMsg                      = "DATE_INVALID"
	CreditcardLimitExceededMsg          = "CREDITCARD_LIMIT_EXCEEDED"
	InvalidData                         = "INVALID_DATA"
	InvalidID                           = "INVALID_ID"
	CategoryNotFoundMsg                 = "CATEGORY_NOT_FOUND"
	CreditcardNotFoundMsg               = "CREDITCARD_NOT_FOUND"
	TransactionNotFoundMsg              = "TRANSACTION_NOT_FOUND"
	MonthlyTransactionNotFoundMsg       = "MONTHLY_TRANSACTION_NOT_FOUND"
	AnnualTransactionNotFoundMsg        = "ANNUAL_TRANSACTION_NOT_FOUND"
	InstallmentTransactionNotFoundMsg   = "INSTALLMENT_TRANSACTION_NOT_FOUND"
	DayInvalidMsg                       = "DAY_INVALID"
	MonthInvalidMsg                     = "MONTH_INVALID"
	YearInvalidMsg                      = "YEAR_INVALID"
	YearMustBe1970OrLaterMsg            = "YEAR_MUST_BE_1970_OR_LATER"
	InitialDateEmptyMsg                 = "INITIAL_DATE_EMPTY"
	FinalDateEmptyMsg                   = "FINAL_DATE_EMPTY"
	FinalDateBeforeInitialDateMsg       = "FINAL_DATE_BEFORE_INITIAL_DATE"
	InitialDateEqualsFinalDateMsg       = "INITIAL_DATE_EQUALS_FINAL_DATE"
	AnnualAndInstallmentTransactionMsg  = "ANNUAL_AND_INSTALLMENT_TRANSACTION"
	AnnualAndMonthlyTransactionMsg      = "ANNUAL_AND_MONTHLY_TRANSACTION"
	InstallmentAndMonthlyTransactionMsg = "INSTALLMENT_AND_MONTHLY_TRANSACTION"
)

// ERRORS
const (
	UserIDNotFound       = "USER_ID_NOT_FOUND"
	UserIDInvalid        = "USER_ID_INVALID"
	InternalServerError  = "INTERNAL_SERVER_ERROR"
	NilValueError        = "NIL_VALUE"
	UnsupportedTypeError = "UNSUPPORTED_TYPE"
	DecodeJsonError      = "DECODE_JSON_ERROR"
	EncodeJsonError      = "ENCODE_JSON_ERROR"
	InvalidFieldError    = "INVALID_FIELD_ERROR"
	LimitError           = "LIMIT_ERROR"
	NotFoundError        = "NOT_FOUND_ERROR"
	StoreError           = "STORE_ERROR"
	UnauthorizedError    = "UNAUTHORIZED_ERROR"
	InvalidPageParam     = "INVALID_PAGE_PARAM"
	CustomError          = "CUSTOM_ERROR"
	BadRequestError      = "BAD_REQUEST_ERROR"
)
