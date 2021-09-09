// Code generated by go-swagger; DO NOT EDIT.

package design_documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// DesignDocDeleteReader is a Reader for the DesignDocDelete structure.
type DesignDocDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignDocDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignDocDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDesignDocDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDesignDocDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDesignDocDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDesignDocDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDesignDocDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDesignDocDeleteOK creates a DesignDocDeleteOK with default headers values
func NewDesignDocDeleteOK() *DesignDocDeleteOK {
	return &DesignDocDeleteOK{}
}

/* DesignDocDeleteOK describes a response with status code 200, with default header values.

Document successfully removed
*/
type DesignDocDeleteOK struct {

	/* Double quoted document’s revision token
	 */
	ETag string

	Payload *models.DocumentOK
}

func (o *DesignDocDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designDocDeleteOK  %+v", 200, o.Payload)
}
func (o *DesignDocDeleteOK) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DesignDocDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
	}

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocDeleteAccepted creates a DesignDocDeleteAccepted with default headers values
func NewDesignDocDeleteAccepted() *DesignDocDeleteAccepted {
	return &DesignDocDeleteAccepted{}
}

/* DesignDocDeleteAccepted describes a response with status code 202, with default header values.

Request was accepted, but changes are not yet stored on disk
*/
type DesignDocDeleteAccepted struct {

	/* Double quoted document’s revision token
	 */
	ETag string

	Payload *models.DocumentOK
}

func (o *DesignDocDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designDocDeleteAccepted  %+v", 202, o.Payload)
}
func (o *DesignDocDeleteAccepted) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DesignDocDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
	}

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocDeleteBadRequest creates a DesignDocDeleteBadRequest with default headers values
func NewDesignDocDeleteBadRequest() *DesignDocDeleteBadRequest {
	return &DesignDocDeleteBadRequest{}
}

/* DesignDocDeleteBadRequest describes a response with status code 400, with default header values.

Invalid request body or parameters
*/
type DesignDocDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designDocDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *DesignDocDeleteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocDeleteUnauthorized creates a DesignDocDeleteUnauthorized with default headers values
func NewDesignDocDeleteUnauthorized() *DesignDocDeleteUnauthorized {
	return &DesignDocDeleteUnauthorized{}
}

/* DesignDocDeleteUnauthorized describes a response with status code 401, with default header values.

Write privileges required
*/
type DesignDocDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designDocDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *DesignDocDeleteUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocDeleteNotFound creates a DesignDocDeleteNotFound with default headers values
func NewDesignDocDeleteNotFound() *DesignDocDeleteNotFound {
	return &DesignDocDeleteNotFound{}
}

/* DesignDocDeleteNotFound describes a response with status code 404, with default header values.

Specified database or document ID doesn’t exists
*/
type DesignDocDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designDocDeleteNotFound  %+v", 404, o.Payload)
}
func (o *DesignDocDeleteNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocDeleteConflict creates a DesignDocDeleteConflict with default headers values
func NewDesignDocDeleteConflict() *DesignDocDeleteConflict {
	return &DesignDocDeleteConflict{}
}

/* DesignDocDeleteConflict describes a response with status code 409, with default header values.

Specified revision is not the latest for target document
*/
type DesignDocDeleteConflict struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designDocDeleteConflict  %+v", 409, o.Payload)
}
func (o *DesignDocDeleteConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
