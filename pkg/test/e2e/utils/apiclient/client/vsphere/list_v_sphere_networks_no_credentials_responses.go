// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListVSphereNetworksNoCredentialsReader is a Reader for the ListVSphereNetworksNoCredentials structure.
type ListVSphereNetworksNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVSphereNetworksNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVSphereNetworksNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVSphereNetworksNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVSphereNetworksNoCredentialsOK creates a ListVSphereNetworksNoCredentialsOK with default headers values
func NewListVSphereNetworksNoCredentialsOK() *ListVSphereNetworksNoCredentialsOK {
	return &ListVSphereNetworksNoCredentialsOK{}
}

/*
ListVSphereNetworksNoCredentialsOK describes a response with status code 200, with default header values.

VSphereNetwork
*/
type ListVSphereNetworksNoCredentialsOK struct {
	Payload []*models.VSphereNetwork
}

// IsSuccess returns true when this list v sphere networks no credentials o k response has a 2xx status code
func (o *ListVSphereNetworksNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list v sphere networks no credentials o k response has a 3xx status code
func (o *ListVSphereNetworksNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list v sphere networks no credentials o k response has a 4xx status code
func (o *ListVSphereNetworksNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list v sphere networks no credentials o k response has a 5xx status code
func (o *ListVSphereNetworksNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list v sphere networks no credentials o k response a status code equal to that given
func (o *ListVSphereNetworksNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListVSphereNetworksNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks][%d] listVSphereNetworksNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVSphereNetworksNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks][%d] listVSphereNetworksNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVSphereNetworksNoCredentialsOK) GetPayload() []*models.VSphereNetwork {
	return o.Payload
}

func (o *ListVSphereNetworksNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVSphereNetworksNoCredentialsDefault creates a ListVSphereNetworksNoCredentialsDefault with default headers values
func NewListVSphereNetworksNoCredentialsDefault(code int) *ListVSphereNetworksNoCredentialsDefault {
	return &ListVSphereNetworksNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListVSphereNetworksNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVSphereNetworksNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v sphere networks no credentials default response
func (o *ListVSphereNetworksNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list v sphere networks no credentials default response has a 2xx status code
func (o *ListVSphereNetworksNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list v sphere networks no credentials default response has a 3xx status code
func (o *ListVSphereNetworksNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list v sphere networks no credentials default response has a 4xx status code
func (o *ListVSphereNetworksNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list v sphere networks no credentials default response has a 5xx status code
func (o *ListVSphereNetworksNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list v sphere networks no credentials default response a status code equal to that given
func (o *ListVSphereNetworksNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListVSphereNetworksNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks][%d] listVSphereNetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVSphereNetworksNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks][%d] listVSphereNetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVSphereNetworksNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVSphereNetworksNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
