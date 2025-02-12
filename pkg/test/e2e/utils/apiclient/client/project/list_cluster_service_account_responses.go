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

// ListClusterServiceAccountReader is a Reader for the ListClusterServiceAccount structure.
type ListClusterServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClusterServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClusterServiceAccountOK creates a ListClusterServiceAccountOK with default headers values
func NewListClusterServiceAccountOK() *ListClusterServiceAccountOK {
	return &ListClusterServiceAccountOK{}
}

/*
ListClusterServiceAccountOK describes a response with status code 200, with default header values.

ClusterServiceAccount
*/
type ListClusterServiceAccountOK struct {
	Payload []*models.ClusterServiceAccount
}

// IsSuccess returns true when this list cluster service account o k response has a 2xx status code
func (o *ListClusterServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list cluster service account o k response has a 3xx status code
func (o *ListClusterServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster service account o k response has a 4xx status code
func (o *ListClusterServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list cluster service account o k response has a 5xx status code
func (o *ListClusterServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster service account o k response a status code equal to that given
func (o *ListClusterServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListClusterServiceAccountOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccountOK  %+v", 200, o.Payload)
}

func (o *ListClusterServiceAccountOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccountOK  %+v", 200, o.Payload)
}

func (o *ListClusterServiceAccountOK) GetPayload() []*models.ClusterServiceAccount {
	return o.Payload
}

func (o *ListClusterServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterServiceAccountUnauthorized creates a ListClusterServiceAccountUnauthorized with default headers values
func NewListClusterServiceAccountUnauthorized() *ListClusterServiceAccountUnauthorized {
	return &ListClusterServiceAccountUnauthorized{}
}

/*
ListClusterServiceAccountUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListClusterServiceAccountUnauthorized struct {
}

// IsSuccess returns true when this list cluster service account unauthorized response has a 2xx status code
func (o *ListClusterServiceAccountUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster service account unauthorized response has a 3xx status code
func (o *ListClusterServiceAccountUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster service account unauthorized response has a 4xx status code
func (o *ListClusterServiceAccountUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster service account unauthorized response has a 5xx status code
func (o *ListClusterServiceAccountUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster service account unauthorized response a status code equal to that given
func (o *ListClusterServiceAccountUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListClusterServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccountUnauthorized ", 401)
}

func (o *ListClusterServiceAccountUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccountUnauthorized ", 401)
}

func (o *ListClusterServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterServiceAccountForbidden creates a ListClusterServiceAccountForbidden with default headers values
func NewListClusterServiceAccountForbidden() *ListClusterServiceAccountForbidden {
	return &ListClusterServiceAccountForbidden{}
}

/*
ListClusterServiceAccountForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListClusterServiceAccountForbidden struct {
}

// IsSuccess returns true when this list cluster service account forbidden response has a 2xx status code
func (o *ListClusterServiceAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list cluster service account forbidden response has a 3xx status code
func (o *ListClusterServiceAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cluster service account forbidden response has a 4xx status code
func (o *ListClusterServiceAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list cluster service account forbidden response has a 5xx status code
func (o *ListClusterServiceAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list cluster service account forbidden response a status code equal to that given
func (o *ListClusterServiceAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListClusterServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccountForbidden ", 403)
}

func (o *ListClusterServiceAccountForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccountForbidden ", 403)
}

func (o *ListClusterServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterServiceAccountDefault creates a ListClusterServiceAccountDefault with default headers values
func NewListClusterServiceAccountDefault(code int) *ListClusterServiceAccountDefault {
	return &ListClusterServiceAccountDefault{
		_statusCode: code,
	}
}

/*
ListClusterServiceAccountDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListClusterServiceAccountDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list cluster service account default response
func (o *ListClusterServiceAccountDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list cluster service account default response has a 2xx status code
func (o *ListClusterServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list cluster service account default response has a 3xx status code
func (o *ListClusterServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list cluster service account default response has a 4xx status code
func (o *ListClusterServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list cluster service account default response has a 5xx status code
func (o *ListClusterServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list cluster service account default response a status code equal to that given
func (o *ListClusterServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListClusterServiceAccountDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterServiceAccountDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/serviceaccount][%d] listClusterServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterServiceAccountDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClusterServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
