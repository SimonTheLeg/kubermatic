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

// ListNamespaceReader is a Reader for the ListNamespace structure.
type ListNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNamespaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNamespaceOK creates a ListNamespaceOK with default headers values
func NewListNamespaceOK() *ListNamespaceOK {
	return &ListNamespaceOK{}
}

/* ListNamespaceOK describes a response with status code 200, with default header values.

Namespace
*/
type ListNamespaceOK struct {
	Payload []*models.Namespace
}

// IsSuccess returns true when this list namespace o k response has a 2xx status code
func (o *ListNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list namespace o k response has a 3xx status code
func (o *ListNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list namespace o k response has a 4xx status code
func (o *ListNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list namespace o k response has a 5xx status code
func (o *ListNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list namespace o k response a status code equal to that given
func (o *ListNamespaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespaceOK  %+v", 200, o.Payload)
}

func (o *ListNamespaceOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespaceOK  %+v", 200, o.Payload)
}

func (o *ListNamespaceOK) GetPayload() []*models.Namespace {
	return o.Payload
}

func (o *ListNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNamespaceUnauthorized creates a ListNamespaceUnauthorized with default headers values
func NewListNamespaceUnauthorized() *ListNamespaceUnauthorized {
	return &ListNamespaceUnauthorized{}
}

/* ListNamespaceUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListNamespaceUnauthorized struct {
}

// IsSuccess returns true when this list namespace unauthorized response has a 2xx status code
func (o *ListNamespaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list namespace unauthorized response has a 3xx status code
func (o *ListNamespaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list namespace unauthorized response has a 4xx status code
func (o *ListNamespaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list namespace unauthorized response has a 5xx status code
func (o *ListNamespaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list namespace unauthorized response a status code equal to that given
func (o *ListNamespaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespaceUnauthorized ", 401)
}

func (o *ListNamespaceUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespaceUnauthorized ", 401)
}

func (o *ListNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNamespaceForbidden creates a ListNamespaceForbidden with default headers values
func NewListNamespaceForbidden() *ListNamespaceForbidden {
	return &ListNamespaceForbidden{}
}

/* ListNamespaceForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListNamespaceForbidden struct {
}

// IsSuccess returns true when this list namespace forbidden response has a 2xx status code
func (o *ListNamespaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list namespace forbidden response has a 3xx status code
func (o *ListNamespaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list namespace forbidden response has a 4xx status code
func (o *ListNamespaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list namespace forbidden response has a 5xx status code
func (o *ListNamespaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list namespace forbidden response a status code equal to that given
func (o *ListNamespaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListNamespaceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespaceForbidden ", 403)
}

func (o *ListNamespaceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespaceForbidden ", 403)
}

func (o *ListNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNamespaceDefault creates a ListNamespaceDefault with default headers values
func NewListNamespaceDefault(code int) *ListNamespaceDefault {
	return &ListNamespaceDefault{
		_statusCode: code,
	}
}

/* ListNamespaceDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListNamespaceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list namespace default response
func (o *ListNamespaceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list namespace default response has a 2xx status code
func (o *ListNamespaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list namespace default response has a 3xx status code
func (o *ListNamespaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list namespace default response has a 4xx status code
func (o *ListNamespaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list namespace default response has a 5xx status code
func (o *ListNamespaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list namespace default response a status code equal to that given
func (o *ListNamespaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListNamespaceDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *ListNamespaceDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces][%d] listNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *ListNamespaceDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNamespaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
