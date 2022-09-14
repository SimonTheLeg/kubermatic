// Code generated by go-swagger; DO NOT EDIT.

package nutanix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListNutanixCategoryValuesNoCredentialsReader is a Reader for the ListNutanixCategoryValuesNoCredentials structure.
type ListNutanixCategoryValuesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNutanixCategoryValuesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNutanixCategoryValuesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListNutanixCategoryValuesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNutanixCategoryValuesNoCredentialsOK creates a ListNutanixCategoryValuesNoCredentialsOK with default headers values
func NewListNutanixCategoryValuesNoCredentialsOK() *ListNutanixCategoryValuesNoCredentialsOK {
	return &ListNutanixCategoryValuesNoCredentialsOK{}
}

/* ListNutanixCategoryValuesNoCredentialsOK describes a response with status code 200, with default header values.

NutanixCategoryValueList
*/
type ListNutanixCategoryValuesNoCredentialsOK struct {
	Payload models.NutanixCategoryValueList
}

// IsSuccess returns true when this list nutanix category values no credentials o k response has a 2xx status code
func (o *ListNutanixCategoryValuesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list nutanix category values no credentials o k response has a 3xx status code
func (o *ListNutanixCategoryValuesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list nutanix category values no credentials o k response has a 4xx status code
func (o *ListNutanixCategoryValuesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list nutanix category values no credentials o k response has a 5xx status code
func (o *ListNutanixCategoryValuesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list nutanix category values no credentials o k response a status code equal to that given
func (o *ListNutanixCategoryValuesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListNutanixCategoryValuesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/categories/{category}/values][%d] listNutanixCategoryValuesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListNutanixCategoryValuesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/categories/{category}/values][%d] listNutanixCategoryValuesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListNutanixCategoryValuesNoCredentialsOK) GetPayload() models.NutanixCategoryValueList {
	return o.Payload
}

func (o *ListNutanixCategoryValuesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNutanixCategoryValuesNoCredentialsDefault creates a ListNutanixCategoryValuesNoCredentialsDefault with default headers values
func NewListNutanixCategoryValuesNoCredentialsDefault(code int) *ListNutanixCategoryValuesNoCredentialsDefault {
	return &ListNutanixCategoryValuesNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListNutanixCategoryValuesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListNutanixCategoryValuesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list nutanix category values no credentials default response
func (o *ListNutanixCategoryValuesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list nutanix category values no credentials default response has a 2xx status code
func (o *ListNutanixCategoryValuesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list nutanix category values no credentials default response has a 3xx status code
func (o *ListNutanixCategoryValuesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list nutanix category values no credentials default response has a 4xx status code
func (o *ListNutanixCategoryValuesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list nutanix category values no credentials default response has a 5xx status code
func (o *ListNutanixCategoryValuesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list nutanix category values no credentials default response a status code equal to that given
func (o *ListNutanixCategoryValuesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListNutanixCategoryValuesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/categories/{category}/values][%d] listNutanixCategoryValuesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListNutanixCategoryValuesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/categories/{category}/values][%d] listNutanixCategoryValuesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListNutanixCategoryValuesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNutanixCategoryValuesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
