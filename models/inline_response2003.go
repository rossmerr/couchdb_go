// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InlineResponse2003 inline response 200 3
//
// swagger:model inline_response_200_3
type InlineResponse2003 struct {

	// tokens
	Tokens []string `json:"tokens"`
}

// Validate validates this inline response 200 3
func (m *InlineResponse2003) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inline response 200 3 based on context it is used
func (m *InlineResponse2003) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2003) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2003) UnmarshalBinary(b []byte) error {
	var res InlineResponse2003
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
