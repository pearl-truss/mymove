// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListMTOsHandlerFunc turns a function with the right signature into a list m t os handler
type ListMTOsHandlerFunc func(ListMTOsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListMTOsHandlerFunc) Handle(params ListMTOsParams) middleware.Responder {
	return fn(params)
}

// ListMTOsHandler interface for that can handle valid list m t os params
type ListMTOsHandler interface {
	Handle(ListMTOsParams) middleware.Responder
}

// NewListMTOs creates a new http.Handler for the list m t os operation
func NewListMTOs(ctx *middleware.Context, handler ListMTOsHandler) *ListMTOs {
	return &ListMTOs{Context: ctx, Handler: handler}
}

/*ListMTOs swagger:route GET /move-task-orders moveTaskOrder listMTOs

listMTOs

Gets all move task orders. Provides all move task orders regardless of whether or not they have been made available to prime.


*/
type ListMTOs struct {
	Context *middleware.Context
	Handler ListMTOsHandler
}

func (o *ListMTOs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListMTOsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
