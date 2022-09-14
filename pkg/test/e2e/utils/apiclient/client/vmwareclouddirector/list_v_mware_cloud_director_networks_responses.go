// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListVMwareCloudDirectorNetworksReader is a Reader for the ListVMwareCloudDirectorNetworks structure.
type ListVMwareCloudDirectorNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVMwareCloudDirectorNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVMwareCloudDirectorNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVMwareCloudDirectorNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVMwareCloudDirectorNetworksOK creates a ListVMwareCloudDirectorNetworksOK with default headers values
func NewListVMwareCloudDirectorNetworksOK() *ListVMwareCloudDirectorNetworksOK {
	return &ListVMwareCloudDirectorNetworksOK{}
}

/* ListVMwareCloudDirectorNetworksOK describes a response with status code 200, with default header values.

VMwareCloudDirectorNetworkList
*/
type ListVMwareCloudDirectorNetworksOK struct {
	Payload models.VMwareCloudDirectorNetworkList
}

// IsSuccess returns true when this list v mware cloud director networks o k response has a 2xx status code
func (o *ListVMwareCloudDirectorNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list v mware cloud director networks o k response has a 3xx status code
func (o *ListVMwareCloudDirectorNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list v mware cloud director networks o k response has a 4xx status code
func (o *ListVMwareCloudDirectorNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list v mware cloud director networks o k response has a 5xx status code
func (o *ListVMwareCloudDirectorNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list v mware cloud director networks o k response a status code equal to that given
func (o *ListVMwareCloudDirectorNetworksOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVMwareCloudDirectorNetworksOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/networks][%d] listVMwareCloudDirectorNetworksOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/networks][%d] listVMwareCloudDirectorNetworksOK  %+v", 200, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksOK) GetPayload() models.VMwareCloudDirectorNetworkList {
	return o.Payload
}

func (o *ListVMwareCloudDirectorNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVMwareCloudDirectorNetworksDefault creates a ListVMwareCloudDirectorNetworksDefault with default headers values
func NewListVMwareCloudDirectorNetworksDefault(code int) *ListVMwareCloudDirectorNetworksDefault {
	return &ListVMwareCloudDirectorNetworksDefault{
		_statusCode: code,
	}
}

/* ListVMwareCloudDirectorNetworksDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVMwareCloudDirectorNetworksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v mware cloud director networks default response
func (o *ListVMwareCloudDirectorNetworksDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list v mware cloud director networks default response has a 2xx status code
func (o *ListVMwareCloudDirectorNetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list v mware cloud director networks default response has a 3xx status code
func (o *ListVMwareCloudDirectorNetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list v mware cloud director networks default response has a 4xx status code
func (o *ListVMwareCloudDirectorNetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list v mware cloud director networks default response has a 5xx status code
func (o *ListVMwareCloudDirectorNetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list v mware cloud director networks default response a status code equal to that given
func (o *ListVMwareCloudDirectorNetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListVMwareCloudDirectorNetworksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/networks][%d] listVMwareCloudDirectorNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/vmwareclouddirector/{dc}/networks][%d] listVMwareCloudDirectorNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *ListVMwareCloudDirectorNetworksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVMwareCloudDirectorNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
