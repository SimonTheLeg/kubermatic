// Code generated by go-swagger; DO NOT EDIT.

package hetzner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListHetznerSizesReader is a Reader for the ListHetznerSizes structure.
type ListHetznerSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHetznerSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListHetznerSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListHetznerSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListHetznerSizesOK creates a ListHetznerSizesOK with default headers values
func NewListHetznerSizesOK() *ListHetznerSizesOK {
	return &ListHetznerSizesOK{}
}

/* ListHetznerSizesOK describes a response with status code 200, with default header values.

HetznerSizeList
*/
type ListHetznerSizesOK struct {
	Payload *models.HetznerSizeList
}

// IsSuccess returns true when this list hetzner sizes o k response has a 2xx status code
func (o *ListHetznerSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list hetzner sizes o k response has a 3xx status code
func (o *ListHetznerSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list hetzner sizes o k response has a 4xx status code
func (o *ListHetznerSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list hetzner sizes o k response has a 5xx status code
func (o *ListHetznerSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list hetzner sizes o k response a status code equal to that given
func (o *ListHetznerSizesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListHetznerSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/hetzner/sizes][%d] listHetznerSizesOK  %+v", 200, o.Payload)
}

func (o *ListHetznerSizesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/hetzner/sizes][%d] listHetznerSizesOK  %+v", 200, o.Payload)
}

func (o *ListHetznerSizesOK) GetPayload() *models.HetznerSizeList {
	return o.Payload
}

func (o *ListHetznerSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HetznerSizeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHetznerSizesDefault creates a ListHetznerSizesDefault with default headers values
func NewListHetznerSizesDefault(code int) *ListHetznerSizesDefault {
	return &ListHetznerSizesDefault{
		_statusCode: code,
	}
}

/* ListHetznerSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListHetznerSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list hetzner sizes default response
func (o *ListHetznerSizesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list hetzner sizes default response has a 2xx status code
func (o *ListHetznerSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list hetzner sizes default response has a 3xx status code
func (o *ListHetznerSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list hetzner sizes default response has a 4xx status code
func (o *ListHetznerSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list hetzner sizes default response has a 5xx status code
func (o *ListHetznerSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list hetzner sizes default response a status code equal to that given
func (o *ListHetznerSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListHetznerSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/hetzner/sizes][%d] listHetznerSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListHetznerSizesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/hetzner/sizes][%d] listHetznerSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListHetznerSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListHetznerSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
