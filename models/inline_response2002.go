// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InlineResponse2002 inline response 200 2
//
// swagger:model inline_response_200_2
type InlineResponse2002 struct {

	// all nodes
	AllNodes []string `json:"all_nodes"`

	// cluster nodes
	ClusterNodes []string `json:"cluster_nodes"`
}

// Validate validates this inline response 200 2
func (m *InlineResponse2002) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inline response 200 2 based on context it is used
func (m *InlineResponse2002) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2002) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2002) UnmarshalBinary(b []byte) error {
	var res InlineResponse2002
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}