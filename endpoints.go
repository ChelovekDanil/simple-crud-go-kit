package crud

import (
	"context"

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
		if err != nil {
			return nil, err
		}
		return getResponse{*user}, err
	}
}

func makeGetAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return getAllResponse{users}, nil
	}
}

func makeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		id, err := s.Create(ctx, User{FirstName: req.User.FirstName, LastName: req.User.LastName})
		if err != nil {
			return nil, err
		}
		return createResponse{User{Id: id, FirstName: req.User.FirstName, LastName: req.User.LastName}}, nil
	}
}

func makeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.Update(ctx, req.User)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func makeDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requestWithId)
		err := s.Delete(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

type requestWithId struct {
	Id string
}

type getResponse struct {
	User User `json:"user,omitempty"`
}

type getAllResponse struct {
	Users []User `json:"users,omitempty"`
}

type createRequest struct {
	User User
}

type createResponse struct {
	User User `json:"user,omitempty"`
}

type updateRequest struct {
	User User
}
