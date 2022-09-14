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

// ListEKSSecurityGroupsReader is a Reader for the ListEKSSecurityGroups structure.
type ListEKSSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSSecurityGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSSecurityGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSSecurityGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSSecurityGroupsOK creates a ListEKSSecurityGroupsOK with default headers values
func NewListEKSSecurityGroupsOK() *ListEKSSecurityGroupsOK {
	return &ListEKSSecurityGroupsOK{}
}

/* ListEKSSecurityGroupsOK describes a response with status code 200, with default header values.

EKSSecurityGroupList
*/
type ListEKSSecurityGroupsOK struct {
	Payload models.EKSSecurityGroupList
}

// IsSuccess returns true when this list e k s security groups o k response has a 2xx status code
func (o *ListEKSSecurityGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list e k s security groups o k response has a 3xx status code
func (o *ListEKSSecurityGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s security groups o k response has a 4xx status code
func (o *ListEKSSecurityGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list e k s security groups o k response has a 5xx status code
func (o *ListEKSSecurityGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s security groups o k response a status code equal to that given
func (o *ListEKSSecurityGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEKSSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *ListEKSSecurityGroupsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *ListEKSSecurityGroupsOK) GetPayload() models.EKSSecurityGroupList {
	return o.Payload
}

func (o *ListEKSSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSSecurityGroupsUnauthorized creates a ListEKSSecurityGroupsUnauthorized with default headers values
func NewListEKSSecurityGroupsUnauthorized() *ListEKSSecurityGroupsUnauthorized {
	return &ListEKSSecurityGroupsUnauthorized{}
}

/* ListEKSSecurityGroupsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSSecurityGroupsUnauthorized struct {
}

// IsSuccess returns true when this list e k s security groups unauthorized response has a 2xx status code
func (o *ListEKSSecurityGroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s security groups unauthorized response has a 3xx status code
func (o *ListEKSSecurityGroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s security groups unauthorized response has a 4xx status code
func (o *ListEKSSecurityGroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s security groups unauthorized response has a 5xx status code
func (o *ListEKSSecurityGroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s security groups unauthorized response a status code equal to that given
func (o *ListEKSSecurityGroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEKSSecurityGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroupsUnauthorized ", 401)
}

func (o *ListEKSSecurityGroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroupsUnauthorized ", 401)
}

func (o *ListEKSSecurityGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSSecurityGroupsForbidden creates a ListEKSSecurityGroupsForbidden with default headers values
func NewListEKSSecurityGroupsForbidden() *ListEKSSecurityGroupsForbidden {
	return &ListEKSSecurityGroupsForbidden{}
}

/* ListEKSSecurityGroupsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSSecurityGroupsForbidden struct {
}

// IsSuccess returns true when this list e k s security groups forbidden response has a 2xx status code
func (o *ListEKSSecurityGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list e k s security groups forbidden response has a 3xx status code
func (o *ListEKSSecurityGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list e k s security groups forbidden response has a 4xx status code
func (o *ListEKSSecurityGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list e k s security groups forbidden response has a 5xx status code
func (o *ListEKSSecurityGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list e k s security groups forbidden response a status code equal to that given
func (o *ListEKSSecurityGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListEKSSecurityGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroupsForbidden ", 403)
}

func (o *ListEKSSecurityGroupsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroupsForbidden ", 403)
}

func (o *ListEKSSecurityGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSSecurityGroupsDefault creates a ListEKSSecurityGroupsDefault with default headers values
func NewListEKSSecurityGroupsDefault(code int) *ListEKSSecurityGroupsDefault {
	return &ListEKSSecurityGroupsDefault{
		_statusCode: code,
	}
}

/* ListEKSSecurityGroupsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSSecurityGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s security groups default response
func (o *ListEKSSecurityGroupsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list e k s security groups default response has a 2xx status code
func (o *ListEKSSecurityGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list e k s security groups default response has a 3xx status code
func (o *ListEKSSecurityGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list e k s security groups default response has a 4xx status code
func (o *ListEKSSecurityGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list e k s security groups default response has a 5xx status code
func (o *ListEKSSecurityGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list e k s security groups default response a status code equal to that given
func (o *ListEKSSecurityGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListEKSSecurityGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSSecurityGroupsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/securitygroups][%d] listEKSSecurityGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListEKSSecurityGroupsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSSecurityGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
