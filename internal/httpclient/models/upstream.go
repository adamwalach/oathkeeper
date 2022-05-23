// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Upstream Upstream Upstream Upstream Upstream Upstream Upstream Upstream Upstream Upstream Upstream Upstream Upstream upstream
//
// swagger:model Upstream
type Upstream struct {

	// PreserveHost, if false (the default), tells ORY Oathkeeper to set the upstream request's Host header to the
	// hostname of the API's upstream's URL. Setting this flag to true instructs ORY Oathkeeper not to do so.
	PreserveHost bool `json:"preserve_host,omitempty"`

	// StripPath if set, replaces the provided path prefix when forwarding the requested URL to the upstream URL.
	StripPath string `json:"strip_path,omitempty"`

	// URL is the URL the request will be proxied to.
	URL string `json:"url,omitempty"`
}

// Validate validates this upstream
func (m *Upstream) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Upstream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Upstream) UnmarshalBinary(b []byte) error {
	var res Upstream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
