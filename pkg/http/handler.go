package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"
	endpoint "todo/pkg/endpoint"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeGetHandler creates the handler logic
//func makeGetHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
//	m.Methods("POST").Path("/get").Handler(
//	handlers.CORS(
//	handlers.AllowedMethods([]string{"POST"}),
//	handlers.AllowedOrigins([]string{"*"})
//	)(http.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)))
//}

// makeGetHandler creates the handler logic
func makeGetHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)),
	)
}

// decodeGetRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
//func decodeGetRequest(_ context.Context, r *http1.Request) (interface{}, error) {
//	req := endpoint.GetRequest{}
//	err := json.NewDecoder(r.Body).Decode(&req)
//	return req, err
//}

func decodeGetRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetRequest{}
	return req, nil
}

// encodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddHandler creates the handler logic
//func makeAddHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
//	m.Methods("POST").Path("/add").Handler(
//	handlers.CORS(
//	handlers.AllowedMethods([]string{"POST"}),
//	handlers.AllowedOrigins([]string{"*"})
//	)(http.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...)))
//}

func makeAddHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST", "OPTIONS").Path("/add").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"POST"}),
		)(http.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...)))
}

// decodeAddRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.AddRequest{}
	err := json.NewDecoder(r.Body).Decode(&req.Todo)
	return req, err
}

// encodeAddResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSetCompleteHandler creates the handler logic
//func makeSetCompleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
//	m.Methods("POST").Path("/set-complete").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.SetCompleteEndpoint, decodeSetCompleteRequest, encodeSetCompleteResponse, options...)))
//}

func makeSetCompleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT", "OPTIONS").Path("/set-complete").Handler(
		handlers.CORS(
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.SetCompleteEndpoint, decodeSetCompleteRequest, encodeSetCompleteResponse, options...)))
}

// decodeSetCompleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSetCompleteRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.SetCompleteRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSetCompleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSetCompleteResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeRemoveCompleteHandler creates the handler logic
//func makeRemoveCompleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
//	m.Methods("POST").Path("/remove-complete").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.RemoveCompleteEndpoint, decodeRemoveCompleteRequest, encodeRemoveCompleteResponse, options...)))
//}

func makeRemoveCompleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT", "OPTIONS").Path("/remove-complete").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.RemoveCompleteEndpoint, decodeRemoveCompleteRequest, encodeRemoveCompleteResponse, options...)))
}

// decodeRemoveCompleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRemoveCompleteRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RemoveCompleteRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRemoveCompleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRemoveCompleteResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteHandler creates the handler logic
//func makeDeleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
//	m.Methods("POST").Path("/delete").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...)))
//}

func makeDeleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/delete/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...)))
}

// decodeDeleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
//func decodeDeleteRequest(_ context.Context, r *http1.Request) (interface{}, error) {
//	req := endpoint.DeleteRequest{}
//	err := json.NewDecoder(r.Body).Decode(&req)
//	return req, err
//}

func decodeDeleteRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

func makeUpdateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT", "OPTIONS").Path("/update").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"PUT"}),
		)(http.NewServer(endpoints.UpdateEndpoint, decodeUpdateRequest, encodeUpdateResponse, options...)))
}

// decodeUpdateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req.Todo)
	return req, err
}

// encodeUpdateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSetStarHandler creates the handler logic
func makeSetStarHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/set-star").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.SetStarEndpoint, decodeSetStarRequest, encodeSetStarResponse, options...)))
}

// decodeSetStarRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSetStarRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.SetStarRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSetStarResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSetStarResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
