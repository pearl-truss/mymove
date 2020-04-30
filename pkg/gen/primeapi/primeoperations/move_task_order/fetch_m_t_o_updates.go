// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FetchMTOUpdatesHandlerFunc turns a function with the right signature into a fetch m t o updates handler
type FetchMTOUpdatesHandlerFunc func(FetchMTOUpdatesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FetchMTOUpdatesHandlerFunc) Handle(params FetchMTOUpdatesParams) middleware.Responder {
	return fn(params)
}

// FetchMTOUpdatesHandler interface for that can handle valid fetch m t o updates params
type FetchMTOUpdatesHandler interface {
	Handle(FetchMTOUpdatesParams) middleware.Responder
}

// NewFetchMTOUpdates creates a new http.Handler for the fetch m t o updates operation
func NewFetchMTOUpdates(ctx *middleware.Context, handler FetchMTOUpdatesHandler) *FetchMTOUpdates {
	return &FetchMTOUpdates{Context: ctx, Handler: handler}
}

/*FetchMTOUpdates swagger:route GET /move-task-orders moveTaskOrder fetchMTOUpdates

Gets all move task orders where `isAvailableToPrime` is TRUE.

Gets all move task orders where `isAvailableToPrime` is TRUE. This prevents viewing any move task orders that have not been made available to the Prime.


*/
type FetchMTOUpdates struct {
	Context *middleware.Context
	Handler FetchMTOUpdatesHandler
}

func (o *FetchMTOUpdates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFetchMTOUpdatesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
