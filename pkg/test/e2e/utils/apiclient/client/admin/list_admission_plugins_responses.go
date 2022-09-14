// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAdmissionPluginsReader is a Reader for the ListAdmissionPlugins structure.
type ListAdmissionPluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAdmissionPluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAdmissionPluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAdmissionPluginsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAdmissionPluginsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAdmissionPluginsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAdmissionPluginsOK creates a ListAdmissionPluginsOK with default headers values
func NewListAdmissionPluginsOK() *ListAdmissionPluginsOK {
	return &ListAdmissionPluginsOK{}
}

/* ListAdmissionPluginsOK describes a response with status code 200, with default header values.

AdmissionPlugin
*/
type ListAdmissionPluginsOK struct {
	Payload []*models.AdmissionPlugin
}

// IsSuccess returns true when this list admission plugins o k response has a 2xx status code
func (o *ListAdmissionPluginsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list admission plugins o k response has a 3xx status code
func (o *ListAdmissionPluginsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list admission plugins o k response has a 4xx status code
func (o *ListAdmissionPluginsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list admission plugins o k response has a 5xx status code
func (o *ListAdmissionPluginsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list admission plugins o k response a status code equal to that given
func (o *ListAdmissionPluginsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAdmissionPluginsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPluginsOK  %+v", 200, o.Payload)
}

func (o *ListAdmissionPluginsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPluginsOK  %+v", 200, o.Payload)
}

func (o *ListAdmissionPluginsOK) GetPayload() []*models.AdmissionPlugin {
	return o.Payload
}

func (o *ListAdmissionPluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAdmissionPluginsUnauthorized creates a ListAdmissionPluginsUnauthorized with default headers values
func NewListAdmissionPluginsUnauthorized() *ListAdmissionPluginsUnauthorized {
	return &ListAdmissionPluginsUnauthorized{}
}

/* ListAdmissionPluginsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListAdmissionPluginsUnauthorized struct {
}

// IsSuccess returns true when this list admission plugins unauthorized response has a 2xx status code
func (o *ListAdmissionPluginsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list admission plugins unauthorized response has a 3xx status code
func (o *ListAdmissionPluginsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list admission plugins unauthorized response has a 4xx status code
func (o *ListAdmissionPluginsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list admission plugins unauthorized response has a 5xx status code
func (o *ListAdmissionPluginsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list admission plugins unauthorized response a status code equal to that given
func (o *ListAdmissionPluginsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListAdmissionPluginsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPluginsUnauthorized ", 401)
}

func (o *ListAdmissionPluginsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPluginsUnauthorized ", 401)
}

func (o *ListAdmissionPluginsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAdmissionPluginsForbidden creates a ListAdmissionPluginsForbidden with default headers values
func NewListAdmissionPluginsForbidden() *ListAdmissionPluginsForbidden {
	return &ListAdmissionPluginsForbidden{}
}

/* ListAdmissionPluginsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListAdmissionPluginsForbidden struct {
}

// IsSuccess returns true when this list admission plugins forbidden response has a 2xx status code
func (o *ListAdmissionPluginsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list admission plugins forbidden response has a 3xx status code
func (o *ListAdmissionPluginsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list admission plugins forbidden response has a 4xx status code
func (o *ListAdmissionPluginsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list admission plugins forbidden response has a 5xx status code
func (o *ListAdmissionPluginsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list admission plugins forbidden response a status code equal to that given
func (o *ListAdmissionPluginsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListAdmissionPluginsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPluginsForbidden ", 403)
}

func (o *ListAdmissionPluginsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPluginsForbidden ", 403)
}

func (o *ListAdmissionPluginsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAdmissionPluginsDefault creates a ListAdmissionPluginsDefault with default headers values
func NewListAdmissionPluginsDefault(code int) *ListAdmissionPluginsDefault {
	return &ListAdmissionPluginsDefault{
		_statusCode: code,
	}
}

/* ListAdmissionPluginsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAdmissionPluginsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list admission plugins default response
func (o *ListAdmissionPluginsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list admission plugins default response has a 2xx status code
func (o *ListAdmissionPluginsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list admission plugins default response has a 3xx status code
func (o *ListAdmissionPluginsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list admission plugins default response has a 4xx status code
func (o *ListAdmissionPluginsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list admission plugins default response has a 5xx status code
func (o *ListAdmissionPluginsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list admission plugins default response a status code equal to that given
func (o *ListAdmissionPluginsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAdmissionPluginsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPlugins default  %+v", o._statusCode, o.Payload)
}

func (o *ListAdmissionPluginsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/admission/plugins][%d] listAdmissionPlugins default  %+v", o._statusCode, o.Payload)
}

func (o *ListAdmissionPluginsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAdmissionPluginsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
