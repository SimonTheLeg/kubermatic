// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListKubevirtStorageClassesReader is a Reader for the ListKubevirtStorageClasses structure.
type ListKubevirtStorageClassesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubevirtStorageClassesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubevirtStorageClassesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubevirtStorageClassesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubevirtStorageClassesOK creates a ListKubevirtStorageClassesOK with default headers values
func NewListKubevirtStorageClassesOK() *ListKubevirtStorageClassesOK {
	return &ListKubevirtStorageClassesOK{}
}

/* ListKubevirtStorageClassesOK describes a response with status code 200, with default header values.

StorageClassList
*/
type ListKubevirtStorageClassesOK struct {
	Payload models.StorageClassList
}

// IsSuccess returns true when this list kubevirt storage classes o k response has a 2xx status code
func (o *ListKubevirtStorageClassesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list kubevirt storage classes o k response has a 3xx status code
func (o *ListKubevirtStorageClassesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list kubevirt storage classes o k response has a 4xx status code
func (o *ListKubevirtStorageClassesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list kubevirt storage classes o k response has a 5xx status code
func (o *ListKubevirtStorageClassesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list kubevirt storage classes o k response a status code equal to that given
func (o *ListKubevirtStorageClassesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListKubevirtStorageClassesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubevirtStorageClassesOK  %+v", 200, o.Payload)
}

func (o *ListKubevirtStorageClassesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubevirtStorageClassesOK  %+v", 200, o.Payload)
}

func (o *ListKubevirtStorageClassesOK) GetPayload() models.StorageClassList {
	return o.Payload
}

func (o *ListKubevirtStorageClassesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubevirtStorageClassesDefault creates a ListKubevirtStorageClassesDefault with default headers values
func NewListKubevirtStorageClassesDefault(code int) *ListKubevirtStorageClassesDefault {
	return &ListKubevirtStorageClassesDefault{
		_statusCode: code,
	}
}

/* ListKubevirtStorageClassesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListKubevirtStorageClassesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list kubevirt storage classes default response
func (o *ListKubevirtStorageClassesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list kubevirt storage classes default response has a 2xx status code
func (o *ListKubevirtStorageClassesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list kubevirt storage classes default response has a 3xx status code
func (o *ListKubevirtStorageClassesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list kubevirt storage classes default response has a 4xx status code
func (o *ListKubevirtStorageClassesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list kubevirt storage classes default response has a 5xx status code
func (o *ListKubevirtStorageClassesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list kubevirt storage classes default response a status code equal to that given
func (o *ListKubevirtStorageClassesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListKubevirtStorageClassesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubevirtStorageClasses default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubevirtStorageClassesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubevirtStorageClasses default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubevirtStorageClassesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListKubevirtStorageClassesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
