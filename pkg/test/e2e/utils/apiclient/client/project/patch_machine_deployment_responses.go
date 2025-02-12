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

// PatchMachineDeploymentReader is a Reader for the PatchMachineDeployment structure.
type PatchMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMachineDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchMachineDeploymentOK creates a PatchMachineDeploymentOK with default headers values
func NewPatchMachineDeploymentOK() *PatchMachineDeploymentOK {
	return &PatchMachineDeploymentOK{}
}

/*
PatchMachineDeploymentOK describes a response with status code 200, with default header values.

NodeDeployment
*/
type PatchMachineDeploymentOK struct {
	Payload *models.NodeDeployment
}

// IsSuccess returns true when this patch machine deployment o k response has a 2xx status code
func (o *PatchMachineDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch machine deployment o k response has a 3xx status code
func (o *PatchMachineDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch machine deployment o k response has a 4xx status code
func (o *PatchMachineDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch machine deployment o k response has a 5xx status code
func (o *PatchMachineDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch machine deployment o k response a status code equal to that given
func (o *PatchMachineDeploymentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchMachineDeploymentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeploymentOK  %+v", 200, o.Payload)
}

func (o *PatchMachineDeploymentOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeploymentOK  %+v", 200, o.Payload)
}

func (o *PatchMachineDeploymentOK) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *PatchMachineDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMachineDeploymentUnauthorized creates a PatchMachineDeploymentUnauthorized with default headers values
func NewPatchMachineDeploymentUnauthorized() *PatchMachineDeploymentUnauthorized {
	return &PatchMachineDeploymentUnauthorized{}
}

/*
PatchMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchMachineDeploymentUnauthorized struct {
}

// IsSuccess returns true when this patch machine deployment unauthorized response has a 2xx status code
func (o *PatchMachineDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch machine deployment unauthorized response has a 3xx status code
func (o *PatchMachineDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch machine deployment unauthorized response has a 4xx status code
func (o *PatchMachineDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch machine deployment unauthorized response has a 5xx status code
func (o *PatchMachineDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch machine deployment unauthorized response a status code equal to that given
func (o *PatchMachineDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeploymentUnauthorized ", 401)
}

func (o *PatchMachineDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeploymentUnauthorized ", 401)
}

func (o *PatchMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMachineDeploymentForbidden creates a PatchMachineDeploymentForbidden with default headers values
func NewPatchMachineDeploymentForbidden() *PatchMachineDeploymentForbidden {
	return &PatchMachineDeploymentForbidden{}
}

/*
PatchMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchMachineDeploymentForbidden struct {
}

// IsSuccess returns true when this patch machine deployment forbidden response has a 2xx status code
func (o *PatchMachineDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch machine deployment forbidden response has a 3xx status code
func (o *PatchMachineDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch machine deployment forbidden response has a 4xx status code
func (o *PatchMachineDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch machine deployment forbidden response has a 5xx status code
func (o *PatchMachineDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch machine deployment forbidden response a status code equal to that given
func (o *PatchMachineDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeploymentForbidden ", 403)
}

func (o *PatchMachineDeploymentForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeploymentForbidden ", 403)
}

func (o *PatchMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMachineDeploymentDefault creates a PatchMachineDeploymentDefault with default headers values
func NewPatchMachineDeploymentDefault(code int) *PatchMachineDeploymentDefault {
	return &PatchMachineDeploymentDefault{
		_statusCode: code,
	}
}

/*
PatchMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch machine deployment default response
func (o *PatchMachineDeploymentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch machine deployment default response has a 2xx status code
func (o *PatchMachineDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch machine deployment default response has a 3xx status code
func (o *PatchMachineDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch machine deployment default response has a 4xx status code
func (o *PatchMachineDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch machine deployment default response has a 5xx status code
func (o *PatchMachineDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch machine deployment default response a status code equal to that given
func (o *PatchMachineDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchMachineDeploymentDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *PatchMachineDeploymentDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] patchMachineDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *PatchMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
