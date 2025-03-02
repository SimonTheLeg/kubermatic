// Code generated by go-swagger; DO NOT EDIT.

package gke

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListGKEVMSizesReader is a Reader for the ListGKEVMSizes structure.
type ListGKEVMSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGKEVMSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGKEVMSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGKEVMSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGKEVMSizesOK creates a ListGKEVMSizesOK with default headers values
func NewListGKEVMSizesOK() *ListGKEVMSizesOK {
	return &ListGKEVMSizesOK{}
}

/*
ListGKEVMSizesOK describes a response with status code 200, with default header values.

GCPMachineSizeList
*/
type ListGKEVMSizesOK struct {
	Payload models.GCPMachineSizeList
}

// IsSuccess returns true when this list g k e Vm sizes o k response has a 2xx status code
func (o *ListGKEVMSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g k e Vm sizes o k response has a 3xx status code
func (o *ListGKEVMSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g k e Vm sizes o k response has a 4xx status code
func (o *ListGKEVMSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g k e Vm sizes o k response has a 5xx status code
func (o *ListGKEVMSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g k e Vm sizes o k response a status code equal to that given
func (o *ListGKEVMSizesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGKEVMSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/vmsizes][%d] listGKEVmSizesOK  %+v", 200, o.Payload)
}

func (o *ListGKEVMSizesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/vmsizes][%d] listGKEVmSizesOK  %+v", 200, o.Payload)
}

func (o *ListGKEVMSizesOK) GetPayload() models.GCPMachineSizeList {
	return o.Payload
}

func (o *ListGKEVMSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGKEVMSizesDefault creates a ListGKEVMSizesDefault with default headers values
func NewListGKEVMSizesDefault(code int) *ListGKEVMSizesDefault {
	return &ListGKEVMSizesDefault{
		_statusCode: code,
	}
}

/*
ListGKEVMSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGKEVMSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g k e VM sizes default response
func (o *ListGKEVMSizesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g k e VM sizes default response has a 2xx status code
func (o *ListGKEVMSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g k e VM sizes default response has a 3xx status code
func (o *ListGKEVMSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g k e VM sizes default response has a 4xx status code
func (o *ListGKEVMSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g k e VM sizes default response has a 5xx status code
func (o *ListGKEVMSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g k e VM sizes default response a status code equal to that given
func (o *ListGKEVMSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGKEVMSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/vmsizes][%d] listGKEVMSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListGKEVMSizesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/vmsizes][%d] listGKEVMSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListGKEVMSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGKEVMSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
