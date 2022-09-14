// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteAddonV2Reader is a Reader for the DeleteAddonV2 structure.
type DeleteAddonV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAddonV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAddonV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAddonV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAddonV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAddonV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAddonV2OK creates a DeleteAddonV2OK with default headers values
func NewDeleteAddonV2OK() *DeleteAddonV2OK {
	return &DeleteAddonV2OK{}
}

/* DeleteAddonV2OK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteAddonV2OK struct {
}

// IsSuccess returns true when this delete addon v2 o k response has a 2xx status code
func (o *DeleteAddonV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete addon v2 o k response has a 3xx status code
func (o *DeleteAddonV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete addon v2 o k response has a 4xx status code
func (o *DeleteAddonV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete addon v2 o k response has a 5xx status code
func (o *DeleteAddonV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete addon v2 o k response a status code equal to that given
func (o *DeleteAddonV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAddonV2OK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2OK ", 200)
}

func (o *DeleteAddonV2OK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2OK ", 200)
}

func (o *DeleteAddonV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAddonV2Unauthorized creates a DeleteAddonV2Unauthorized with default headers values
func NewDeleteAddonV2Unauthorized() *DeleteAddonV2Unauthorized {
	return &DeleteAddonV2Unauthorized{}
}

/* DeleteAddonV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteAddonV2Unauthorized struct {
}

// IsSuccess returns true when this delete addon v2 unauthorized response has a 2xx status code
func (o *DeleteAddonV2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete addon v2 unauthorized response has a 3xx status code
func (o *DeleteAddonV2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete addon v2 unauthorized response has a 4xx status code
func (o *DeleteAddonV2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete addon v2 unauthorized response has a 5xx status code
func (o *DeleteAddonV2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete addon v2 unauthorized response a status code equal to that given
func (o *DeleteAddonV2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteAddonV2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2Unauthorized ", 401)
}

func (o *DeleteAddonV2Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2Unauthorized ", 401)
}

func (o *DeleteAddonV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAddonV2Forbidden creates a DeleteAddonV2Forbidden with default headers values
func NewDeleteAddonV2Forbidden() *DeleteAddonV2Forbidden {
	return &DeleteAddonV2Forbidden{}
}

/* DeleteAddonV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteAddonV2Forbidden struct {
}

// IsSuccess returns true when this delete addon v2 forbidden response has a 2xx status code
func (o *DeleteAddonV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete addon v2 forbidden response has a 3xx status code
func (o *DeleteAddonV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete addon v2 forbidden response has a 4xx status code
func (o *DeleteAddonV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete addon v2 forbidden response has a 5xx status code
func (o *DeleteAddonV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete addon v2 forbidden response a status code equal to that given
func (o *DeleteAddonV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteAddonV2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2Forbidden ", 403)
}

func (o *DeleteAddonV2Forbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2Forbidden ", 403)
}

func (o *DeleteAddonV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAddonV2Default creates a DeleteAddonV2Default with default headers values
func NewDeleteAddonV2Default(code int) *DeleteAddonV2Default {
	return &DeleteAddonV2Default{
		_statusCode: code,
	}
}

/* DeleteAddonV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteAddonV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete addon v2 default response
func (o *DeleteAddonV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete addon v2 default response has a 2xx status code
func (o *DeleteAddonV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete addon v2 default response has a 3xx status code
func (o *DeleteAddonV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete addon v2 default response has a 4xx status code
func (o *DeleteAddonV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete addon v2 default response has a 5xx status code
func (o *DeleteAddonV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete addon v2 default response a status code equal to that given
func (o *DeleteAddonV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteAddonV2Default) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAddonV2Default) String() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}][%d] deleteAddonV2 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAddonV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteAddonV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
