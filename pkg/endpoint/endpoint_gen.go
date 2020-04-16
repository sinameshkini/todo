// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "todo/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetEndpoint            endpoint.Endpoint
	AddEndpoint            endpoint.Endpoint
	SetCompleteEndpoint    endpoint.Endpoint
	RemoveCompleteEndpoint endpoint.Endpoint
	DeleteEndpoint         endpoint.Endpoint
	UpdateEndpoint         endpoint.Endpoint
	SetStarEndpoint        endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.TodoService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddEndpoint:            MakeAddEndpoint(s),
		DeleteEndpoint:         MakeDeleteEndpoint(s),
		GetEndpoint:            MakeGetEndpoint(s),
		RemoveCompleteEndpoint: MakeRemoveCompleteEndpoint(s),
		SetCompleteEndpoint:    MakeSetCompleteEndpoint(s),
		SetStarEndpoint:        MakeSetStarEndpoint(s),
		UpdateEndpoint:         MakeUpdateEndpoint(s),
	}
	for _, m := range mdw["Get"] {
		eps.GetEndpoint = m(eps.GetEndpoint)
	}
	for _, m := range mdw["Add"] {
		eps.AddEndpoint = m(eps.AddEndpoint)
	}
	for _, m := range mdw["SetComplete"] {
		eps.SetCompleteEndpoint = m(eps.SetCompleteEndpoint)
	}
	for _, m := range mdw["RemoveComplete"] {
		eps.RemoveCompleteEndpoint = m(eps.RemoveCompleteEndpoint)
	}
	for _, m := range mdw["Delete"] {
		eps.DeleteEndpoint = m(eps.DeleteEndpoint)
	}
	for _, m := range mdw["Update"] {
		eps.UpdateEndpoint = m(eps.UpdateEndpoint)
	}
	for _, m := range mdw["SetStar"] {
		eps.SetStarEndpoint = m(eps.SetStarEndpoint)
	}
	return eps
}
