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

// ListExternalClusterMachineDeploymentsReader is a Reader for the ListExternalClusterMachineDeployments structure.
type ListExternalClusterMachineDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalClusterMachineDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalClusterMachineDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListExternalClusterMachineDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExternalClusterMachineDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListExternalClusterMachineDeploymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListExternalClusterMachineDeploymentsOK creates a ListExternalClusterMachineDeploymentsOK with default headers values
func NewListExternalClusterMachineDeploymentsOK() *ListExternalClusterMachineDeploymentsOK {
	return &ListExternalClusterMachineDeploymentsOK{}
}

/* ListExternalClusterMachineDeploymentsOK describes a response with status code 200, with default header values.

ExternalClusterMachineDeployment
*/
type ListExternalClusterMachineDeploymentsOK struct {
	Payload []*models.ExternalClusterMachineDeployment
}

// IsSuccess returns true when this list external cluster machine deployments o k response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list external cluster machine deployments o k response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployments o k response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list external cluster machine deployments o k response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployments o k response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListExternalClusterMachineDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeploymentsOK  %+v", 200, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeploymentsOK  %+v", 200, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentsOK) GetPayload() []*models.ExternalClusterMachineDeployment {
	return o.Payload
}

func (o *ListExternalClusterMachineDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalClusterMachineDeploymentsUnauthorized creates a ListExternalClusterMachineDeploymentsUnauthorized with default headers values
func NewListExternalClusterMachineDeploymentsUnauthorized() *ListExternalClusterMachineDeploymentsUnauthorized {
	return &ListExternalClusterMachineDeploymentsUnauthorized{}
}

/* ListExternalClusterMachineDeploymentsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterMachineDeploymentsUnauthorized struct {
}

// IsSuccess returns true when this list external cluster machine deployments unauthorized response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster machine deployments unauthorized response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployments unauthorized response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster machine deployments unauthorized response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployments unauthorized response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListExternalClusterMachineDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeploymentsUnauthorized ", 401)
}

func (o *ListExternalClusterMachineDeploymentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeploymentsUnauthorized ", 401)
}

func (o *ListExternalClusterMachineDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterMachineDeploymentsForbidden creates a ListExternalClusterMachineDeploymentsForbidden with default headers values
func NewListExternalClusterMachineDeploymentsForbidden() *ListExternalClusterMachineDeploymentsForbidden {
	return &ListExternalClusterMachineDeploymentsForbidden{}
}

/* ListExternalClusterMachineDeploymentsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterMachineDeploymentsForbidden struct {
}

// IsSuccess returns true when this list external cluster machine deployments forbidden response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list external cluster machine deployments forbidden response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list external cluster machine deployments forbidden response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list external cluster machine deployments forbidden response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list external cluster machine deployments forbidden response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListExternalClusterMachineDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeploymentsForbidden ", 403)
}

func (o *ListExternalClusterMachineDeploymentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeploymentsForbidden ", 403)
}

func (o *ListExternalClusterMachineDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterMachineDeploymentsDefault creates a ListExternalClusterMachineDeploymentsDefault with default headers values
func NewListExternalClusterMachineDeploymentsDefault(code int) *ListExternalClusterMachineDeploymentsDefault {
	return &ListExternalClusterMachineDeploymentsDefault{
		_statusCode: code,
	}
}

/* ListExternalClusterMachineDeploymentsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListExternalClusterMachineDeploymentsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list external cluster machine deployments default response
func (o *ListExternalClusterMachineDeploymentsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list external cluster machine deployments default response has a 2xx status code
func (o *ListExternalClusterMachineDeploymentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list external cluster machine deployments default response has a 3xx status code
func (o *ListExternalClusterMachineDeploymentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list external cluster machine deployments default response has a 4xx status code
func (o *ListExternalClusterMachineDeploymentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list external cluster machine deployments default response has a 5xx status code
func (o *ListExternalClusterMachineDeploymentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list external cluster machine deployments default response a status code equal to that given
func (o *ListExternalClusterMachineDeploymentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListExternalClusterMachineDeploymentsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeployments default  %+v", o._statusCode, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/machinedeployments][%d] listExternalClusterMachineDeployments default  %+v", o._statusCode, o.Payload)
}

func (o *ListExternalClusterMachineDeploymentsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListExternalClusterMachineDeploymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
