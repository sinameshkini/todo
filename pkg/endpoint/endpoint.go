package endpoint

import (
	"context"
	io "todo/pkg/io"
	service "todo/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	T     []io.Todo `json:"t"`
	Error error     `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t, error := s.Get(ctx)
		return GetResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Todo io.Todo `json:"todo"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	T     io.Todo `json:"t"`
	Error error   `json:"error"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		t, error := s.Add(ctx, req.Todo)

		return AddResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Error
}

// SetCompleteRequest collects the request parameters for the SetComplete method.
type SetCompleteRequest struct {
	Id string `json:"id"`
}

// SetCompleteResponse collects the response parameters for the SetComplete method.
type SetCompleteResponse struct {
	Error error `json:"error"`
}

// MakeSetCompleteEndpoint returns an endpoint that invokes SetComplete on the service.
func MakeSetCompleteEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetCompleteRequest)
		error := s.SetComplete(ctx, req.Id)
		return SetCompleteResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r SetCompleteResponse) Failed() error {
	return r.Error
}

// RemoveCompleteRequest collects the request parameters for the RemoveComplete method.
type RemoveCompleteRequest struct {
	Id string `json:"id"`
}

// RemoveCompleteResponse collects the response parameters for the RemoveComplete method.
type RemoveCompleteResponse struct {
	Error error `json:"error"`
}

// MakeRemoveCompleteEndpoint returns an endpoint that invokes RemoveComplete on the service.
func MakeRemoveCompleteEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveCompleteRequest)
		error := s.RemoveComplete(ctx, req.Id)
		return RemoveCompleteResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r RemoveCompleteResponse) Failed() error {
	return r.Error
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Error error `json:"error"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		error := s.Delete(ctx, req.Id)
		return DeleteResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (t []io.Todo, error error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).T, response.(GetResponse).Error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	request := AddRequest{Todo: todo}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).T, response.(AddResponse).Error
}

// SetComplete implements Service. Primarily useful in a client.
func (e Endpoints) SetComplete(ctx context.Context, id string) (error error) {
	request := SetCompleteRequest{Id: id}
	response, err := e.SetCompleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SetCompleteResponse).Error
}

// RemoveComplete implements Service. Primarily useful in a client.
func (e Endpoints) RemoveComplete(ctx context.Context, id string) (error error) {
	request := RemoveCompleteRequest{Id: id}
	response, err := e.RemoveCompleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RemoveCompleteResponse).Error
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (error error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Error
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	Todo io.Todo `json:"todo"`
}

// UpdateResponse collects the response parameters for the Update method.
type UpdateResponse struct {
	T     io.Todo `json:"t"`
	Error error   `json:"error"`
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the service.
func MakeUpdateEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		t, error := s.Update(ctx, req.Todo)
		return UpdateResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateResponse) Failed() error {
	return r.Error
}

// Update implements Service. Primarily useful in a client.
func (e Endpoints) Update(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	request := UpdateRequest{Todo: todo}
	response, err := e.UpdateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateResponse).T, response.(UpdateResponse).Error
}

// Failed implements Failer.
func (r SetStarResponse) Failed() error {
	return r.Error
}

// SetStar implements Service. Primarily useful in a client.
func (e Endpoints) SetStar(ctx context.Context, star uint8) (error error) {
	request := SetStarRequest{Star: star}
	response, err := e.SetStarEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SetStarResponse).Error
}

// SetStarRequest collects the request parameters for the SetStar method.
type SetStarRequest struct {
	Id   string `json:"id"`
	Star uint8  `json:"star"`
}

// SetStarResponse collects the response parameters for the SetStar method.
type SetStarResponse struct {
	Error error `json:"error"`
}

// MakeSetStarEndpoint returns an endpoint that invokes SetStar on the service.
func MakeSetStarEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetStarRequest)
		error := s.SetStar(ctx, req.Id, req.Star)
		return SetStarResponse{Error: error}, nil
	}
}

// ReplyToRequest collects the request parameters for the ReplyTo method.
type ReplyToRequest struct {
	ParentId uint    `json:"parent_id"`
	Todo     io.Todo `json:"todo"`
}

// ReplyToResponse collects the response parameters for the ReplyTo method.
type ReplyToResponse struct {
	T     io.Todo `json:"t"`
	Error error   `json:"error"`
}

// MakeReplyToEndpoint returns an endpoint that invokes ReplyTo on the service.
func MakeReplyToEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReplyToRequest)
		t, error := s.ReplyTo(ctx, req.ParentId, req.Todo)
		return ReplyToResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r ReplyToResponse) Failed() error {
	return r.Error
}

// ReplyTo implements Service. Primarily useful in a client.
func (e Endpoints) ReplyTo(ctx context.Context, parentId uint, todo io.Todo) (t io.Todo, error error) {
	request := ReplyToRequest{
		ParentId: parentId,
		Todo:     todo,
	}
	response, err := e.ReplyToEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReplyToResponse).T, response.(ReplyToResponse).Error
}

// GetChildesRequest collects the request parameters for the GetChildes method.
type GetChildesRequest struct {
	Id string `json:"id"`
}

// GetChildesResponse collects the response parameters for the GetChildes method.
type GetChildesResponse struct {
	T     []io.Todo `json:"t"`
	Error error     `json:"error"`
}

// MakeGetChildesEndpoint returns an endpoint that invokes GetChildes on the service.
func MakeGetChildesEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetChildesRequest)
		t, error := s.GetChildes(ctx, req.Id)
		return GetChildesResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetChildesResponse) Failed() error {
	return r.Error
}

// GetChildes implements Service. Primarily useful in a client.
func (e Endpoints) GetChildes(ctx context.Context, id string) (t []io.Todo, error error) {
	request := GetChildesRequest{Id: id}
	response, err := e.GetChildesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetChildesResponse).T, response.(GetChildesResponse).Error
}
