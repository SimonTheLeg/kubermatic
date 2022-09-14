// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListConstraintsReader is a Reader for the ListConstraints structure.
type ListConstraintsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConstraintsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConstraintsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListConstraintsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListConstraintsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListConstraintsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListConstraintsOK creates a ListConstraintsOK with default headers values
func NewListConstraintsOK() *ListConstraintsOK {
	return &ListConstraintsOK{}
}

/* ListConstraintsOK describes a response with status code 200, with default header values.

Constraint
*/
type ListConstraintsOK struct {
	Payload []*models.Constraint
}

// IsSuccess returns true when this list constraints o k response has a 2xx status code
func (o *ListConstraintsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list constraints o k response has a 3xx status code
func (o *ListConstraintsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list constraints o k response has a 4xx status code
func (o *ListConstraintsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list constraints o k response has a 5xx status code
func (o *ListConstraintsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list constraints o k response a status code equal to that given
func (o *ListConstraintsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListConstraintsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraintsOK  %+v", 200, o.Payload)
}

func (o *ListConstraintsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraintsOK  %+v", 200, o.Payload)
}

func (o *ListConstraintsOK) GetPayload() []*models.Constraint {
	return o.Payload
}

func (o *ListConstraintsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConstraintsUnauthorized creates a ListConstraintsUnauthorized with default headers values
func NewListConstraintsUnauthorized() *ListConstraintsUnauthorized {
	return &ListConstraintsUnauthorized{}
}

/* ListConstraintsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListConstraintsUnauthorized struct {
}

// IsSuccess returns true when this list constraints unauthorized response has a 2xx status code
func (o *ListConstraintsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list constraints unauthorized response has a 3xx status code
func (o *ListConstraintsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list constraints unauthorized response has a 4xx status code
func (o *ListConstraintsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list constraints unauthorized response has a 5xx status code
func (o *ListConstraintsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list constraints unauthorized response a status code equal to that given
func (o *ListConstraintsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListConstraintsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraintsUnauthorized ", 401)
}

func (o *ListConstraintsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraintsUnauthorized ", 401)
}

func (o *ListConstraintsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConstraintsForbidden creates a ListConstraintsForbidden with default headers values
func NewListConstraintsForbidden() *ListConstraintsForbidden {
	return &ListConstraintsForbidden{}
}

/* ListConstraintsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListConstraintsForbidden struct {
}

// IsSuccess returns true when this list constraints forbidden response has a 2xx status code
func (o *ListConstraintsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list constraints forbidden response has a 3xx status code
func (o *ListConstraintsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list constraints forbidden response has a 4xx status code
func (o *ListConstraintsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list constraints forbidden response has a 5xx status code
func (o *ListConstraintsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list constraints forbidden response a status code equal to that given
func (o *ListConstraintsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListConstraintsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraintsForbidden ", 403)
}

func (o *ListConstraintsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraintsForbidden ", 403)
}

func (o *ListConstraintsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConstraintsDefault creates a ListConstraintsDefault with default headers values
func NewListConstraintsDefault(code int) *ListConstraintsDefault {
	return &ListConstraintsDefault{
		_statusCode: code,
	}
}

/* ListConstraintsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListConstraintsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list constraints default response
func (o *ListConstraintsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list constraints default response has a 2xx status code
func (o *ListConstraintsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list constraints default response has a 3xx status code
func (o *ListConstraintsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list constraints default response has a 4xx status code
func (o *ListConstraintsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list constraints default response has a 5xx status code
func (o *ListConstraintsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list constraints default response a status code equal to that given
func (o *ListConstraintsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListConstraintsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraints default  %+v", o._statusCode, o.Payload)
}

func (o *ListConstraintsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints][%d] listConstraints default  %+v", o._statusCode, o.Payload)
}

func (o *ListConstraintsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListConstraintsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
