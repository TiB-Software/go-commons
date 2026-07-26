package validator

import "github.com/TB-Systems/go-commons/errors"

type Validator interface {
	Validate() []errors.ApiErrorItem
}
