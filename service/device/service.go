package device

import (
	"github.com/go-kit/kit/log"
	"github.com/matisszilard/devops-palinta/pkg/model"

	"errors"
	"strings"
)

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	GetDevices() ([]model.Device, error)
}

type stringService struct{}

// New creates a new string service
func New(logger log.Logger) StringService {
	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}
	return svc
}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) GetDevices() ([]model.Device, error) {
	// Return hardcoded values
	return []model.Device{{Name: "Aragorn"}, {Name: "Legolas"}, {Name: "Gimli"}}, nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
