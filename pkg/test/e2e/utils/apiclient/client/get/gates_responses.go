// Code generated by go-swagger; DO NOT EDIT.

package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GatesReader is a Reader for the Gates structure.
type GatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGatesOK creates a GatesOK with default headers values
func NewGatesOK() *GatesOK {
	return &GatesOK{}
}

/* GatesOK describes a response with status code 200, with default header values.

FeatureGates
*/
type GatesOK struct {
	Payload *models.FeatureGates
}

// IsSuccess returns true when this gates o k response has a 2xx status code
func (o *GatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gates o k response has a 3xx status code
func (o *GatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gates o k response has a 4xx status code
func (o *GatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this gates o k response has a 5xx status code
func (o *GatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this gates o k response a status code equal to that given
func (o *GatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GatesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gatesOK  %+v", 200, o.Payload)
}

func (o *GatesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gatesOK  %+v", 200, o.Payload)
}

func (o *GatesOK) GetPayload() *models.FeatureGates {
	return o.Payload
}

func (o *GatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeatureGates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGatesUnauthorized creates a GatesUnauthorized with default headers values
func NewGatesUnauthorized() *GatesUnauthorized {
	return &GatesUnauthorized{}
}

/* GatesUnauthorized describes a response with status code 401, with default header values.

errorResponse
*/
type GatesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this gates unauthorized response has a 2xx status code
func (o *GatesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this gates unauthorized response has a 3xx status code
func (o *GatesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gates unauthorized response has a 4xx status code
func (o *GatesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this gates unauthorized response has a 5xx status code
func (o *GatesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this gates unauthorized response a status code equal to that given
func (o *GatesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GatesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GatesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGatesForbidden creates a GatesForbidden with default headers values
func NewGatesForbidden() *GatesForbidden {
	return &GatesForbidden{}
}

/* GatesForbidden describes a response with status code 403, with default header values.

errorResponse
*/
type GatesForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this gates forbidden response has a 2xx status code
func (o *GatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this gates forbidden response has a 3xx status code
func (o *GatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gates forbidden response has a 4xx status code
func (o *GatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this gates forbidden response has a 5xx status code
func (o *GatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this gates forbidden response a status code equal to that given
func (o *GatesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GatesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gatesForbidden  %+v", 403, o.Payload)
}

func (o *GatesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gatesForbidden  %+v", 403, o.Payload)
}

func (o *GatesForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGatesDefault creates a GatesDefault with default headers values
func NewGatesDefault(code int) *GatesDefault {
	return &GatesDefault{
		_statusCode: code,
	}
}

/* GatesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GatesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gates default response
func (o *GatesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this gates default response has a 2xx status code
func (o *GatesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gates default response has a 3xx status code
func (o *GatesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gates default response has a 4xx status code
func (o *GatesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gates default response has a 5xx status code
func (o *GatesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gates default response a status code equal to that given
func (o *GatesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GatesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gates default  %+v", o._statusCode, o.Payload)
}

func (o *GatesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/featuregates][%d] gates default  %+v", o._statusCode, o.Payload)
}

func (o *GatesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
