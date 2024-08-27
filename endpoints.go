package crud

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetEndpoint    endpoint.Endpoint
	GetAllEndpoint endpoint.Endpoint
	CreateEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetEndpoint:    makeGetEndpoint(s),
		GetAllEndpoint: makeGetAllEndpoint(s),
		CreateEndpoint: makeCreateEndpoint(s),
		UpdateEndpoint: makeUpdateEndpoint(s),
		DeleteEndpoint: makeDeleteEndpoint(s),
	}
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestWithId)
		user, err := s.Get(ctx, req.Id)
		return getResponse{*user, err}, nil
	}
}

func makeGetAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		return getAllResponse{users, err}, nil
	}
}

func makeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		id, err := s.Create(ctx, User{FirstName: req.User.FirstName, LastName: req.User.LastName})
		return createResponse{User{Id: strconv.Itoa(int(id)), FirstName: req.User.FirstName, LastName: req.User.LastName}, err}, nil
	}
}

func makeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.Update(ctx, req.User)
		return updateResponse{Err: err}, nil
	}
}

func makeDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestWithId)
		err := s.Delete(ctx, req.Id)
		return deleteResponse{Err: err}, nil
	}
}

type requestWithId struct {
	Id string
}

type getResponse struct {
	User User  `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type getAllResponse struct {
	Users []User `json:"users,omitempty"`
	Err   error  `json:"error,omitempty"`
}

type createRequest struct {
	User User
}

type createResponse struct {
	User User  `json:"user,omitempty"`
	Err  error `json:"error,omitempty"`
}

type updateRequest struct {
	User User
}

type updateResponse struct {
	Err error `json:"error,omitempty"`
}

type deleteResponse struct {
	Err error `json:"error,omitempty"`
}
