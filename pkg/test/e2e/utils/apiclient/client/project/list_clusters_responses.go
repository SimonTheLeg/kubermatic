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

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/*
ListClustersOK describes a response with status code 200, with default header values.

ClusterList
*/
type ListClustersOK struct {
	Payload models.ClusterList
}

// IsSuccess returns true when this list clusters o k response has a 2xx status code
func (o *ListClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list clusters o k response has a 3xx status code
func (o *ListClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters o k response has a 4xx status code
func (o *ListClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list clusters o k response has a 5xx status code
func (o *ListClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list clusters o k response a status code equal to that given
func (o *ListClustersOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) GetPayload() models.ClusterList {
	return o.Payload
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersUnauthorized creates a ListClustersUnauthorized with default headers values
func NewListClustersUnauthorized() *ListClustersUnauthorized {
	return &ListClustersUnauthorized{}
}

/*
ListClustersUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListClustersUnauthorized struct {
}

// IsSuccess returns true when this list clusters unauthorized response has a 2xx status code
func (o *ListClustersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list clusters unauthorized response has a 3xx status code
func (o *ListClustersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters unauthorized response has a 4xx status code
func (o *ListClustersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list clusters unauthorized response has a 5xx status code
func (o *ListClustersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list clusters unauthorized response a status code equal to that given
func (o *ListClustersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersUnauthorized ", 401)
}

func (o *ListClustersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersUnauthorized ", 401)
}

func (o *ListClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersForbidden creates a ListClustersForbidden with default headers values
func NewListClustersForbidden() *ListClustersForbidden {
	return &ListClustersForbidden{}
}

/*
ListClustersForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListClustersForbidden struct {
}

// IsSuccess returns true when this list clusters forbidden response has a 2xx status code
func (o *ListClustersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list clusters forbidden response has a 3xx status code
func (o *ListClustersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list clusters forbidden response has a 4xx status code
func (o *ListClustersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list clusters forbidden response has a 5xx status code
func (o *ListClustersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list clusters forbidden response a status code equal to that given
func (o *ListClustersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersForbidden ", 403)
}

func (o *ListClustersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersForbidden ", 403)
}

func (o *ListClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersDefault creates a ListClustersDefault with default headers values
func NewListClustersDefault(code int) *ListClustersDefault {
	return &ListClustersDefault{
		_statusCode: code,
	}
}

/*
ListClustersDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListClustersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list clusters default response
func (o *ListClustersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list clusters default response has a 2xx status code
func (o *ListClustersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list clusters default response has a 3xx status code
func (o *ListClustersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list clusters default response has a 4xx status code
func (o *ListClustersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list clusters default response has a 5xx status code
func (o *ListClustersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list clusters default response a status code equal to that given
func (o *ListClustersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListClustersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListClustersDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListClustersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
