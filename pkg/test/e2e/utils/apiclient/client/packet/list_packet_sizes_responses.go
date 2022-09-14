// Code generated by go-swagger; DO NOT EDIT.

package packet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListPacketSizesReader is a Reader for the ListPacketSizes structure.
type ListPacketSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPacketSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPacketSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPacketSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPacketSizesOK creates a ListPacketSizesOK with default headers values
func NewListPacketSizesOK() *ListPacketSizesOK {
	return &ListPacketSizesOK{}
}

/* ListPacketSizesOK describes a response with status code 200, with default header values.

PacketSizeList
*/
type ListPacketSizesOK struct {
	Payload []models.PacketSizeList
}

// IsSuccess returns true when this list packet sizes o k response has a 2xx status code
func (o *ListPacketSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list packet sizes o k response has a 3xx status code
func (o *ListPacketSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list packet sizes o k response has a 4xx status code
func (o *ListPacketSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list packet sizes o k response has a 5xx status code
func (o *ListPacketSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list packet sizes o k response a status code equal to that given
func (o *ListPacketSizesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListPacketSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/packet/sizes][%d] listPacketSizesOK  %+v", 200, o.Payload)
}

func (o *ListPacketSizesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/packet/sizes][%d] listPacketSizesOK  %+v", 200, o.Payload)
}

func (o *ListPacketSizesOK) GetPayload() []models.PacketSizeList {
	return o.Payload
}

func (o *ListPacketSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPacketSizesDefault creates a ListPacketSizesDefault with default headers values
func NewListPacketSizesDefault(code int) *ListPacketSizesDefault {
	return &ListPacketSizesDefault{
		_statusCode: code,
	}
}

/* ListPacketSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListPacketSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list packet sizes default response
func (o *ListPacketSizesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list packet sizes default response has a 2xx status code
func (o *ListPacketSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list packet sizes default response has a 3xx status code
func (o *ListPacketSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list packet sizes default response has a 4xx status code
func (o *ListPacketSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list packet sizes default response has a 5xx status code
func (o *ListPacketSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list packet sizes default response a status code equal to that given
func (o *ListPacketSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListPacketSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/packet/sizes][%d] listPacketSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListPacketSizesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/packet/sizes][%d] listPacketSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListPacketSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListPacketSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
