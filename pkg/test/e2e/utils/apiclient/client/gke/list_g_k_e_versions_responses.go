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

// ListGKEVersionsReader is a Reader for the ListGKEVersions structure.
type ListGKEVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGKEVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGKEVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGKEVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGKEVersionsOK creates a ListGKEVersionsOK with default headers values
func NewListGKEVersionsOK() *ListGKEVersionsOK {
	return &ListGKEVersionsOK{}
}

/*
ListGKEVersionsOK describes a response with status code 200, with default header values.

MasterVersion
*/
type ListGKEVersionsOK struct {
	Payload []*models.MasterVersion
}

// IsSuccess returns true when this list g k e versions o k response has a 2xx status code
func (o *ListGKEVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g k e versions o k response has a 3xx status code
func (o *ListGKEVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g k e versions o k response has a 4xx status code
func (o *ListGKEVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g k e versions o k response has a 5xx status code
func (o *ListGKEVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g k e versions o k response a status code equal to that given
func (o *ListGKEVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGKEVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/versions][%d] listGKEVersionsOK  %+v", 200, o.Payload)
}

func (o *ListGKEVersionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/versions][%d] listGKEVersionsOK  %+v", 200, o.Payload)
}

func (o *ListGKEVersionsOK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *ListGKEVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGKEVersionsDefault creates a ListGKEVersionsDefault with default headers values
func NewListGKEVersionsDefault(code int) *ListGKEVersionsDefault {
	return &ListGKEVersionsDefault{
		_statusCode: code,
	}
}

/*
ListGKEVersionsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGKEVersionsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g k e versions default response
func (o *ListGKEVersionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g k e versions default response has a 2xx status code
func (o *ListGKEVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g k e versions default response has a 3xx status code
func (o *ListGKEVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g k e versions default response has a 4xx status code
func (o *ListGKEVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g k e versions default response has a 5xx status code
func (o *ListGKEVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g k e versions default response a status code equal to that given
func (o *ListGKEVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGKEVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/versions][%d] listGKEVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListGKEVersionsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/gke/versions][%d] listGKEVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListGKEVersionsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGKEVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
