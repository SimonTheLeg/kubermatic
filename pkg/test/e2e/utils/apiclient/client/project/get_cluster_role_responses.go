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

// GetClusterRoleReader is a Reader for the GetClusterRole structure.
type GetClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterRoleOK creates a GetClusterRoleOK with default headers values
func NewGetClusterRoleOK() *GetClusterRoleOK {
	return &GetClusterRoleOK{}
}

/* GetClusterRoleOK describes a response with status code 200, with default header values.

ClusterRole
*/
type GetClusterRoleOK struct {
	Payload *models.ClusterRole
}

// IsSuccess returns true when this get cluster role o k response has a 2xx status code
func (o *GetClusterRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster role o k response has a 3xx status code
func (o *GetClusterRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster role o k response has a 4xx status code
func (o *GetClusterRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster role o k response has a 5xx status code
func (o *GetClusterRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster role o k response a status code equal to that given
func (o *GetClusterRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterRoleOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRoleOK  %+v", 200, o.Payload)
}

func (o *GetClusterRoleOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRoleOK  %+v", 200, o.Payload)
}

func (o *GetClusterRoleOK) GetPayload() *models.ClusterRole {
	return o.Payload
}

func (o *GetClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterRoleUnauthorized creates a GetClusterRoleUnauthorized with default headers values
func NewGetClusterRoleUnauthorized() *GetClusterRoleUnauthorized {
	return &GetClusterRoleUnauthorized{}
}

/* GetClusterRoleUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterRoleUnauthorized struct {
}

// IsSuccess returns true when this get cluster role unauthorized response has a 2xx status code
func (o *GetClusterRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster role unauthorized response has a 3xx status code
func (o *GetClusterRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster role unauthorized response has a 4xx status code
func (o *GetClusterRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster role unauthorized response has a 5xx status code
func (o *GetClusterRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster role unauthorized response a status code equal to that given
func (o *GetClusterRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRoleUnauthorized ", 401)
}

func (o *GetClusterRoleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRoleUnauthorized ", 401)
}

func (o *GetClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterRoleForbidden creates a GetClusterRoleForbidden with default headers values
func NewGetClusterRoleForbidden() *GetClusterRoleForbidden {
	return &GetClusterRoleForbidden{}
}

/* GetClusterRoleForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterRoleForbidden struct {
}

// IsSuccess returns true when this get cluster role forbidden response has a 2xx status code
func (o *GetClusterRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster role forbidden response has a 3xx status code
func (o *GetClusterRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster role forbidden response has a 4xx status code
func (o *GetClusterRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster role forbidden response has a 5xx status code
func (o *GetClusterRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster role forbidden response a status code equal to that given
func (o *GetClusterRoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRoleForbidden ", 403)
}

func (o *GetClusterRoleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRoleForbidden ", 403)
}

func (o *GetClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterRoleDefault creates a GetClusterRoleDefault with default headers values
func NewGetClusterRoleDefault(code int) *GetClusterRoleDefault {
	return &GetClusterRoleDefault{
		_statusCode: code,
	}
}

/* GetClusterRoleDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster role default response
func (o *GetClusterRoleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster role default response has a 2xx status code
func (o *GetClusterRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster role default response has a 3xx status code
func (o *GetClusterRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster role default response has a 4xx status code
func (o *GetClusterRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster role default response has a 5xx status code
func (o *GetClusterRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster role default response a status code equal to that given
func (o *GetClusterRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterRoleDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterRoleDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}][%d] getClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
