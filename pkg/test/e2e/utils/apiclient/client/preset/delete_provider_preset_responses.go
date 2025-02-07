// Code generated by go-swagger; DO NOT EDIT.

package preset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteProviderPresetReader is a Reader for the DeleteProviderPreset structure.
type DeleteProviderPresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProviderPresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProviderPresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteProviderPresetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProviderPresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteProviderPresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProviderPresetOK creates a DeleteProviderPresetOK with default headers values
func NewDeleteProviderPresetOK() *DeleteProviderPresetOK {
	return &DeleteProviderPresetOK{}
}

/*
DeleteProviderPresetOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteProviderPresetOK struct {
}

// IsSuccess returns true when this delete provider preset o k response has a 2xx status code
func (o *DeleteProviderPresetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete provider preset o k response has a 3xx status code
func (o *DeleteProviderPresetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete provider preset o k response has a 4xx status code
func (o *DeleteProviderPresetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete provider preset o k response has a 5xx status code
func (o *DeleteProviderPresetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete provider preset o k response a status code equal to that given
func (o *DeleteProviderPresetOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteProviderPresetOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPresetOK ", 200)
}

func (o *DeleteProviderPresetOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPresetOK ", 200)
}

func (o *DeleteProviderPresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProviderPresetUnauthorized creates a DeleteProviderPresetUnauthorized with default headers values
func NewDeleteProviderPresetUnauthorized() *DeleteProviderPresetUnauthorized {
	return &DeleteProviderPresetUnauthorized{}
}

/*
DeleteProviderPresetUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteProviderPresetUnauthorized struct {
}

// IsSuccess returns true when this delete provider preset unauthorized response has a 2xx status code
func (o *DeleteProviderPresetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete provider preset unauthorized response has a 3xx status code
func (o *DeleteProviderPresetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete provider preset unauthorized response has a 4xx status code
func (o *DeleteProviderPresetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete provider preset unauthorized response has a 5xx status code
func (o *DeleteProviderPresetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete provider preset unauthorized response a status code equal to that given
func (o *DeleteProviderPresetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteProviderPresetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPresetUnauthorized ", 401)
}

func (o *DeleteProviderPresetUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPresetUnauthorized ", 401)
}

func (o *DeleteProviderPresetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProviderPresetForbidden creates a DeleteProviderPresetForbidden with default headers values
func NewDeleteProviderPresetForbidden() *DeleteProviderPresetForbidden {
	return &DeleteProviderPresetForbidden{}
}

/*
DeleteProviderPresetForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteProviderPresetForbidden struct {
}

// IsSuccess returns true when this delete provider preset forbidden response has a 2xx status code
func (o *DeleteProviderPresetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete provider preset forbidden response has a 3xx status code
func (o *DeleteProviderPresetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete provider preset forbidden response has a 4xx status code
func (o *DeleteProviderPresetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete provider preset forbidden response has a 5xx status code
func (o *DeleteProviderPresetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete provider preset forbidden response a status code equal to that given
func (o *DeleteProviderPresetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteProviderPresetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPresetForbidden ", 403)
}

func (o *DeleteProviderPresetForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPresetForbidden ", 403)
}

func (o *DeleteProviderPresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProviderPresetDefault creates a DeleteProviderPresetDefault with default headers values
func NewDeleteProviderPresetDefault(code int) *DeleteProviderPresetDefault {
	return &DeleteProviderPresetDefault{
		_statusCode: code,
	}
}

/*
DeleteProviderPresetDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteProviderPresetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete provider preset default response
func (o *DeleteProviderPresetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete provider preset default response has a 2xx status code
func (o *DeleteProviderPresetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete provider preset default response has a 3xx status code
func (o *DeleteProviderPresetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete provider preset default response has a 4xx status code
func (o *DeleteProviderPresetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete provider preset default response has a 5xx status code
func (o *DeleteProviderPresetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete provider preset default response a status code equal to that given
func (o *DeleteProviderPresetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteProviderPresetDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPreset default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProviderPresetDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/providers/{provider_name}/presets/{preset_name}][%d] deleteProviderPreset default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProviderPresetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProviderPresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
