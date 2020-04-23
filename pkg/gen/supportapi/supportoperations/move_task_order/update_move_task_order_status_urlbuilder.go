// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// UpdateMoveTaskOrderStatusURL generates an URL for the update move task order status operation
type UpdateMoveTaskOrderStatusURL struct {
	MoveTaskOrderID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateMoveTaskOrderStatusURL) WithBasePath(bp string) *UpdateMoveTaskOrderStatusURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdateMoveTaskOrderStatusURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *UpdateMoveTaskOrderStatusURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/move-task-orders/{moveTaskOrderID}/status"

	moveTaskOrderID := o.MoveTaskOrderID
	if moveTaskOrderID != "" {
		_path = strings.Replace(_path, "{moveTaskOrderID}", moveTaskOrderID, -1)
	} else {
		return nil, errors.New("moveTaskOrderId is required on UpdateMoveTaskOrderStatusURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/support/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UpdateMoveTaskOrderStatusURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UpdateMoveTaskOrderStatusURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UpdateMoveTaskOrderStatusURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UpdateMoveTaskOrderStatusURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UpdateMoveTaskOrderStatusURL")
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
func (o *UpdateMoveTaskOrderStatusURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}