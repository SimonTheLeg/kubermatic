// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// UpdateDCReader is a Reader for the UpdateDC structure.
type UpdateDCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDCOK creates a UpdateDCOK with default headers values
func NewUpdateDCOK() *UpdateDCOK {
	return &UpdateDCOK{}
}

/*
UpdateDCOK describes a response with status code 200, with default header values.

Datacenter
*/
type UpdateDCOK struct {
	Payload *models.Datacenter
}

// IsSuccess returns true when this update d c o k response has a 2xx status code
func (o *UpdateDCOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update d c o k response has a 3xx status code
func (o *UpdateDCOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update d c o k response has a 4xx status code
func (o *UpdateDCOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update d c o k response has a 5xx status code
func (o *UpdateDCOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update d c o k response a status code equal to that given
func (o *UpdateDCOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateDCOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDCOK  %+v", 200, o.Payload)
}

func (o *UpdateDCOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDCOK  %+v", 200, o.Payload)
}

func (o *UpdateDCOK) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *UpdateDCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDCUnauthorized creates a UpdateDCUnauthorized with default headers values
func NewUpdateDCUnauthorized() *UpdateDCUnauthorized {
	return &UpdateDCUnauthorized{}
}

/*
UpdateDCUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpdateDCUnauthorized struct {
}

// IsSuccess returns true when this update d c unauthorized response has a 2xx status code
func (o *UpdateDCUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update d c unauthorized response has a 3xx status code
func (o *UpdateDCUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update d c unauthorized response has a 4xx status code
func (o *UpdateDCUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update d c unauthorized response has a 5xx status code
func (o *UpdateDCUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update d c unauthorized response a status code equal to that given
func (o *UpdateDCUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateDCUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDCUnauthorized ", 401)
}

func (o *UpdateDCUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDCUnauthorized ", 401)
}

func (o *UpdateDCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDCForbidden creates a UpdateDCForbidden with default headers values
func NewUpdateDCForbidden() *UpdateDCForbidden {
	return &UpdateDCForbidden{}
}

/*
UpdateDCForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpdateDCForbidden struct {
}

// IsSuccess returns true when this update d c forbidden response has a 2xx status code
func (o *UpdateDCForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update d c forbidden response has a 3xx status code
func (o *UpdateDCForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update d c forbidden response has a 4xx status code
func (o *UpdateDCForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update d c forbidden response has a 5xx status code
func (o *UpdateDCForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update d c forbidden response a status code equal to that given
func (o *UpdateDCForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateDCForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDCForbidden ", 403)
}

func (o *UpdateDCForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDCForbidden ", 403)
}

func (o *UpdateDCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDCDefault creates a UpdateDCDefault with default headers values
func NewUpdateDCDefault(code int) *UpdateDCDefault {
	return &UpdateDCDefault{
		_statusCode: code,
	}
}

/*
UpdateDCDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateDCDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update d c default response
func (o *UpdateDCDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update d c default response has a 2xx status code
func (o *UpdateDCDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update d c default response has a 3xx status code
func (o *UpdateDCDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update d c default response has a 4xx status code
func (o *UpdateDCDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update d c default response has a 5xx status code
func (o *UpdateDCDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update d c default response a status code equal to that given
func (o *UpdateDCDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateDCDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDC default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDCDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/seed/{seed_name}/dc/{dc}][%d] updateDC default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDCDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDCBody update d c body
swagger:model UpdateDCBody
*/
type UpdateDCBody struct {

	// name
	Name string `json:"name,omitempty"`

	// spec
	Spec *models.DatacenterSpec `json:"spec,omitempty"`
}

// Validate validates this update d c body
func (o *UpdateDCBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDCBody) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(o.Spec) { // not required
		return nil
	}

	if o.Spec != nil {
		if err := o.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update d c body based on the context it is used
func (o *UpdateDCBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDCBody) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if o.Spec != nil {
		if err := o.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDCBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDCBody) UnmarshalBinary(b []byte) error {
	var res UpdateDCBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
