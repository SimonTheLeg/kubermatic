// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// MigrateClusterToExternalCCMReader is a Reader for the MigrateClusterToExternalCCM structure.
type MigrateClusterToExternalCCMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateClusterToExternalCCMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateClusterToExternalCCMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewMigrateClusterToExternalCCMUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMigrateClusterToExternalCCMForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewMigrateClusterToExternalCCMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMigrateClusterToExternalCCMOK creates a MigrateClusterToExternalCCMOK with default headers values
func NewMigrateClusterToExternalCCMOK() *MigrateClusterToExternalCCMOK {
	return &MigrateClusterToExternalCCMOK{}
}

/* MigrateClusterToExternalCCMOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type MigrateClusterToExternalCCMOK struct {
}

// IsSuccess returns true when this migrate cluster to external c c m o k response has a 2xx status code
func (o *MigrateClusterToExternalCCMOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this migrate cluster to external c c m o k response has a 3xx status code
func (o *MigrateClusterToExternalCCMOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate cluster to external c c m o k response has a 4xx status code
func (o *MigrateClusterToExternalCCMOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this migrate cluster to external c c m o k response has a 5xx status code
func (o *MigrateClusterToExternalCCMOK) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate cluster to external c c m o k response a status code equal to that given
func (o *MigrateClusterToExternalCCMOK) IsCode(code int) bool {
	return code == 200
}

func (o *MigrateClusterToExternalCCMOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMOK ", 200)
}

func (o *MigrateClusterToExternalCCMOK) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMOK ", 200)
}

func (o *MigrateClusterToExternalCCMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMigrateClusterToExternalCCMUnauthorized creates a MigrateClusterToExternalCCMUnauthorized with default headers values
func NewMigrateClusterToExternalCCMUnauthorized() *MigrateClusterToExternalCCMUnauthorized {
	return &MigrateClusterToExternalCCMUnauthorized{}
}

/* MigrateClusterToExternalCCMUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type MigrateClusterToExternalCCMUnauthorized struct {
}

// IsSuccess returns true when this migrate cluster to external c c m unauthorized response has a 2xx status code
func (o *MigrateClusterToExternalCCMUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this migrate cluster to external c c m unauthorized response has a 3xx status code
func (o *MigrateClusterToExternalCCMUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate cluster to external c c m unauthorized response has a 4xx status code
func (o *MigrateClusterToExternalCCMUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this migrate cluster to external c c m unauthorized response has a 5xx status code
func (o *MigrateClusterToExternalCCMUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate cluster to external c c m unauthorized response a status code equal to that given
func (o *MigrateClusterToExternalCCMUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *MigrateClusterToExternalCCMUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMUnauthorized ", 401)
}

func (o *MigrateClusterToExternalCCMUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMUnauthorized ", 401)
}

func (o *MigrateClusterToExternalCCMUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMigrateClusterToExternalCCMForbidden creates a MigrateClusterToExternalCCMForbidden with default headers values
func NewMigrateClusterToExternalCCMForbidden() *MigrateClusterToExternalCCMForbidden {
	return &MigrateClusterToExternalCCMForbidden{}
}

/* MigrateClusterToExternalCCMForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type MigrateClusterToExternalCCMForbidden struct {
}

// IsSuccess returns true when this migrate cluster to external c c m forbidden response has a 2xx status code
func (o *MigrateClusterToExternalCCMForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this migrate cluster to external c c m forbidden response has a 3xx status code
func (o *MigrateClusterToExternalCCMForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate cluster to external c c m forbidden response has a 4xx status code
func (o *MigrateClusterToExternalCCMForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this migrate cluster to external c c m forbidden response has a 5xx status code
func (o *MigrateClusterToExternalCCMForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate cluster to external c c m forbidden response a status code equal to that given
func (o *MigrateClusterToExternalCCMForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *MigrateClusterToExternalCCMForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMForbidden ", 403)
}

func (o *MigrateClusterToExternalCCMForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMForbidden ", 403)
}

func (o *MigrateClusterToExternalCCMForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMigrateClusterToExternalCCMDefault creates a MigrateClusterToExternalCCMDefault with default headers values
func NewMigrateClusterToExternalCCMDefault(code int) *MigrateClusterToExternalCCMDefault {
	return &MigrateClusterToExternalCCMDefault{
		_statusCode: code,
	}
}

/* MigrateClusterToExternalCCMDefault describes a response with status code -1, with default header values.

errorResponse
*/
type MigrateClusterToExternalCCMDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the migrate cluster to external c c m default response
func (o *MigrateClusterToExternalCCMDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this migrate cluster to external c c m default response has a 2xx status code
func (o *MigrateClusterToExternalCCMDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this migrate cluster to external c c m default response has a 3xx status code
func (o *MigrateClusterToExternalCCMDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this migrate cluster to external c c m default response has a 4xx status code
func (o *MigrateClusterToExternalCCMDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this migrate cluster to external c c m default response has a 5xx status code
func (o *MigrateClusterToExternalCCMDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this migrate cluster to external c c m default response a status code equal to that given
func (o *MigrateClusterToExternalCCMDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *MigrateClusterToExternalCCMDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCM default  %+v", o._statusCode, o.Payload)
}

func (o *MigrateClusterToExternalCCMDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCM default  %+v", o._statusCode, o.Payload)
}

func (o *MigrateClusterToExternalCCMDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MigrateClusterToExternalCCMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
