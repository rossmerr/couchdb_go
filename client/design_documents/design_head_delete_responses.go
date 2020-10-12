// Code generated by go-swagger; DO NOT EDIT.

package design_documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RossMerr/couchdb_go/models"
)

// DesignHeadDeleteReader is a Reader for the DesignHeadDelete structure.
type DesignHeadDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignHeadDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignHeadDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDesignHeadDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDesignHeadDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDesignHeadDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDesignHeadDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDesignHeadDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDesignHeadDeleteOK creates a DesignHeadDeleteOK with default headers values
func NewDesignHeadDeleteOK() *DesignHeadDeleteOK {
	return &DesignHeadDeleteOK{}
}

/*DesignHeadDeleteOK handles this case with default header values.

Document successfully removed
*/
type DesignHeadDeleteOK struct {
	/*Double quoted document’s revision token
	 */
	ETag string

	Payload *models.DocumentOK
}

func (o *DesignHeadDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designHeadDeleteOK  %+v", 200, o.Payload)
}

func (o *DesignHeadDeleteOK) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DesignHeadDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignHeadDeleteAccepted creates a DesignHeadDeleteAccepted with default headers values
func NewDesignHeadDeleteAccepted() *DesignHeadDeleteAccepted {
	return &DesignHeadDeleteAccepted{}
}

/*DesignHeadDeleteAccepted handles this case with default header values.

Request was accepted, but changes are not yet stored on disk
*/
type DesignHeadDeleteAccepted struct {
	/*Double quoted document’s revision token
	 */
	ETag string

	Payload *models.DocumentOK
}

func (o *DesignHeadDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designHeadDeleteAccepted  %+v", 202, o.Payload)
}

func (o *DesignHeadDeleteAccepted) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DesignHeadDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignHeadDeleteBadRequest creates a DesignHeadDeleteBadRequest with default headers values
func NewDesignHeadDeleteBadRequest() *DesignHeadDeleteBadRequest {
	return &DesignHeadDeleteBadRequest{}
}

/*DesignHeadDeleteBadRequest handles this case with default header values.

Invalid request body or parameters
*/
type DesignHeadDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DesignHeadDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designHeadDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DesignHeadDeleteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignHeadDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignHeadDeleteUnauthorized creates a DesignHeadDeleteUnauthorized with default headers values
func NewDesignHeadDeleteUnauthorized() *DesignHeadDeleteUnauthorized {
	return &DesignHeadDeleteUnauthorized{}
}

/*DesignHeadDeleteUnauthorized handles this case with default header values.

Write privileges required
*/
type DesignHeadDeleteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DesignHeadDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designHeadDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *DesignHeadDeleteUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignHeadDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignHeadDeleteNotFound creates a DesignHeadDeleteNotFound with default headers values
func NewDesignHeadDeleteNotFound() *DesignHeadDeleteNotFound {
	return &DesignHeadDeleteNotFound{}
}

/*DesignHeadDeleteNotFound handles this case with default header values.

Specified database or document ID doesn’t exists
*/
type DesignHeadDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DesignHeadDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designHeadDeleteNotFound  %+v", 404, o.Payload)
}

func (o *DesignHeadDeleteNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignHeadDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignHeadDeleteConflict creates a DesignHeadDeleteConflict with default headers values
func NewDesignHeadDeleteConflict() *DesignHeadDeleteConflict {
	return &DesignHeadDeleteConflict{}
}

/*DesignHeadDeleteConflict handles this case with default header values.

Specified revision is not the latest for target document
*/
type DesignHeadDeleteConflict struct {
	Payload *models.ErrorResponse
}

func (o *DesignHeadDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /{db}/_design/{ddoc}][%d] designHeadDeleteConflict  %+v", 409, o.Payload)
}

func (o *DesignHeadDeleteConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignHeadDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
