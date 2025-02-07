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

// GetKubermaticCustomLinksReader is a Reader for the GetKubermaticCustomLinks structure.
type GetKubermaticCustomLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubermaticCustomLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubermaticCustomLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKubermaticCustomLinksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubermaticCustomLinksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetKubermaticCustomLinksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKubermaticCustomLinksOK creates a GetKubermaticCustomLinksOK with default headers values
func NewGetKubermaticCustomLinksOK() *GetKubermaticCustomLinksOK {
	return &GetKubermaticCustomLinksOK{}
}

/*
GetKubermaticCustomLinksOK describes a response with status code 200, with default header values.

GlobalCustomLinks
*/
type GetKubermaticCustomLinksOK struct {
	Payload models.GlobalCustomLinks
}

// IsSuccess returns true when this get kubermatic custom links o k response has a 2xx status code
func (o *GetKubermaticCustomLinksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubermatic custom links o k response has a 3xx status code
func (o *GetKubermaticCustomLinksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubermatic custom links o k response has a 4xx status code
func (o *GetKubermaticCustomLinksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubermatic custom links o k response has a 5xx status code
func (o *GetKubermaticCustomLinksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubermatic custom links o k response a status code equal to that given
func (o *GetKubermaticCustomLinksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKubermaticCustomLinksOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksOK  %+v", 200, o.Payload)
}

func (o *GetKubermaticCustomLinksOK) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksOK  %+v", 200, o.Payload)
}

func (o *GetKubermaticCustomLinksOK) GetPayload() models.GlobalCustomLinks {
	return o.Payload
}

func (o *GetKubermaticCustomLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubermaticCustomLinksUnauthorized creates a GetKubermaticCustomLinksUnauthorized with default headers values
func NewGetKubermaticCustomLinksUnauthorized() *GetKubermaticCustomLinksUnauthorized {
	return &GetKubermaticCustomLinksUnauthorized{}
}

/*
GetKubermaticCustomLinksUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetKubermaticCustomLinksUnauthorized struct {
}

// IsSuccess returns true when this get kubermatic custom links unauthorized response has a 2xx status code
func (o *GetKubermaticCustomLinksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubermatic custom links unauthorized response has a 3xx status code
func (o *GetKubermaticCustomLinksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubermatic custom links unauthorized response has a 4xx status code
func (o *GetKubermaticCustomLinksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubermatic custom links unauthorized response has a 5xx status code
func (o *GetKubermaticCustomLinksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubermatic custom links unauthorized response a status code equal to that given
func (o *GetKubermaticCustomLinksUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKubermaticCustomLinksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksUnauthorized ", 401)
}

func (o *GetKubermaticCustomLinksUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksUnauthorized ", 401)
}

func (o *GetKubermaticCustomLinksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubermaticCustomLinksForbidden creates a GetKubermaticCustomLinksForbidden with default headers values
func NewGetKubermaticCustomLinksForbidden() *GetKubermaticCustomLinksForbidden {
	return &GetKubermaticCustomLinksForbidden{}
}

/*
GetKubermaticCustomLinksForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetKubermaticCustomLinksForbidden struct {
}

// IsSuccess returns true when this get kubermatic custom links forbidden response has a 2xx status code
func (o *GetKubermaticCustomLinksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubermatic custom links forbidden response has a 3xx status code
func (o *GetKubermaticCustomLinksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubermatic custom links forbidden response has a 4xx status code
func (o *GetKubermaticCustomLinksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubermatic custom links forbidden response has a 5xx status code
func (o *GetKubermaticCustomLinksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubermatic custom links forbidden response a status code equal to that given
func (o *GetKubermaticCustomLinksForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKubermaticCustomLinksForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksForbidden ", 403)
}

func (o *GetKubermaticCustomLinksForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinksForbidden ", 403)
}

func (o *GetKubermaticCustomLinksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubermaticCustomLinksDefault creates a GetKubermaticCustomLinksDefault with default headers values
func NewGetKubermaticCustomLinksDefault(code int) *GetKubermaticCustomLinksDefault {
	return &GetKubermaticCustomLinksDefault{
		_statusCode: code,
	}
}

/*
GetKubermaticCustomLinksDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetKubermaticCustomLinksDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get kubermatic custom links default response
func (o *GetKubermaticCustomLinksDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get kubermatic custom links default response has a 2xx status code
func (o *GetKubermaticCustomLinksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get kubermatic custom links default response has a 3xx status code
func (o *GetKubermaticCustomLinksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get kubermatic custom links default response has a 4xx status code
func (o *GetKubermaticCustomLinksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get kubermatic custom links default response has a 5xx status code
func (o *GetKubermaticCustomLinksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get kubermatic custom links default response a status code equal to that given
func (o *GetKubermaticCustomLinksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetKubermaticCustomLinksDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinks default  %+v", o._statusCode, o.Payload)
}

func (o *GetKubermaticCustomLinksDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/settings/customlinks][%d] getKubermaticCustomLinks default  %+v", o._statusCode, o.Payload)
}

func (o *GetKubermaticCustomLinksDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetKubermaticCustomLinksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
