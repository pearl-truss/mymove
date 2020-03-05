// Code generated by go-swagger; DO NOT EDIT.

package mto_agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
)

// FetchMTOAgentListURL generates an URL for the fetch m t o agent list operation
type FetchMTOAgentListURL struct {
	MoveTaskOrderID strfmt.UUID
	ShipmentID      strfmt.UUID

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FetchMTOAgentListURL) WithBasePath(bp string) *FetchMTOAgentListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FetchMTOAgentListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FetchMTOAgentListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/move_task_orders/{moveTaskOrderID}/mto_shipments/{shipmentID}/mto-agents"

	moveTaskOrderID := o.MoveTaskOrderID.String()
	if moveTaskOrderID != "" {
		_path = strings.Replace(_path, "{moveTaskOrderID}", moveTaskOrderID, -1)
	} else {
		return nil, errors.New("moveTaskOrderId is required on FetchMTOAgentListURL")
	}

	shipmentID := o.ShipmentID.String()
	if shipmentID != "" {
		_path = strings.Replace(_path, "{shipmentID}", shipmentID, -1)
	} else {
		return nil, errors.New("shipmentId is required on FetchMTOAgentListURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/ghc/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FetchMTOAgentListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FetchMTOAgentListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FetchMTOAgentListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FetchMTOAgentListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FetchMTOAgentListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FetchMTOAgentListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}