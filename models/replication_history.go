// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicationHistory replication history
//
// swagger:model ReplicationHistory
type ReplicationHistory struct {

	// Number of document write failures
	DocWriteFailures int64 `json:"doc_write_failures,omitempty"`

	// Number of documents read
	DocsRead int64 `json:"docs_read,omitempty"`

	// Number of documents written to target
	DocsWritten int64 `json:"docs_written,omitempty"`

	// Last sequence number in changes stream
	EndLastSeq int64 `json:"end_last_seq,omitempty"`

	// Date/Time replication operation completed in RFC 2822 format
	EndTime string `json:"end_time,omitempty"`

	// Number of missing documents checked
	MissingChecked int64 `json:"missing_checked,omitempty"`

	// Number of missing documents found
	MissingFound int64 `json:"missing_found,omitempty"`

	// Last recorded sequence number
	RecordedSeq int64 `json:"recorded_seq,omitempty"`

	// Session ID for this replication operation
	SessionID string `json:"session_id,omitempty"`

	// First sequence number in changes stream
	StartLastSeq int64 `json:"start_last_seq,omitempty"`

	// Date/Time replication operation started in RFC 2822 format
	StartTime string `json:"start_time,omitempty"`
}

// Validate validates this replication history
func (m *ReplicationHistory) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this replication history based on context it is used
func (m *ReplicationHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationHistory) UnmarshalBinary(b []byte) error {
	var res ReplicationHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
