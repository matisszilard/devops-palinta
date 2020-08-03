package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

// MakeGetUsersEndpoint creates the get users endpoint
func MakeGetUsersEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		d, err := svc.GetUsers()
		if err != nil {
			return nil, err
		}
		return d, nil
	}
}

// EncodeResponse encodes the response into JSON.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeGetUsersRequest decodes the request from the body.
func DecodeGetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
