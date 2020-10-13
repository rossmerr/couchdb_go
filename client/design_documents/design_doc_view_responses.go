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

// DesignDocViewReader is a Reader for the DesignDocView structure.
type DesignDocViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignDocViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignDocViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewDesignDocViewNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewDesignDocViewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDesignDocViewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDesignDocViewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDesignDocViewOK creates a DesignDocViewOK with default headers values
func NewDesignDocViewOK() *DesignDocViewOK {
	return &DesignDocViewOK{}
}

/*DesignDocViewOK handles this case with default header values.

Request completed successfully
*/
type DesignDocViewOK struct {
	/*Response signature
	 */
	ETag string
	/*chunked
	 */
	TransferEncoding string

	Payload *models.Pagination
}

func (o *DesignDocViewOK) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_view/{view}][%d] designDocViewOK  %+v", 200, o.Payload)
}

func (o *DesignDocViewOK) GetPayload() *models.Pagination {
	return o.Payload
}

func (o *DesignDocViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Transfer-Encoding
	o.TransferEncoding = response.GetHeader("Transfer-Encoding")

	o.Payload = new(models.Pagination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocViewNotModified creates a DesignDocViewNotModified with default headers values
func NewDesignDocViewNotModified() *DesignDocViewNotModified {
	return &DesignDocViewNotModified{}
}

/*DesignDocViewNotModified handles this case with default header values.

Document wasn’t modified since specified revision
*/
type DesignDocViewNotModified struct {
}

func (o *DesignDocViewNotModified) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_view/{view}][%d] designDocViewNotModified ", 304)
}

func (o *DesignDocViewNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDesignDocViewBadRequest creates a DesignDocViewBadRequest with default headers values
func NewDesignDocViewBadRequest() *DesignDocViewBadRequest {
	return &DesignDocViewBadRequest{}
}

/*DesignDocViewBadRequest handles this case with default header values.

Invalid request
*/
type DesignDocViewBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocViewBadRequest) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_view/{view}][%d] designDocViewBadRequest  %+v", 400, o.Payload)
}

func (o *DesignDocViewBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocViewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocViewUnauthorized creates a DesignDocViewUnauthorized with default headers values
func NewDesignDocViewUnauthorized() *DesignDocViewUnauthorized {
	return &DesignDocViewUnauthorized{}
}

/*DesignDocViewUnauthorized handles this case with default header values.

Read permission required
*/
type DesignDocViewUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocViewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_view/{view}][%d] designDocViewUnauthorized  %+v", 401, o.Payload)
}

func (o *DesignDocViewUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocViewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocViewNotFound creates a DesignDocViewNotFound with default headers values
func NewDesignDocViewNotFound() *DesignDocViewNotFound {
	return &DesignDocViewNotFound{}
}

/*DesignDocViewNotFound handles this case with default header values.

Specified database, design document or view is missed
*/
type DesignDocViewNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocViewNotFound) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_view/{view}][%d] designDocViewNotFound  %+v", 404, o.Payload)
}

func (o *DesignDocViewNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocViewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
