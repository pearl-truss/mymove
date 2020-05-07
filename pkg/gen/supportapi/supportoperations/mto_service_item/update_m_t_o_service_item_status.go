// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateMTOServiceItemStatusHandlerFunc turns a function with the right signature into a update m t o service item status handler
type UpdateMTOServiceItemStatusHandlerFunc func(UpdateMTOServiceItemStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateMTOServiceItemStatusHandlerFunc) Handle(params UpdateMTOServiceItemStatusParams) middleware.Responder {
	return fn(params)
}

// UpdateMTOServiceItemStatusHandler interface for that can handle valid update m t o service item status params
type UpdateMTOServiceItemStatusHandler interface {
	Handle(UpdateMTOServiceItemStatusParams) middleware.Responder
}

// NewUpdateMTOServiceItemStatus creates a new http.Handler for the update m t o service item status operation
func NewUpdateMTOServiceItemStatus(ctx *middleware.Context, handler UpdateMTOServiceItemStatusHandler) *UpdateMTOServiceItemStatus {
	return &UpdateMTOServiceItemStatus{Context: ctx, Handler: handler}
}

/*UpdateMTOServiceItemStatus swagger:route PATCH /service-items/{mtoServiceItemID}/status mtoServiceItem updateMTOServiceItemStatus

updateMTOServiceItemStatus

Updates the status of a service item for a move order to APPROVED or REJECTED. <br />
<br />
This is a support endpoint and will not be available in production.


*/
type UpdateMTOServiceItemStatus struct {
	Context *middleware.Context
	Handler UpdateMTOServiceItemStatusHandler
}

func (o *UpdateMTOServiceItemStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateMTOServiceItemStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
