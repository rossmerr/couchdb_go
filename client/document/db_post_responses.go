// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RossMerr/couchdb_go/models"
)

// DbPostReader is a Reader for the DbPost structure.
type DbPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DbPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDbPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDbPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDbPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDbPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDbPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDbPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDbPostCreated creates a DbPostCreated with default headers values
func NewDbPostCreated() *DbPostCreated {
	return &DbPostCreated{}
}

/*DbPostCreated handles this case with default header values.

Document created and stored on disk
*/
type DbPostCreated struct {
	Payload *models.DocumentOK
}

func (o *DbPostCreated) Error() string {
	return fmt.Sprintf("[POST /{db}][%d] dbPostCreated  %+v", 201, o.Payload)
}

func (o *DbPostCreated) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DbPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDbPostAccepted creates a DbPostAccepted with default headers values
func NewDbPostAccepted() *DbPostAccepted {
	return &DbPostAccepted{}
}

/*DbPostAccepted handles this case with default header values.

Document data accepted, but not yet stored on disk
*/
type DbPostAccepted struct {
	Payload *models.DocumentOK
}

func (o *DbPostAccepted) Error() string {
	return fmt.Sprintf("[POST /{db}][%d] dbPostAccepted  %+v", 202, o.Payload)
}

func (o *DbPostAccepted) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DbPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDbPostBadRequest creates a DbPostBadRequest with default headers values
func NewDbPostBadRequest() *DbPostBadRequest {
	return &DbPostBadRequest{}
}

/*DbPostBadRequest handles this case with default header values.

Invalid database name
*/
type DbPostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DbPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /{db}][%d] dbPostBadRequest  %+v", 400, o.Payload)
}

func (o *DbPostBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DbPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDbPostUnauthorized creates a DbPostUnauthorized with default headers values
func NewDbPostUnauthorized() *DbPostUnauthorized {
	return &DbPostUnauthorized{}
}

/*DbPostUnauthorized handles this case with default header values.

Write privileges required
*/
type DbPostUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DbPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{db}][%d] dbPostUnauthorized  %+v", 401, o.Payload)
}

func (o *DbPostUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DbPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDbPostNotFound creates a DbPostNotFound with default headers values
func NewDbPostNotFound() *DbPostNotFound {
	return &DbPostNotFound{}
}

/*DbPostNotFound handles this case with default header values.

Database doesn’t exist
*/
type DbPostNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DbPostNotFound) Error() string {
	return fmt.Sprintf("[POST /{db}][%d] dbPostNotFound  %+v", 404, o.Payload)
}

func (o *DbPostNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DbPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDbPostConflict creates a DbPostConflict with default headers values
func NewDbPostConflict() *DbPostConflict {
	return &DbPostConflict{}
}

/*DbPostConflict handles this case with default header values.

A Conflicting Document with same ID already exists
*/
type DbPostConflict struct {
	Payload *models.ErrorResponse
}

func (o *DbPostConflict) Error() string {
	return fmt.Sprintf("[POST /{db}][%d] dbPostConflict  %+v", 409, o.Payload)
}

func (o *DbPostConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DbPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}