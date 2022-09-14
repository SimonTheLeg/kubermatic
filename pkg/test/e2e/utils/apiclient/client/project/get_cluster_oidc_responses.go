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

// GetClusterOidcReader is a Reader for the GetClusterOidc structure.
type GetClusterOidcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterOidcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterOidcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterOidcUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterOidcForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterOidcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterOidcOK creates a GetClusterOidcOK with default headers values
func NewGetClusterOidcOK() *GetClusterOidcOK {
	return &GetClusterOidcOK{}
}

/* GetClusterOidcOK describes a response with status code 200, with default header values.

OIDCSpec
*/
type GetClusterOidcOK struct {
	Payload *models.OIDCSpec
}

// IsSuccess returns true when this get cluster oidc o k response has a 2xx status code
func (o *GetClusterOidcOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster oidc o k response has a 3xx status code
func (o *GetClusterOidcOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster oidc o k response has a 4xx status code
func (o *GetClusterOidcOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster oidc o k response has a 5xx status code
func (o *GetClusterOidcOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster oidc o k response a status code equal to that given
func (o *GetClusterOidcOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterOidcOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidcOK  %+v", 200, o.Payload)
}

func (o *GetClusterOidcOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidcOK  %+v", 200, o.Payload)
}

func (o *GetClusterOidcOK) GetPayload() *models.OIDCSpec {
	return o.Payload
}

func (o *GetClusterOidcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OIDCSpec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterOidcUnauthorized creates a GetClusterOidcUnauthorized with default headers values
func NewGetClusterOidcUnauthorized() *GetClusterOidcUnauthorized {
	return &GetClusterOidcUnauthorized{}
}

/* GetClusterOidcUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterOidcUnauthorized struct {
}

// IsSuccess returns true when this get cluster oidc unauthorized response has a 2xx status code
func (o *GetClusterOidcUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster oidc unauthorized response has a 3xx status code
func (o *GetClusterOidcUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster oidc unauthorized response has a 4xx status code
func (o *GetClusterOidcUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster oidc unauthorized response has a 5xx status code
func (o *GetClusterOidcUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster oidc unauthorized response a status code equal to that given
func (o *GetClusterOidcUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterOidcUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidcUnauthorized ", 401)
}

func (o *GetClusterOidcUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidcUnauthorized ", 401)
}

func (o *GetClusterOidcUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterOidcForbidden creates a GetClusterOidcForbidden with default headers values
func NewGetClusterOidcForbidden() *GetClusterOidcForbidden {
	return &GetClusterOidcForbidden{}
}

/* GetClusterOidcForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterOidcForbidden struct {
}

// IsSuccess returns true when this get cluster oidc forbidden response has a 2xx status code
func (o *GetClusterOidcForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster oidc forbidden response has a 3xx status code
func (o *GetClusterOidcForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster oidc forbidden response has a 4xx status code
func (o *GetClusterOidcForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster oidc forbidden response has a 5xx status code
func (o *GetClusterOidcForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster oidc forbidden response a status code equal to that given
func (o *GetClusterOidcForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterOidcForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidcForbidden ", 403)
}

func (o *GetClusterOidcForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidcForbidden ", 403)
}

func (o *GetClusterOidcForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterOidcDefault creates a GetClusterOidcDefault with default headers values
func NewGetClusterOidcDefault(code int) *GetClusterOidcDefault {
	return &GetClusterOidcDefault{
		_statusCode: code,
	}
}

/* GetClusterOidcDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterOidcDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster oidc default response
func (o *GetClusterOidcDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster oidc default response has a 2xx status code
func (o *GetClusterOidcDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster oidc default response has a 3xx status code
func (o *GetClusterOidcDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster oidc default response has a 4xx status code
func (o *GetClusterOidcDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster oidc default response has a 5xx status code
func (o *GetClusterOidcDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster oidc default response a status code equal to that given
func (o *GetClusterOidcDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterOidcDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidc default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterOidcDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc][%d] getClusterOidc default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterOidcDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterOidcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
