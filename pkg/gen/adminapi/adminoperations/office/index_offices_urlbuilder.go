// Code generated by go-swagger; DO NOT EDIT.

package office

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// IndexOfficesURL generates an URL for the index offices operation
type IndexOfficesURL struct {
	Filter  *string
	Page    *int64
	PerPage *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *IndexOfficesURL) WithBasePath(bp string) *IndexOfficesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *IndexOfficesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *IndexOfficesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/offices"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/admin/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var filterQ string
	if o.Filter != nil {
		filterQ = *o.Filter
	}
	if filterQ != "" {
		qs.Set("filter", filterQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt64(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var perPageQ string
	if o.PerPage != nil {
		perPageQ = swag.FormatInt64(*o.PerPage)
	}
	if perPageQ != "" {
		qs.Set("perPage", perPageQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *IndexOfficesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *IndexOfficesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *IndexOfficesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on IndexOfficesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on IndexOfficesURL")
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
func (o *IndexOfficesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
