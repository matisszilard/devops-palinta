package user

import (
	"github.com/matisszilard/devops-palinta/pkg/model"
	"github.com/sirupsen/logrus"

	"errors"
)

// UserService ..
type UserService interface {
	GetUsers() ([]model.User, error)
}

type userService struct{}

// New creates a new string service
func New(logger *logrus.Logger) UserService {
	var svc UserService
	svc = userService{}
	svc = loggingMiddleware{logger, svc}
	return svc
}

func (userService) GetUsers() ([]model.User, error) {
	// Return hardcoded values
	return []model.User{{Name: "Aragorn"}, {Name: "Legolas"}, {Name: "Gimli"}}, nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
