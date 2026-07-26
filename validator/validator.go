package validator

import "github.com/TiB-Software/go-commons/errors"

type Validator interface {
	Validate() []errors.ApiErrorItem
}
