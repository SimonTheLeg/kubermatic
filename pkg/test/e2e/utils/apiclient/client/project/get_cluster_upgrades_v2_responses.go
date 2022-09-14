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

// GetClusterUpgradesV2Reader is a Reader for the GetClusterUpgradesV2 structure.
type GetClusterUpgradesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterUpgradesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterUpgradesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterUpgradesV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterUpgradesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterUpgradesV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterUpgradesV2OK creates a GetClusterUpgradesV2OK with default headers values
func NewGetClusterUpgradesV2OK() *GetClusterUpgradesV2OK {
	return &GetClusterUpgradesV2OK{}
}

/* GetClusterUpgradesV2OK describes a response with status code 200, with default header values.

MasterVersion
*/
type GetClusterUpgradesV2OK struct {
	Payload []*models.MasterVersion
}

// IsSuccess returns true when this get cluster upgrades v2 o k response has a 2xx status code
func (o *GetClusterUpgradesV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster upgrades v2 o k response has a 3xx status code
func (o *GetClusterUpgradesV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster upgrades v2 o k response has a 4xx status code
func (o *GetClusterUpgradesV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster upgrades v2 o k response has a 5xx status code
func (o *GetClusterUpgradesV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster upgrades v2 o k response a status code equal to that given
func (o *GetClusterUpgradesV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterUpgradesV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2OK  %+v", 200, o.Payload)
}

func (o *GetClusterUpgradesV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2OK  %+v", 200, o.Payload)
}

func (o *GetClusterUpgradesV2OK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *GetClusterUpgradesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterUpgradesV2Unauthorized creates a GetClusterUpgradesV2Unauthorized with default headers values
func NewGetClusterUpgradesV2Unauthorized() *GetClusterUpgradesV2Unauthorized {
	return &GetClusterUpgradesV2Unauthorized{}
}

/* GetClusterUpgradesV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetClusterUpgradesV2Unauthorized struct {
}

// IsSuccess returns true when this get cluster upgrades v2 unauthorized response has a 2xx status code
func (o *GetClusterUpgradesV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster upgrades v2 unauthorized response has a 3xx status code
func (o *GetClusterUpgradesV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster upgrades v2 unauthorized response has a 4xx status code
func (o *GetClusterUpgradesV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster upgrades v2 unauthorized response has a 5xx status code
func (o *GetClusterUpgradesV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster upgrades v2 unauthorized response a status code equal to that given
func (o *GetClusterUpgradesV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterUpgradesV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2Unauthorized ", 401)
}

func (o *GetClusterUpgradesV2Unauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2Unauthorized ", 401)
}

func (o *GetClusterUpgradesV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterUpgradesV2Forbidden creates a GetClusterUpgradesV2Forbidden with default headers values
func NewGetClusterUpgradesV2Forbidden() *GetClusterUpgradesV2Forbidden {
	return &GetClusterUpgradesV2Forbidden{}
}

/* GetClusterUpgradesV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetClusterUpgradesV2Forbidden struct {
}

// IsSuccess returns true when this get cluster upgrades v2 forbidden response has a 2xx status code
func (o *GetClusterUpgradesV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster upgrades v2 forbidden response has a 3xx status code
func (o *GetClusterUpgradesV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster upgrades v2 forbidden response has a 4xx status code
func (o *GetClusterUpgradesV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster upgrades v2 forbidden response has a 5xx status code
func (o *GetClusterUpgradesV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster upgrades v2 forbidden response a status code equal to that given
func (o *GetClusterUpgradesV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterUpgradesV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2Forbidden ", 403)
}

func (o *GetClusterUpgradesV2Forbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2Forbidden ", 403)
}

func (o *GetClusterUpgradesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterUpgradesV2Default creates a GetClusterUpgradesV2Default with default headers values
func NewGetClusterUpgradesV2Default(code int) *GetClusterUpgradesV2Default {
	return &GetClusterUpgradesV2Default{
		_statusCode: code,
	}
}

/* GetClusterUpgradesV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type GetClusterUpgradesV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster upgrades v2 default response
func (o *GetClusterUpgradesV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster upgrades v2 default response has a 2xx status code
func (o *GetClusterUpgradesV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster upgrades v2 default response has a 3xx status code
func (o *GetClusterUpgradesV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster upgrades v2 default response has a 4xx status code
func (o *GetClusterUpgradesV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster upgrades v2 default response has a 5xx status code
func (o *GetClusterUpgradesV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster upgrades v2 default response a status code equal to that given
func (o *GetClusterUpgradesV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterUpgradesV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterUpgradesV2Default) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterUpgradesV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterUpgradesV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
