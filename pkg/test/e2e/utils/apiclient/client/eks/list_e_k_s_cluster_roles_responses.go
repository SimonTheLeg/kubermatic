// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListEKSClusterRolesReader is a Reader for the ListEKSClusterRoles structure.
type ListEKSClusterRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSClusterRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSClusterRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSClusterRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSClusterRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSClusterRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSClusterRolesOK creates a ListEKSClusterRolesOK with default headers values
func NewListEKSClusterRolesOK() *ListEKSClusterRolesOK {
	return &ListEKSClusterRolesOK{}
}

/*
ListEKSClusterRolesOK describes a response with status code 200, with default header values.

EKSClusterRoleList
*/
type ListEKSClusterRolesOK struct {
	Payload models.EKSClusterRoleList
}

// IsSuccess returns true when this list e k s cluster roles o k response has a 2xx status code
func (o *ListEKSClusterRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s cluster roles o k response has a 3xx status code
func (o *ListEKSClusterRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s cluster roles o k response has a 4xx status code
func (o *ListEKSClusterRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s cluster roles o k response has a 5xx status code
func (o *ListEKSClusterRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s cluster roles o k response a status code equal to that given
func (o *ListEKSClusterRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSClusterRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesOK  %+v", 200, o.Payload)
}

func (o *ListEKSClusterRolesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesOK  %+v", 200, o.Payload)
}

func (o *ListEKSClusterRolesOK) GetPayload() models.EKSClusterRoleList {
	return o.Payload
}

func (o *ListEKSClusterRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSClusterRolesUnauthorized creates a ListEKSClusterRolesUnauthorized with default headers values
func NewListEKSClusterRolesUnauthorized() *ListEKSClusterRolesUnauthorized {
	return &ListEKSClusterRolesUnauthorized{}
}

/*
ListEKSClusterRolesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSClusterRolesUnauthorized struct {
}

// IsSuccess returns true when this list e k s cluster roles unauthorized response has a 2xx status code
func (o *ListEKSClusterRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s cluster roles unauthorized response has a 3xx status code
func (o *ListEKSClusterRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s cluster roles unauthorized response has a 4xx status code
func (o *ListEKSClusterRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s cluster roles unauthorized response has a 5xx status code
func (o *ListEKSClusterRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s cluster roles unauthorized response a status code equal to that given
func (o *ListEKSClusterRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEKSClusterRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesUnauthorized ", 401)
}

func (o *ListEKSClusterRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesUnauthorized ", 401)
}

func (o *ListEKSClusterRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSClusterRolesForbidden creates a ListEKSClusterRolesForbidden with default headers values
func NewListEKSClusterRolesForbidden() *ListEKSClusterRolesForbidden {
	return &ListEKSClusterRolesForbidden{}
}

/*
ListEKSClusterRolesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSClusterRolesForbidden struct {
}

// IsSuccess returns true when this list e k s cluster roles forbidden response has a 2xx status code
func (o *ListEKSClusterRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s cluster roles forbidden response has a 3xx status code
func (o *ListEKSClusterRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s cluster roles forbidden response has a 4xx status code
func (o *ListEKSClusterRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s cluster roles forbidden response has a 5xx status code
func (o *ListEKSClusterRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s cluster roles forbidden response a status code equal to that given
func (o *ListEKSClusterRolesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListEKSClusterRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesForbidden ", 403)
}

func (o *ListEKSClusterRolesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRolesForbidden ", 403)
}

func (o *ListEKSClusterRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSClusterRolesDefault creates a ListEKSClusterRolesDefault with default headers values
func NewListEKSClusterRolesDefault(code int) *ListEKSClusterRolesDefault {
	return &ListEKSClusterRolesDefault{
		_statusCode: code,
	}
}

/*
ListEKSClusterRolesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSClusterRolesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s cluster roles default response
func (o *ListEKSClusterRolesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s cluster roles default response has a 2xx status code
func (o *ListEKSClusterRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s cluster roles default response has a 3xx status code
func (o *ListEKSClusterRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s cluster roles default response has a 4xx status code
func (o *ListEKSClusterRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s cluster roles default response has a 5xx status code
func (o *ListEKSClusterRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s cluster roles default response a status code equal to that given
func (o *ListEKSClusterRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSClusterRolesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRoles default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSClusterRolesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/clusterroles][%d] listEKSClusterRoles default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSClusterRolesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSClusterRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
