// Code generated by go-swagger; DO NOT EDIT.

package aks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAKSNodePoolModesReader is a Reader for the ListAKSNodePoolModes structure.
type ListAKSNodePoolModesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAKSNodePoolModesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAKSNodePoolModesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAKSNodePoolModesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAKSNodePoolModesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAKSNodePoolModesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAKSNodePoolModesOK creates a ListAKSNodePoolModesOK with default headers values
func NewListAKSNodePoolModesOK() *ListAKSNodePoolModesOK {
	return &ListAKSNodePoolModesOK{}
}

/*
ListAKSNodePoolModesOK describes a response with status code 200, with default header values.

AKSNodePoolModes
*/
type ListAKSNodePoolModesOK struct {
	Payload models.AKSNodePoolModes
}

// IsSuccess returns true when this list a k s node pool modes o k response has a 2xx status code
func (o *ListAKSNodePoolModesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list a k s node pool modes o k response has a 3xx status code
func (o *ListAKSNodePoolModesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s node pool modes o k response has a 4xx status code
func (o *ListAKSNodePoolModesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list a k s node pool modes o k response has a 5xx status code
func (o *ListAKSNodePoolModesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s node pool modes o k response a status code equal to that given
func (o *ListAKSNodePoolModesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAKSNodePoolModesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModesOK  %+v", 200, o.Payload)
}

func (o *ListAKSNodePoolModesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModesOK  %+v", 200, o.Payload)
}

func (o *ListAKSNodePoolModesOK) GetPayload() models.AKSNodePoolModes {
	return o.Payload
}

func (o *ListAKSNodePoolModesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAKSNodePoolModesUnauthorized creates a ListAKSNodePoolModesUnauthorized with default headers values
func NewListAKSNodePoolModesUnauthorized() *ListAKSNodePoolModesUnauthorized {
	return &ListAKSNodePoolModesUnauthorized{}
}

/*
ListAKSNodePoolModesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListAKSNodePoolModesUnauthorized struct {
}

// IsSuccess returns true when this list a k s node pool modes unauthorized response has a 2xx status code
func (o *ListAKSNodePoolModesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list a k s node pool modes unauthorized response has a 3xx status code
func (o *ListAKSNodePoolModesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s node pool modes unauthorized response has a 4xx status code
func (o *ListAKSNodePoolModesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list a k s node pool modes unauthorized response has a 5xx status code
func (o *ListAKSNodePoolModesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s node pool modes unauthorized response a status code equal to that given
func (o *ListAKSNodePoolModesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListAKSNodePoolModesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModesUnauthorized ", 401)
}

func (o *ListAKSNodePoolModesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModesUnauthorized ", 401)
}

func (o *ListAKSNodePoolModesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAKSNodePoolModesForbidden creates a ListAKSNodePoolModesForbidden with default headers values
func NewListAKSNodePoolModesForbidden() *ListAKSNodePoolModesForbidden {
	return &ListAKSNodePoolModesForbidden{}
}

/*
ListAKSNodePoolModesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListAKSNodePoolModesForbidden struct {
}

// IsSuccess returns true when this list a k s node pool modes forbidden response has a 2xx status code
func (o *ListAKSNodePoolModesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list a k s node pool modes forbidden response has a 3xx status code
func (o *ListAKSNodePoolModesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a k s node pool modes forbidden response has a 4xx status code
func (o *ListAKSNodePoolModesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list a k s node pool modes forbidden response has a 5xx status code
func (o *ListAKSNodePoolModesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list a k s node pool modes forbidden response a status code equal to that given
func (o *ListAKSNodePoolModesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListAKSNodePoolModesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModesForbidden ", 403)
}

func (o *ListAKSNodePoolModesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModesForbidden ", 403)
}

func (o *ListAKSNodePoolModesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAKSNodePoolModesDefault creates a ListAKSNodePoolModesDefault with default headers values
func NewListAKSNodePoolModesDefault(code int) *ListAKSNodePoolModesDefault {
	return &ListAKSNodePoolModesDefault{
		_statusCode: code,
	}
}

/*
ListAKSNodePoolModesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAKSNodePoolModesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a k s node pool modes default response
func (o *ListAKSNodePoolModesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list a k s node pool modes default response has a 2xx status code
func (o *ListAKSNodePoolModesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list a k s node pool modes default response has a 3xx status code
func (o *ListAKSNodePoolModesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list a k s node pool modes default response has a 4xx status code
func (o *ListAKSNodePoolModesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list a k s node pool modes default response has a 5xx status code
func (o *ListAKSNodePoolModesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list a k s node pool modes default response a status code equal to that given
func (o *ListAKSNodePoolModesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAKSNodePoolModesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModes default  %+v", o._statusCode, o.Payload)
}

func (o *ListAKSNodePoolModesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/aks/modes][%d] listAKSNodePoolModes default  %+v", o._statusCode, o.Payload)
}

func (o *ListAKSNodePoolModesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAKSNodePoolModesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
