package device

import (
	"github.com/go-kit/kit/log"

	"errors"
	"strings"
)

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
}

type stringService struct{}

// New creates a new string service
func New(logger log.Logger) StringService {
	var svc StringService
	svc = stringService{}
	svc = LoggingMiddleware{logger, svc}
	return svc
}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
