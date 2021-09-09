// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Body2 List of documents objects
//
// swagger:model body_2
type Body2 struct {

	// docs
	Docs []Document `json:"docs"`
}

// Validate validates this body 2
func (m *Body2) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this body 2 based on context it is used
func (m *Body2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Body2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body2) UnmarshalBinary(b []byte) error {
	var res Body2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
