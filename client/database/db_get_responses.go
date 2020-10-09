// Code generated by go-swagger; DO NOT EDIT.

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RossMerr/couchdb-go/models"
)

// DbGetReader is a Reader for the DbGet structure.
type DbGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DbGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDbGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDbGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDbGetOK creates a DbGetOK with default headers values
func NewDbGetOK() *DbGetOK {
	return &DbGetOK{}
}

/*DbGetOK handles this case with default header values.

Request completed successfully
*/
type DbGetOK struct {
	Payload *models.Database
}

func (o *DbGetOK) Error() string {
	return fmt.Sprintf("[GET /{db}][%d] dbGetOK  %+v", 200, o.Payload)
}

func (o *DbGetOK) GetPayload() *models.Database {
	return o.Payload
}

func (o *DbGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Database)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDbGetNotFound creates a DbGetNotFound with default headers values
func NewDbGetNotFound() *DbGetNotFound {
	return &DbGetNotFound{}
}

/*DbGetNotFound handles this case with default header values.

Requested database not found
*/
type DbGetNotFound struct {
	Payload *models.Error
}

func (o *DbGetNotFound) Error() string {
	return fmt.Sprintf("[GET /{db}][%d] dbGetNotFound  %+v", 404, o.Payload)
}

func (o *DbGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DbGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
