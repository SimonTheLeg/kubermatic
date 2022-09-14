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

// GetMachineDeploymentReader is a Reader for the GetMachineDeployment structure.
type GetMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMachineDeploymentOK creates a GetMachineDeploymentOK with default headers values
func NewGetMachineDeploymentOK() *GetMachineDeploymentOK {
	return &GetMachineDeploymentOK{}
}

/* GetMachineDeploymentOK describes a response with status code 200, with default header values.

NodeDeployment
*/
type GetMachineDeploymentOK struct {
	Payload *models.NodeDeployment
}

// IsSuccess returns true when this get machine deployment o k response has a 2xx status code
func (o *GetMachineDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get machine deployment o k response has a 3xx status code
func (o *GetMachineDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine deployment o k response has a 4xx status code
func (o *GetMachineDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get machine deployment o k response has a 5xx status code
func (o *GetMachineDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine deployment o k response a status code equal to that given
func (o *GetMachineDeploymentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMachineDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetMachineDeploymentOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetMachineDeploymentOK) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *GetMachineDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineDeploymentUnauthorized creates a GetMachineDeploymentUnauthorized with default headers values
func NewGetMachineDeploymentUnauthorized() *GetMachineDeploymentUnauthorized {
	return &GetMachineDeploymentUnauthorized{}
}

/* GetMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetMachineDeploymentUnauthorized struct {
}

// IsSuccess returns true when this get machine deployment unauthorized response has a 2xx status code
func (o *GetMachineDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get machine deployment unauthorized response has a 3xx status code
func (o *GetMachineDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine deployment unauthorized response has a 4xx status code
func (o *GetMachineDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get machine deployment unauthorized response has a 5xx status code
func (o *GetMachineDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine deployment unauthorized response a status code equal to that given
func (o *GetMachineDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeploymentUnauthorized ", 401)
}

func (o *GetMachineDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeploymentUnauthorized ", 401)
}

func (o *GetMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMachineDeploymentForbidden creates a GetMachineDeploymentForbidden with default headers values
func NewGetMachineDeploymentForbidden() *GetMachineDeploymentForbidden {
	return &GetMachineDeploymentForbidden{}
}

/* GetMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetMachineDeploymentForbidden struct {
}

// IsSuccess returns true when this get machine deployment forbidden response has a 2xx status code
func (o *GetMachineDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get machine deployment forbidden response has a 3xx status code
func (o *GetMachineDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine deployment forbidden response has a 4xx status code
func (o *GetMachineDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get machine deployment forbidden response has a 5xx status code
func (o *GetMachineDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine deployment forbidden response a status code equal to that given
func (o *GetMachineDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeploymentForbidden ", 403)
}

func (o *GetMachineDeploymentForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeploymentForbidden ", 403)
}

func (o *GetMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMachineDeploymentDefault creates a GetMachineDeploymentDefault with default headers values
func NewGetMachineDeploymentDefault(code int) *GetMachineDeploymentDefault {
	return &GetMachineDeploymentDefault{
		_statusCode: code,
	}
}

/* GetMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get machine deployment default response
func (o *GetMachineDeploymentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get machine deployment default response has a 2xx status code
func (o *GetMachineDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get machine deployment default response has a 3xx status code
func (o *GetMachineDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get machine deployment default response has a 4xx status code
func (o *GetMachineDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get machine deployment default response has a 5xx status code
func (o *GetMachineDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get machine deployment default response a status code equal to that given
func (o *GetMachineDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetMachineDeploymentDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *GetMachineDeploymentDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] getMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *GetMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
