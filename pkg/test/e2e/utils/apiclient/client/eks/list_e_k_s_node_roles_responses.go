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

// ListEKSNodeRolesReader is a Reader for the ListEKSNodeRoles structure.
type ListEKSNodeRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSNodeRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSNodeRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSNodeRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSNodeRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSNodeRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSNodeRolesOK creates a ListEKSNodeRolesOK with default headers values
func NewListEKSNodeRolesOK() *ListEKSNodeRolesOK {
	return &ListEKSNodeRolesOK{}
}

/* ListEKSNodeRolesOK describes a response with status code 200, with default header values.

EKSNodeRoleList
*/
type ListEKSNodeRolesOK struct {
	Payload models.EKSNodeRoleList
}

// IsSuccess returns true when this list e k s node roles o k response has a 2xx status code
func (o *ListEKSNodeRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s node roles o k response has a 3xx status code
func (o *ListEKSNodeRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s node roles o k response has a 4xx status code
func (o *ListEKSNodeRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s node roles o k response has a 5xx status code
func (o *ListEKSNodeRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s node roles o k response a status code equal to that given
func (o *ListEKSNodeRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSNodeRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRolesOK  %+v", 200, o.Payload)
}

func (o *ListEKSNodeRolesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRolesOK  %+v", 200, o.Payload)
}

func (o *ListEKSNodeRolesOK) GetPayload() models.EKSNodeRoleList {
	return o.Payload
}

func (o *ListEKSNodeRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSNodeRolesUnauthorized creates a ListEKSNodeRolesUnauthorized with default headers values
func NewListEKSNodeRolesUnauthorized() *ListEKSNodeRolesUnauthorized {
	return &ListEKSNodeRolesUnauthorized{}
}

/* ListEKSNodeRolesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSNodeRolesUnauthorized struct {
}

// IsSuccess returns true when this list e k s node roles unauthorized response has a 2xx status code
func (o *ListEKSNodeRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s node roles unauthorized response has a 3xx status code
func (o *ListEKSNodeRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s node roles unauthorized response has a 4xx status code
func (o *ListEKSNodeRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s node roles unauthorized response has a 5xx status code
func (o *ListEKSNodeRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s node roles unauthorized response a status code equal to that given
func (o *ListEKSNodeRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEKSNodeRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRolesUnauthorized ", 401)
}

func (o *ListEKSNodeRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRolesUnauthorized ", 401)
}

func (o *ListEKSNodeRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSNodeRolesForbidden creates a ListEKSNodeRolesForbidden with default headers values
func NewListEKSNodeRolesForbidden() *ListEKSNodeRolesForbidden {
	return &ListEKSNodeRolesForbidden{}
}

/* ListEKSNodeRolesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSNodeRolesForbidden struct {
}

// IsSuccess returns true when this list e k s node roles forbidden response has a 2xx status code
func (o *ListEKSNodeRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s node roles forbidden response has a 3xx status code
func (o *ListEKSNodeRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s node roles forbidden response has a 4xx status code
func (o *ListEKSNodeRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s node roles forbidden response has a 5xx status code
func (o *ListEKSNodeRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s node roles forbidden response a status code equal to that given
func (o *ListEKSNodeRolesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListEKSNodeRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRolesForbidden ", 403)
}

func (o *ListEKSNodeRolesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRolesForbidden ", 403)
}

func (o *ListEKSNodeRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSNodeRolesDefault creates a ListEKSNodeRolesDefault with default headers values
func NewListEKSNodeRolesDefault(code int) *ListEKSNodeRolesDefault {
	return &ListEKSNodeRolesDefault{
		_statusCode: code,
	}
}

/* ListEKSNodeRolesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSNodeRolesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s node roles default response
func (o *ListEKSNodeRolesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s node roles default response has a 2xx status code
func (o *ListEKSNodeRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s node roles default response has a 3xx status code
func (o *ListEKSNodeRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s node roles default response has a 4xx status code
func (o *ListEKSNodeRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s node roles default response has a 5xx status code
func (o *ListEKSNodeRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s node roles default response a status code equal to that given
func (o *ListEKSNodeRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSNodeRolesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRoles default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSNodeRolesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/noderoles][%d] listEKSNodeRoles default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSNodeRolesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSNodeRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
