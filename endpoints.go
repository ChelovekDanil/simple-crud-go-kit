package crud

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetEndpoint: makeGetEndpoint(s),
	}
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		user, err := s.Get(ctx, req.Id)
		return getResponse{*user, err}, nil
	}
}

type getRequest struct {
	Id string
}

type getResponse struct {
	User User  `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}
