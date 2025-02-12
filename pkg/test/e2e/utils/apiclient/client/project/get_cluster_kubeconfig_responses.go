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

// GetClusterKubeconfigReader is a Reader for the GetClusterKubeconfig structure.
type GetClusterKubeconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterKubeconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterKubeconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterKubeconfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterKubeconfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterKubeconfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterKubeconfigOK creates a GetClusterKubeconfigOK with default headers values
func NewGetClusterKubeconfigOK() *GetClusterKubeconfigOK {
	return &GetClusterKubeconfigOK{}
}

/*
GetClusterKubeconfigOK describes a response with status code 200, with default header values.

Kubeconfig is a clusters kubeconfig
*/
type GetClusterKubeconfigOK struct {
	Payload []uint8
}

// IsSuccess returns true when this get cluster kubeconfig o k response has a 2xx status code
func (o *GetClusterKubeconfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster kubeconfig o k response has a 3xx status code
func (o *GetClusterKubeconfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster kubeconfig o k response has a 4xx status code
func (o *GetClusterKubeconfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster kubeconfig o k response has a 5xx status code
func (o *GetClusterKubeconfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster kubeconfig o k response a status code equal to that given
func (o *GetClusterKubeconfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterKubeconfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigOK  %+v", 200, o.Payload)
}

func (o *GetClusterKubeconfigOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigOK  %+v", 200, o.Payload)
}

func (o *GetClusterKubeconfigOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *GetClusterKubeconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterKubeconfigUnauthorized creates a GetClusterKubeconfigUnauthorized with default headers values
func NewGetClusterKubeconfigUnauthorized() *GetClusterKubeconfigUnauthorized {
	return &GetClusterKubeconfigUnauthorized{}
}

/*
GetClusterKubeconfigUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterKubeconfigUnauthorized struct {
}

// IsSuccess returns true when this get cluster kubeconfig unauthorized response has a 2xx status code
func (o *GetClusterKubeconfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster kubeconfig unauthorized response has a 3xx status code
func (o *GetClusterKubeconfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster kubeconfig unauthorized response has a 4xx status code
func (o *GetClusterKubeconfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster kubeconfig unauthorized response has a 5xx status code
func (o *GetClusterKubeconfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster kubeconfig unauthorized response a status code equal to that given
func (o *GetClusterKubeconfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterKubeconfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigUnauthorized ", 401)
}

func (o *GetClusterKubeconfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigUnauthorized ", 401)
}

func (o *GetClusterKubeconfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterKubeconfigForbidden creates a GetClusterKubeconfigForbidden with default headers values
func NewGetClusterKubeconfigForbidden() *GetClusterKubeconfigForbidden {
	return &GetClusterKubeconfigForbidden{}
}

/*
GetClusterKubeconfigForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterKubeconfigForbidden struct {
}

// IsSuccess returns true when this get cluster kubeconfig forbidden response has a 2xx status code
func (o *GetClusterKubeconfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster kubeconfig forbidden response has a 3xx status code
func (o *GetClusterKubeconfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster kubeconfig forbidden response has a 4xx status code
func (o *GetClusterKubeconfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster kubeconfig forbidden response has a 5xx status code
func (o *GetClusterKubeconfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster kubeconfig forbidden response a status code equal to that given
func (o *GetClusterKubeconfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterKubeconfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigForbidden ", 403)
}

func (o *GetClusterKubeconfigForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfigForbidden ", 403)
}

func (o *GetClusterKubeconfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterKubeconfigDefault creates a GetClusterKubeconfigDefault with default headers values
func NewGetClusterKubeconfigDefault(code int) *GetClusterKubeconfigDefault {
	return &GetClusterKubeconfigDefault{
		_statusCode: code,
	}
}

/*
GetClusterKubeconfigDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterKubeconfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster kubeconfig default response
func (o *GetClusterKubeconfigDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster kubeconfig default response has a 2xx status code
func (o *GetClusterKubeconfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster kubeconfig default response has a 3xx status code
func (o *GetClusterKubeconfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster kubeconfig default response has a 4xx status code
func (o *GetClusterKubeconfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster kubeconfig default response has a 5xx status code
func (o *GetClusterKubeconfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster kubeconfig default response a status code equal to that given
func (o *GetClusterKubeconfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterKubeconfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterKubeconfigDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig][%d] getClusterKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterKubeconfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterKubeconfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
