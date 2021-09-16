// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicateCreateTargetParams replicate create target params
//
// swagger:model Replicate_create_target_params
type ReplicateCreateTargetParams struct {

	// partitioned
	Partitioned bool `json:"partitioned,omitempty"`
}

// Validate validates this replicate create target params
func (m *ReplicateCreateTargetParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this replicate create target params based on context it is used
func (m *ReplicateCreateTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicateCreateTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicateCreateTargetParams) UnmarshalBinary(b []byte) error {
	var res ReplicateCreateTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}