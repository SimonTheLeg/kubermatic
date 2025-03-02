// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateAdminRuleGroupReader is a Reader for the CreateAdminRuleGroup structure.
type CreateAdminRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAdminRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAdminRuleGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAdminRuleGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAdminRuleGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateAdminRuleGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAdminRuleGroupCreated creates a CreateAdminRuleGroupCreated with default headers values
func NewCreateAdminRuleGroupCreated() *CreateAdminRuleGroupCreated {
	return &CreateAdminRuleGroupCreated{}
}

/*
CreateAdminRuleGroupCreated describes a response with status code 201, with default header values.

RuleGroup
*/
type CreateAdminRuleGroupCreated struct {
	Payload *models.RuleGroup
}

// IsSuccess returns true when this create admin rule group created response has a 2xx status code
func (o *CreateAdminRuleGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create admin rule group created response has a 3xx status code
func (o *CreateAdminRuleGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create admin rule group created response has a 4xx status code
func (o *CreateAdminRuleGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create admin rule group created response has a 5xx status code
func (o *CreateAdminRuleGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create admin rule group created response a status code equal to that given
func (o *CreateAdminRuleGroupCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateAdminRuleGroupCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateAdminRuleGroupCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateAdminRuleGroupCreated) GetPayload() *models.RuleGroup {
	return o.Payload
}

func (o *CreateAdminRuleGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAdminRuleGroupUnauthorized creates a CreateAdminRuleGroupUnauthorized with default headers values
func NewCreateAdminRuleGroupUnauthorized() *CreateAdminRuleGroupUnauthorized {
	return &CreateAdminRuleGroupUnauthorized{}
}

/*
CreateAdminRuleGroupUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateAdminRuleGroupUnauthorized struct {
}

// IsSuccess returns true when this create admin rule group unauthorized response has a 2xx status code
func (o *CreateAdminRuleGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create admin rule group unauthorized response has a 3xx status code
func (o *CreateAdminRuleGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create admin rule group unauthorized response has a 4xx status code
func (o *CreateAdminRuleGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create admin rule group unauthorized response has a 5xx status code
func (o *CreateAdminRuleGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create admin rule group unauthorized response a status code equal to that given
func (o *CreateAdminRuleGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateAdminRuleGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroupUnauthorized ", 401)
}

func (o *CreateAdminRuleGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroupUnauthorized ", 401)
}

func (o *CreateAdminRuleGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAdminRuleGroupForbidden creates a CreateAdminRuleGroupForbidden with default headers values
func NewCreateAdminRuleGroupForbidden() *CreateAdminRuleGroupForbidden {
	return &CreateAdminRuleGroupForbidden{}
}

/*
CreateAdminRuleGroupForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateAdminRuleGroupForbidden struct {
}

// IsSuccess returns true when this create admin rule group forbidden response has a 2xx status code
func (o *CreateAdminRuleGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create admin rule group forbidden response has a 3xx status code
func (o *CreateAdminRuleGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create admin rule group forbidden response has a 4xx status code
func (o *CreateAdminRuleGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create admin rule group forbidden response has a 5xx status code
func (o *CreateAdminRuleGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create admin rule group forbidden response a status code equal to that given
func (o *CreateAdminRuleGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateAdminRuleGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroupForbidden ", 403)
}

func (o *CreateAdminRuleGroupForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroupForbidden ", 403)
}

func (o *CreateAdminRuleGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAdminRuleGroupDefault creates a CreateAdminRuleGroupDefault with default headers values
func NewCreateAdminRuleGroupDefault(code int) *CreateAdminRuleGroupDefault {
	return &CreateAdminRuleGroupDefault{
		_statusCode: code,
	}
}

/*
CreateAdminRuleGroupDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateAdminRuleGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create admin rule group default response
func (o *CreateAdminRuleGroupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create admin rule group default response has a 2xx status code
func (o *CreateAdminRuleGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create admin rule group default response has a 3xx status code
func (o *CreateAdminRuleGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create admin rule group default response has a 4xx status code
func (o *CreateAdminRuleGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create admin rule group default response has a 5xx status code
func (o *CreateAdminRuleGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create admin rule group default response a status code equal to that given
func (o *CreateAdminRuleGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateAdminRuleGroupDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAdminRuleGroupDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/rulegroups][%d] createAdminRuleGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAdminRuleGroupDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAdminRuleGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
