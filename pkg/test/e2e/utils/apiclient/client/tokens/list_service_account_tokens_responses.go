// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListServiceAccountTokensReader is a Reader for the ListServiceAccountTokens structure.
type ListServiceAccountTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceAccountTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceAccountTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServiceAccountTokensUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListServiceAccountTokensForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListServiceAccountTokensDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServiceAccountTokensOK creates a ListServiceAccountTokensOK with default headers values
func NewListServiceAccountTokensOK() *ListServiceAccountTokensOK {
	return &ListServiceAccountTokensOK{}
}

/*
ListServiceAccountTokensOK describes a response with status code 200, with default header values.

PublicServiceAccountToken
*/
type ListServiceAccountTokensOK struct {
	Payload []*models.PublicServiceAccountToken
}

// IsSuccess returns true when this list service account tokens o k response has a 2xx status code
func (o *ListServiceAccountTokensOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list service account tokens o k response has a 3xx status code
func (o *ListServiceAccountTokensOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service account tokens o k response has a 4xx status code
func (o *ListServiceAccountTokensOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list service account tokens o k response has a 5xx status code
func (o *ListServiceAccountTokensOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list service account tokens o k response a status code equal to that given
func (o *ListServiceAccountTokensOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListServiceAccountTokensOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokensOK  %+v", 200, o.Payload)
}

func (o *ListServiceAccountTokensOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokensOK  %+v", 200, o.Payload)
}

func (o *ListServiceAccountTokensOK) GetPayload() []*models.PublicServiceAccountToken {
	return o.Payload
}

func (o *ListServiceAccountTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceAccountTokensUnauthorized creates a ListServiceAccountTokensUnauthorized with default headers values
func NewListServiceAccountTokensUnauthorized() *ListServiceAccountTokensUnauthorized {
	return &ListServiceAccountTokensUnauthorized{}
}

/*
ListServiceAccountTokensUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListServiceAccountTokensUnauthorized struct {
}

// IsSuccess returns true when this list service account tokens unauthorized response has a 2xx status code
func (o *ListServiceAccountTokensUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list service account tokens unauthorized response has a 3xx status code
func (o *ListServiceAccountTokensUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service account tokens unauthorized response has a 4xx status code
func (o *ListServiceAccountTokensUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list service account tokens unauthorized response has a 5xx status code
func (o *ListServiceAccountTokensUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list service account tokens unauthorized response a status code equal to that given
func (o *ListServiceAccountTokensUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListServiceAccountTokensUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokensUnauthorized ", 401)
}

func (o *ListServiceAccountTokensUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokensUnauthorized ", 401)
}

func (o *ListServiceAccountTokensUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListServiceAccountTokensForbidden creates a ListServiceAccountTokensForbidden with default headers values
func NewListServiceAccountTokensForbidden() *ListServiceAccountTokensForbidden {
	return &ListServiceAccountTokensForbidden{}
}

/*
ListServiceAccountTokensForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListServiceAccountTokensForbidden struct {
}

// IsSuccess returns true when this list service account tokens forbidden response has a 2xx status code
func (o *ListServiceAccountTokensForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list service account tokens forbidden response has a 3xx status code
func (o *ListServiceAccountTokensForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list service account tokens forbidden response has a 4xx status code
func (o *ListServiceAccountTokensForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list service account tokens forbidden response has a 5xx status code
func (o *ListServiceAccountTokensForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list service account tokens forbidden response a status code equal to that given
func (o *ListServiceAccountTokensForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListServiceAccountTokensForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokensForbidden ", 403)
}

func (o *ListServiceAccountTokensForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokensForbidden ", 403)
}

func (o *ListServiceAccountTokensForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListServiceAccountTokensDefault creates a ListServiceAccountTokensDefault with default headers values
func NewListServiceAccountTokensDefault(code int) *ListServiceAccountTokensDefault {
	return &ListServiceAccountTokensDefault{
		_statusCode: code,
	}
}

/*
ListServiceAccountTokensDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListServiceAccountTokensDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list service account tokens default response
func (o *ListServiceAccountTokensDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list service account tokens default response has a 2xx status code
func (o *ListServiceAccountTokensDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list service account tokens default response has a 3xx status code
func (o *ListServiceAccountTokensDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list service account tokens default response has a 4xx status code
func (o *ListServiceAccountTokensDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list service account tokens default response has a 5xx status code
func (o *ListServiceAccountTokensDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list service account tokens default response a status code equal to that given
func (o *ListServiceAccountTokensDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListServiceAccountTokensDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokens default  %+v", o._statusCode, o.Payload)
}

func (o *ListServiceAccountTokensDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens][%d] listServiceAccountTokens default  %+v", o._statusCode, o.Payload)
}

func (o *ListServiceAccountTokensDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListServiceAccountTokensDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
