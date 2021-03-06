// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new move task order API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for move task order API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateMoveTaskOrder creates move task order

Creates an instance of moveTaskOrder.
Current this will also create a number of nested objects but not all.
It will currently create
* MoveTaskOrder
* MoveOrder
* Customer
* User
* Entitlement

It will not create addresses or duty stations. <br />
<br />
This is a support endpoint and will not be available in production.

*/
func (a *Client) CreateMoveTaskOrder(params *CreateMoveTaskOrderParams) (*CreateMoveTaskOrderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMoveTaskOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMoveTaskOrder",
		Method:             "POST",
		PathPattern:        "/move-task-orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateMoveTaskOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMoveTaskOrderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMoveTaskOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMoveTaskOrder gets move task order

Gets an individual move task order by ID. <br />
<br />
This is a support endpoint and will not be available in production.

*/
func (a *Client) GetMoveTaskOrder(params *GetMoveTaskOrderParams) (*GetMoveTaskOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMoveTaskOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMoveTaskOrder",
		Method:             "GET",
		PathPattern:        "/move-task-orders/{moveTaskOrderID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMoveTaskOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMoveTaskOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMoveTaskOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListMTOs lists m t os

Gets all move task orders. Provides all move task orders regardless of whether or not they have been made available to prime.

*/
func (a *Client) ListMTOs(params *ListMTOsParams) (*ListMTOsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMTOsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listMTOs",
		Method:             "GET",
		PathPattern:        "/move-task-orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListMTOsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMTOsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listMTOs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
MakeMoveTaskOrderAvailable makes move task order available

Updates move task order `availableToPrimeAt` to make it available to prime. No request body required. <br />
<br />
This is a support endpoint and will not be available in production.

*/
func (a *Client) MakeMoveTaskOrderAvailable(params *MakeMoveTaskOrderAvailableParams) (*MakeMoveTaskOrderAvailableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMakeMoveTaskOrderAvailableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "makeMoveTaskOrderAvailable",
		Method:             "PATCH",
		PathPattern:        "/move-task-orders/{moveTaskOrderID}/available-to-prime",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MakeMoveTaskOrderAvailableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MakeMoveTaskOrderAvailableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for makeMoveTaskOrderAvailable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
