package device

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

// MakeUppercaseEndpoint creates the uppercase endpoint
func MakeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

// MakeGetDevicesEndpoint creates the get devices endpoint
func MakeGetDevicesEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		d, err := svc.GetDevices()
		if err != nil {
			return nil, err
		}
		return d, nil
	}
}

// DecodeUppercaseRequest decodes the upper case request from the body
func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeResponse encodes the response into JSON.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeGetDevicesRequest decodes the request from the body.
func DecodeGetDevicesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
