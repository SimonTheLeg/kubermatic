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

// ListGKEClusterDiskTypesReader is a Reader for the ListGKEClusterDiskTypes structure.
type ListGKEClusterDiskTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGKEClusterDiskTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGKEClusterDiskTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListGKEClusterDiskTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListGKEClusterDiskTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListGKEClusterDiskTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGKEClusterDiskTypesOK creates a ListGKEClusterDiskTypesOK with default headers values
func NewListGKEClusterDiskTypesOK() *ListGKEClusterDiskTypesOK {
	return &ListGKEClusterDiskTypesOK{}
}

/* ListGKEClusterDiskTypesOK describes a response with status code 200, with default header values.

GCPDiskTypeList
*/
type ListGKEClusterDiskTypesOK struct {
	Payload models.GCPDiskTypeList
}

// IsSuccess returns true when this list g k e cluster disk types o k response has a 2xx status code
func (o *ListGKEClusterDiskTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list g k e cluster disk types o k response has a 3xx status code
func (o *ListGKEClusterDiskTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g k e cluster disk types o k response has a 4xx status code
func (o *ListGKEClusterDiskTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list g k e cluster disk types o k response has a 5xx status code
func (o *ListGKEClusterDiskTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list g k e cluster disk types o k response a status code equal to that given
func (o *ListGKEClusterDiskTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListGKEClusterDiskTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypesOK  %+v", 200, o.Payload)
}

func (o *ListGKEClusterDiskTypesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypesOK  %+v", 200, o.Payload)
}

func (o *ListGKEClusterDiskTypesOK) GetPayload() models.GCPDiskTypeList {
	return o.Payload
}

func (o *ListGKEClusterDiskTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGKEClusterDiskTypesUnauthorized creates a ListGKEClusterDiskTypesUnauthorized with default headers values
func NewListGKEClusterDiskTypesUnauthorized() *ListGKEClusterDiskTypesUnauthorized {
	return &ListGKEClusterDiskTypesUnauthorized{}
}

/* ListGKEClusterDiskTypesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListGKEClusterDiskTypesUnauthorized struct {
}

// IsSuccess returns true when this list g k e cluster disk types unauthorized response has a 2xx status code
func (o *ListGKEClusterDiskTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list g k e cluster disk types unauthorized response has a 3xx status code
func (o *ListGKEClusterDiskTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g k e cluster disk types unauthorized response has a 4xx status code
func (o *ListGKEClusterDiskTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list g k e cluster disk types unauthorized response has a 5xx status code
func (o *ListGKEClusterDiskTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list g k e cluster disk types unauthorized response a status code equal to that given
func (o *ListGKEClusterDiskTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListGKEClusterDiskTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypesUnauthorized ", 401)
}

func (o *ListGKEClusterDiskTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypesUnauthorized ", 401)
}

func (o *ListGKEClusterDiskTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListGKEClusterDiskTypesForbidden creates a ListGKEClusterDiskTypesForbidden with default headers values
func NewListGKEClusterDiskTypesForbidden() *ListGKEClusterDiskTypesForbidden {
	return &ListGKEClusterDiskTypesForbidden{}
}

/* ListGKEClusterDiskTypesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListGKEClusterDiskTypesForbidden struct {
}

// IsSuccess returns true when this list g k e cluster disk types forbidden response has a 2xx status code
func (o *ListGKEClusterDiskTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list g k e cluster disk types forbidden response has a 3xx status code
func (o *ListGKEClusterDiskTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list g k e cluster disk types forbidden response has a 4xx status code
func (o *ListGKEClusterDiskTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list g k e cluster disk types forbidden response has a 5xx status code
func (o *ListGKEClusterDiskTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list g k e cluster disk types forbidden response a status code equal to that given
func (o *ListGKEClusterDiskTypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListGKEClusterDiskTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypesForbidden ", 403)
}

func (o *ListGKEClusterDiskTypesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypesForbidden ", 403)
}

func (o *ListGKEClusterDiskTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListGKEClusterDiskTypesDefault creates a ListGKEClusterDiskTypesDefault with default headers values
func NewListGKEClusterDiskTypesDefault(code int) *ListGKEClusterDiskTypesDefault {
	return &ListGKEClusterDiskTypesDefault{
		_statusCode: code,
	}
}

/* ListGKEClusterDiskTypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListGKEClusterDiskTypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g k e cluster disk types default response
func (o *ListGKEClusterDiskTypesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list g k e cluster disk types default response has a 2xx status code
func (o *ListGKEClusterDiskTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list g k e cluster disk types default response has a 3xx status code
func (o *ListGKEClusterDiskTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list g k e cluster disk types default response has a 4xx status code
func (o *ListGKEClusterDiskTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list g k e cluster disk types default response has a 5xx status code
func (o *ListGKEClusterDiskTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list g k e cluster disk types default response a status code equal to that given
func (o *ListGKEClusterDiskTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListGKEClusterDiskTypesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListGKEClusterDiskTypesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes][%d] listGKEClusterDiskTypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListGKEClusterDiskTypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGKEClusterDiskTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
