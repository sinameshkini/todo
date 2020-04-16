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

// AddCategoryRequest collects the request parameters for the AddCategory method.
type AddCategoryRequest struct {
	Category io.TodoCategory `json:"category"`
}

// AddCategoryResponse collects the response parameters for the AddCategory method.
type AddCategoryResponse struct {
	C     io.TodoCategory `json:"c"`
	Error error           `json:"error"`
}

// MakeAddCategoryEndpoint returns an endpoint that invokes AddCategory on the service.
func MakeAddCategoryEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddCategoryRequest)
		c, error := s.AddCategory(ctx, req.Category)
		return AddCategoryResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddCategoryResponse) Failed() error {
	return r.Error
}

// AddCategory implements Service. Primarily useful in a client.
func (e Endpoints) AddCategory(ctx context.Context, category io.TodoCategory) (c io.TodoCategory, error error) {
	request := AddCategoryRequest{Category: category}
	response, err := e.AddCategoryEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddCategoryResponse).C, response.(AddCategoryResponse).Error
}

// GetCategoryRequest collects the request parameters for the GetCategory method.
type GetCategoryRequest struct{}

// GetCategoryResponse collects the response parameters for the GetCategory method.
type GetCategoryResponse struct {
	C     []io.TodoCategory `json:"c"`
	Error error             `json:"error"`
}

// MakeGetCategoryEndpoint returns an endpoint that invokes GetCategory on the service.
func MakeGetCategoryEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		c, error := s.GetCategory(ctx)
		return GetCategoryResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetCategoryResponse) Failed() error {
	return r.Error
}

// UpdateCategoryRequest collects the request parameters for the UpdateCategory method.
type UpdateCategoryRequest struct {
	Category io.TodoCategory `json:"category"`
}

// UpdateCategoryResponse collects the response parameters for the UpdateCategory method.
type UpdateCategoryResponse struct {
	C     io.TodoCategory `json:"c"`
	Error error           `json:"error"`
}

// MakeUpdateCategoryEndpoint returns an endpoint that invokes UpdateCategory on the service.
func MakeUpdateCategoryEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCategoryRequest)
		c, error := s.UpdateCategory(ctx, req.Category)
		return UpdateCategoryResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateCategoryResponse) Failed() error {
	return r.Error
}

// DeleteCategoryRequest collects the request parameters for the DeleteCategory method.
type DeleteCategoryRequest struct {
	Id string `json:"id"`
}

// DeleteCategoryResponse collects the response parameters for the DeleteCategory method.
type DeleteCategoryResponse struct {
	Error error `json:"error"`
}

// MakeDeleteCategoryEndpoint returns an endpoint that invokes DeleteCategory on the service.
func MakeDeleteCategoryEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCategoryRequest)
		error := s.DeleteCategory(ctx, req.Id)
		return DeleteCategoryResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteCategoryResponse) Failed() error {
	return r.Error
}

// GetCategory implements Service. Primarily useful in a client.
func (e Endpoints) GetCategory(ctx context.Context) (c []io.TodoCategory, error error) {
	request := GetCategoryRequest{}
	response, err := e.GetCategoryEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetCategoryResponse).C, response.(GetCategoryResponse).Error
}

// UpdateCategory implements Service. Primarily useful in a client.
func (e Endpoints) UpdateCategory(ctx context.Context, category io.TodoCategory) (c io.TodoCategory, error error) {
	request := UpdateCategoryRequest{Category: category}
	response, err := e.UpdateCategoryEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateCategoryResponse).C, response.(UpdateCategoryResponse).Error
}

// DeleteCategory implements Service. Primarily useful in a client.
func (e Endpoints) DeleteCategory(ctx context.Context, id string) (error error) {
	request := DeleteCategoryRequest{Id: id}
	response, err := e.DeleteCategoryEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteCategoryResponse).Error
}

// GetCatChildesRequest collects the request parameters for the GetCatChildes method.
type GetCatChildesRequest struct {
	Id string `json:"id"`
}

// GetCatChildesResponse collects the response parameters for the GetCatChildes method.
type GetCatChildesResponse struct {
	C     []io.TodoCategory `json:"c"`
	Error error             `json:"error"`
}

// MakeGetCatChildesEndpoint returns an endpoint that invokes GetCatChildes on the service.
func MakeGetCatChildesEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCatChildesRequest)
		c, error := s.GetCatChildes(ctx, req.Id)
		return GetCatChildesResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetCatChildesResponse) Failed() error {
	return r.Error
}

// GetCatChildes implements Service. Primarily useful in a client.
func (e Endpoints) GetCatChildes(ctx context.Context, id string) (c []io.TodoCategory, error error) {
	request := GetCatChildesRequest{Id: id}
	response, err := e.GetCatChildesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetCatChildesResponse).C, response.(GetCatChildesResponse).Error
}
