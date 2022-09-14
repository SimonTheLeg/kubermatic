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

// GetGroupProjectBindingReader is a Reader for the GetGroupProjectBinding structure.
type GetGroupProjectBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupProjectBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupProjectBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGroupProjectBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupProjectBindingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetGroupProjectBindingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGroupProjectBindingOK creates a GetGroupProjectBindingOK with default headers values
func NewGetGroupProjectBindingOK() *GetGroupProjectBindingOK {
	return &GetGroupProjectBindingOK{}
}

/* GetGroupProjectBindingOK describes a response with status code 200, with default header values.

GroupProjectBinding
*/
type GetGroupProjectBindingOK struct {
	Payload *models.GroupProjectBinding
}

// IsSuccess returns true when this get group project binding o k response has a 2xx status code
func (o *GetGroupProjectBindingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get group project binding o k response has a 3xx status code
func (o *GetGroupProjectBindingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group project binding o k response has a 4xx status code
func (o *GetGroupProjectBindingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group project binding o k response has a 5xx status code
func (o *GetGroupProjectBindingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get group project binding o k response a status code equal to that given
func (o *GetGroupProjectBindingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGroupProjectBindingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBindingOK  %+v", 200, o.Payload)
}

func (o *GetGroupProjectBindingOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBindingOK  %+v", 200, o.Payload)
}

func (o *GetGroupProjectBindingOK) GetPayload() *models.GroupProjectBinding {
	return o.Payload
}

func (o *GetGroupProjectBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupProjectBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupProjectBindingUnauthorized creates a GetGroupProjectBindingUnauthorized with default headers values
func NewGetGroupProjectBindingUnauthorized() *GetGroupProjectBindingUnauthorized {
	return &GetGroupProjectBindingUnauthorized{}
}

/* GetGroupProjectBindingUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetGroupProjectBindingUnauthorized struct {
}

// IsSuccess returns true when this get group project binding unauthorized response has a 2xx status code
func (o *GetGroupProjectBindingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group project binding unauthorized response has a 3xx status code
func (o *GetGroupProjectBindingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group project binding unauthorized response has a 4xx status code
func (o *GetGroupProjectBindingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group project binding unauthorized response has a 5xx status code
func (o *GetGroupProjectBindingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get group project binding unauthorized response a status code equal to that given
func (o *GetGroupProjectBindingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGroupProjectBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBindingUnauthorized ", 401)
}

func (o *GetGroupProjectBindingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBindingUnauthorized ", 401)
}

func (o *GetGroupProjectBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupProjectBindingForbidden creates a GetGroupProjectBindingForbidden with default headers values
func NewGetGroupProjectBindingForbidden() *GetGroupProjectBindingForbidden {
	return &GetGroupProjectBindingForbidden{}
}

/* GetGroupProjectBindingForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetGroupProjectBindingForbidden struct {
}

// IsSuccess returns true when this get group project binding forbidden response has a 2xx status code
func (o *GetGroupProjectBindingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group project binding forbidden response has a 3xx status code
func (o *GetGroupProjectBindingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group project binding forbidden response has a 4xx status code
func (o *GetGroupProjectBindingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group project binding forbidden response has a 5xx status code
func (o *GetGroupProjectBindingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get group project binding forbidden response a status code equal to that given
func (o *GetGroupProjectBindingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGroupProjectBindingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBindingForbidden ", 403)
}

func (o *GetGroupProjectBindingForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBindingForbidden ", 403)
}

func (o *GetGroupProjectBindingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGroupProjectBindingDefault creates a GetGroupProjectBindingDefault with default headers values
func NewGetGroupProjectBindingDefault(code int) *GetGroupProjectBindingDefault {
	return &GetGroupProjectBindingDefault{
		_statusCode: code,
	}
}

/* GetGroupProjectBindingDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetGroupProjectBindingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get group project binding default response
func (o *GetGroupProjectBindingDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get group project binding default response has a 2xx status code
func (o *GetGroupProjectBindingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get group project binding default response has a 3xx status code
func (o *GetGroupProjectBindingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get group project binding default response has a 4xx status code
func (o *GetGroupProjectBindingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get group project binding default response has a 5xx status code
func (o *GetGroupProjectBindingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get group project binding default response a status code equal to that given
func (o *GetGroupProjectBindingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetGroupProjectBindingDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBinding default  %+v", o._statusCode, o.Payload)
}

func (o *GetGroupProjectBindingDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/groupbindings/{binding_name}][%d] getGroupProjectBinding default  %+v", o._statusCode, o.Payload)
}

func (o *GetGroupProjectBindingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetGroupProjectBindingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
