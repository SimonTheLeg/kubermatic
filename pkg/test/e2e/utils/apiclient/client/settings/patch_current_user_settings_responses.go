// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// PatchCurrentUserSettingsReader is a Reader for the PatchCurrentUserSettings structure.
type PatchCurrentUserSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCurrentUserSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCurrentUserSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCurrentUserSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchCurrentUserSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCurrentUserSettingsOK creates a PatchCurrentUserSettingsOK with default headers values
func NewPatchCurrentUserSettingsOK() *PatchCurrentUserSettingsOK {
	return &PatchCurrentUserSettingsOK{}
}

/* PatchCurrentUserSettingsOK describes a response with status code 200, with default header values.

UserSettings
*/
type PatchCurrentUserSettingsOK struct {
	Payload *models.UserSettings
}

// IsSuccess returns true when this patch current user settings o k response has a 2xx status code
func (o *PatchCurrentUserSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch current user settings o k response has a 3xx status code
func (o *PatchCurrentUserSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch current user settings o k response has a 4xx status code
func (o *PatchCurrentUserSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch current user settings o k response has a 5xx status code
func (o *PatchCurrentUserSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch current user settings o k response a status code equal to that given
func (o *PatchCurrentUserSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchCurrentUserSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/me/settings][%d] patchCurrentUserSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchCurrentUserSettingsOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/me/settings][%d] patchCurrentUserSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchCurrentUserSettingsOK) GetPayload() *models.UserSettings {
	return o.Payload
}

func (o *PatchCurrentUserSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCurrentUserSettingsUnauthorized creates a PatchCurrentUserSettingsUnauthorized with default headers values
func NewPatchCurrentUserSettingsUnauthorized() *PatchCurrentUserSettingsUnauthorized {
	return &PatchCurrentUserSettingsUnauthorized{}
}

/* PatchCurrentUserSettingsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchCurrentUserSettingsUnauthorized struct {
}

// IsSuccess returns true when this patch current user settings unauthorized response has a 2xx status code
func (o *PatchCurrentUserSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch current user settings unauthorized response has a 3xx status code
func (o *PatchCurrentUserSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch current user settings unauthorized response has a 4xx status code
func (o *PatchCurrentUserSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch current user settings unauthorized response has a 5xx status code
func (o *PatchCurrentUserSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch current user settings unauthorized response a status code equal to that given
func (o *PatchCurrentUserSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchCurrentUserSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/me/settings][%d] patchCurrentUserSettingsUnauthorized ", 401)
}

func (o *PatchCurrentUserSettingsUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/me/settings][%d] patchCurrentUserSettingsUnauthorized ", 401)
}

func (o *PatchCurrentUserSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCurrentUserSettingsDefault creates a PatchCurrentUserSettingsDefault with default headers values
func NewPatchCurrentUserSettingsDefault(code int) *PatchCurrentUserSettingsDefault {
	return &PatchCurrentUserSettingsDefault{
		_statusCode: code,
	}
}

/* PatchCurrentUserSettingsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchCurrentUserSettingsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch current user settings default response
func (o *PatchCurrentUserSettingsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch current user settings default response has a 2xx status code
func (o *PatchCurrentUserSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch current user settings default response has a 3xx status code
func (o *PatchCurrentUserSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch current user settings default response has a 4xx status code
func (o *PatchCurrentUserSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch current user settings default response has a 5xx status code
func (o *PatchCurrentUserSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch current user settings default response a status code equal to that given
func (o *PatchCurrentUserSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchCurrentUserSettingsDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/me/settings][%d] patchCurrentUserSettings default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCurrentUserSettingsDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/me/settings][%d] patchCurrentUserSettings default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCurrentUserSettingsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchCurrentUserSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
