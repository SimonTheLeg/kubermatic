// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// AddUserToProjectReader is a Reader for the AddUserToProject structure.
type AddUserToProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserToProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddUserToProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddUserToProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddUserToProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddUserToProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddUserToProjectCreated creates a AddUserToProjectCreated with default headers values
func NewAddUserToProjectCreated() *AddUserToProjectCreated {
	return &AddUserToProjectCreated{}
}

/* AddUserToProjectCreated describes a response with status code 201, with default header values.

User
*/
type AddUserToProjectCreated struct {
	Payload *models.User
}

// IsSuccess returns true when this add user to project created response has a 2xx status code
func (o *AddUserToProjectCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user to project created response has a 3xx status code
func (o *AddUserToProjectCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to project created response has a 4xx status code
func (o *AddUserToProjectCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user to project created response has a 5xx status code
func (o *AddUserToProjectCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to project created response a status code equal to that given
func (o *AddUserToProjectCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddUserToProjectCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProjectCreated  %+v", 201, o.Payload)
}

func (o *AddUserToProjectCreated) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProjectCreated  %+v", 201, o.Payload)
}

func (o *AddUserToProjectCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *AddUserToProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToProjectUnauthorized creates a AddUserToProjectUnauthorized with default headers values
func NewAddUserToProjectUnauthorized() *AddUserToProjectUnauthorized {
	return &AddUserToProjectUnauthorized{}
}

/* AddUserToProjectUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type AddUserToProjectUnauthorized struct {
}

// IsSuccess returns true when this add user to project unauthorized response has a 2xx status code
func (o *AddUserToProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to project unauthorized response has a 3xx status code
func (o *AddUserToProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to project unauthorized response has a 4xx status code
func (o *AddUserToProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user to project unauthorized response has a 5xx status code
func (o *AddUserToProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to project unauthorized response a status code equal to that given
func (o *AddUserToProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddUserToProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProjectUnauthorized ", 401)
}

func (o *AddUserToProjectUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProjectUnauthorized ", 401)
}

func (o *AddUserToProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserToProjectForbidden creates a AddUserToProjectForbidden with default headers values
func NewAddUserToProjectForbidden() *AddUserToProjectForbidden {
	return &AddUserToProjectForbidden{}
}

/* AddUserToProjectForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type AddUserToProjectForbidden struct {
}

// IsSuccess returns true when this add user to project forbidden response has a 2xx status code
func (o *AddUserToProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user to project forbidden response has a 3xx status code
func (o *AddUserToProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user to project forbidden response has a 4xx status code
func (o *AddUserToProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user to project forbidden response has a 5xx status code
func (o *AddUserToProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add user to project forbidden response a status code equal to that given
func (o *AddUserToProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddUserToProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProjectForbidden ", 403)
}

func (o *AddUserToProjectForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProjectForbidden ", 403)
}

func (o *AddUserToProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddUserToProjectDefault creates a AddUserToProjectDefault with default headers values
func NewAddUserToProjectDefault(code int) *AddUserToProjectDefault {
	return &AddUserToProjectDefault{
		_statusCode: code,
	}
}

/* AddUserToProjectDefault describes a response with status code -1, with default header values.

errorResponse
*/
type AddUserToProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add user to project default response
func (o *AddUserToProjectDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add user to project default response has a 2xx status code
func (o *AddUserToProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add user to project default response has a 3xx status code
func (o *AddUserToProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add user to project default response has a 4xx status code
func (o *AddUserToProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add user to project default response has a 5xx status code
func (o *AddUserToProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add user to project default response a status code equal to that given
func (o *AddUserToProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddUserToProjectDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProject default  %+v", o._statusCode, o.Payload)
}

func (o *AddUserToProjectDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/users][%d] addUserToProject default  %+v", o._statusCode, o.Payload)
}

func (o *AddUserToProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUserToProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
