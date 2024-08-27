package crud

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/get/{id}").Handler(httptransport.NewServer(
		e.GetEndpoint,
		decodeGetRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/get/").Handler(httptransport.NewServer(
		e.GetAllEndpoint,
		decodeGetAllRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/create/").Handler(httptransport.NewServer(
		e.CreateEndpoint,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/update/").Handler(httptransport.NewServer(
		e.UpdateEndpoint,
		decodeUpdateRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/delete/{id}").Handler(httptransport.NewServer(
		e.DeleteEndpoint,
		decodeDeleteRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return requestWithId{Id: id}, nil
}

func decodeGetAllRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return nil, nil
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req createRequest
	if err = json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req updateRequest
	if err = json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return requestWithId{Id: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=uft-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encode Error with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=uft-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
