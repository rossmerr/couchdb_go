// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// DBsInfoReader is a Reader for the DBsInfo structure.
type DBsInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DBsInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDBsInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDBsInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDBsInfoOK creates a DBsInfoOK with default headers values
func NewDBsInfoOK() *DBsInfoOK {
	return &DBsInfoOK{}
}

/* DBsInfoOK describes a response with status code 200, with default header values.

Request completed successfully
*/
type DBsInfoOK struct {
	Payload []*models.InlineResponse200
}

func (o *DBsInfoOK) Error() string {
	return fmt.Sprintf("[POST /_dbs_info][%d] dBsInfoOK  %+v", 200, o.Payload)
}
func (o *DBsInfoOK) GetPayload() []*models.InlineResponse200 {
	return o.Payload
}

func (o *DBsInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDBsInfoBadRequest creates a DBsInfoBadRequest with default headers values
func NewDBsInfoBadRequest() *DBsInfoBadRequest {
	return &DBsInfoBadRequest{}
}

/* DBsInfoBadRequest describes a response with status code 400, with default header values.

Missing keys or exceeded keys in request
*/
type DBsInfoBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DBsInfoBadRequest) Error() string {
	return fmt.Sprintf("[POST /_dbs_info][%d] dBsInfoBadRequest  %+v", 400, o.Payload)
}
func (o *DBsInfoBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DBsInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
