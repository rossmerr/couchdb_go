// Code generated by go-swagger; DO NOT EDIT.

package design_documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/RossMerr/couchdb_go/models"
)

// DesignGetDocSearchInfoReader is a Reader for the DesignGetDocSearchInfo structure.
type DesignGetDocSearchInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignGetDocSearchInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignGetDocSearchInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDesignGetDocSearchInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDesignGetDocSearchInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDesignGetDocSearchInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDesignGetDocSearchInfoOK creates a DesignGetDocSearchInfoOK with default headers values
func NewDesignGetDocSearchInfoOK() *DesignGetDocSearchInfoOK {
	return &DesignGetDocSearchInfoOK{}
}

/*DesignGetDocSearchInfoOK handles this case with default header values.

Request completed successfully
*/
type DesignGetDocSearchInfoOK struct {
	/*Response signature
	 */
	ETag string
	/*chunked
	 */
	TransferEncoding string

	Payload *DesignGetDocSearchInfoOKBody
}

func (o *DesignGetDocSearchInfoOK) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_search_info/{index}][%d] designGetDocSearchInfoOK  %+v", 200, o.Payload)
}

func (o *DesignGetDocSearchInfoOK) GetPayload() *DesignGetDocSearchInfoOKBody {
	return o.Payload
}

func (o *DesignGetDocSearchInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Transfer-Encoding
	o.TransferEncoding = response.GetHeader("Transfer-Encoding")

	o.Payload = new(DesignGetDocSearchInfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignGetDocSearchInfoBadRequest creates a DesignGetDocSearchInfoBadRequest with default headers values
func NewDesignGetDocSearchInfoBadRequest() *DesignGetDocSearchInfoBadRequest {
	return &DesignGetDocSearchInfoBadRequest{}
}

/*DesignGetDocSearchInfoBadRequest handles this case with default header values.

Invalid request
*/
type DesignGetDocSearchInfoBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DesignGetDocSearchInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_search_info/{index}][%d] designGetDocSearchInfoBadRequest  %+v", 400, o.Payload)
}

func (o *DesignGetDocSearchInfoBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignGetDocSearchInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignGetDocSearchInfoUnauthorized creates a DesignGetDocSearchInfoUnauthorized with default headers values
func NewDesignGetDocSearchInfoUnauthorized() *DesignGetDocSearchInfoUnauthorized {
	return &DesignGetDocSearchInfoUnauthorized{}
}

/*DesignGetDocSearchInfoUnauthorized handles this case with default header values.

Read permission required
*/
type DesignGetDocSearchInfoUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DesignGetDocSearchInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_search_info/{index}][%d] designGetDocSearchInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *DesignGetDocSearchInfoUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignGetDocSearchInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignGetDocSearchInfoNotFound creates a DesignGetDocSearchInfoNotFound with default headers values
func NewDesignGetDocSearchInfoNotFound() *DesignGetDocSearchInfoNotFound {
	return &DesignGetDocSearchInfoNotFound{}
}

/*DesignGetDocSearchInfoNotFound handles this case with default header values.

Specified database, design document or view is missed
*/
type DesignGetDocSearchInfoNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DesignGetDocSearchInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /{db}/_design/{ddoc}/_search_info/{index}][%d] designGetDocSearchInfoNotFound  %+v", 404, o.Payload)
}

func (o *DesignGetDocSearchInfoNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignGetDocSearchInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DesignGetDocSearchInfoOKBody design get doc search info o k body
swagger:model DesignGetDocSearchInfoOKBody
*/
type DesignGetDocSearchInfoOKBody struct {

	// name
	Name string `json:"name,omitempty"`

	// search index
	SearchIndex *models.SearchIndex `json:"search_index,omitempty"`
}

// Validate validates this design get doc search info o k body
func (o *DesignGetDocSearchInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSearchIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DesignGetDocSearchInfoOKBody) validateSearchIndex(formats strfmt.Registry) error {

	if swag.IsZero(o.SearchIndex) { // not required
		return nil
	}

	if o.SearchIndex != nil {
		if err := o.SearchIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("designGetDocSearchInfoOK" + "." + "search_index")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DesignGetDocSearchInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DesignGetDocSearchInfoOKBody) UnmarshalBinary(b []byte) error {
	var res DesignGetDocSearchInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
