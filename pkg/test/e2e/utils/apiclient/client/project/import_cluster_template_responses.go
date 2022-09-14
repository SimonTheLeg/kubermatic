// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ImportClusterTemplateReader is a Reader for the ImportClusterTemplate structure.
type ImportClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImportClusterTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImportClusterTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportClusterTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImportClusterTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImportClusterTemplateCreated creates a ImportClusterTemplateCreated with default headers values
func NewImportClusterTemplateCreated() *ImportClusterTemplateCreated {
	return &ImportClusterTemplateCreated{}
}

/* ImportClusterTemplateCreated describes a response with status code 201, with default header values.

ClusterTemplate
*/
type ImportClusterTemplateCreated struct {
	Payload *models.ClusterTemplate
}

// IsSuccess returns true when this import cluster template created response has a 2xx status code
func (o *ImportClusterTemplateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import cluster template created response has a 3xx status code
func (o *ImportClusterTemplateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import cluster template created response has a 4xx status code
func (o *ImportClusterTemplateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this import cluster template created response has a 5xx status code
func (o *ImportClusterTemplateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this import cluster template created response a status code equal to that given
func (o *ImportClusterTemplateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ImportClusterTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplateCreated  %+v", 201, o.Payload)
}

func (o *ImportClusterTemplateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplateCreated  %+v", 201, o.Payload)
}

func (o *ImportClusterTemplateCreated) GetPayload() *models.ClusterTemplate {
	return o.Payload
}

func (o *ImportClusterTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportClusterTemplateUnauthorized creates a ImportClusterTemplateUnauthorized with default headers values
func NewImportClusterTemplateUnauthorized() *ImportClusterTemplateUnauthorized {
	return &ImportClusterTemplateUnauthorized{}
}

/* ImportClusterTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ImportClusterTemplateUnauthorized struct {
}

// IsSuccess returns true when this import cluster template unauthorized response has a 2xx status code
func (o *ImportClusterTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import cluster template unauthorized response has a 3xx status code
func (o *ImportClusterTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import cluster template unauthorized response has a 4xx status code
func (o *ImportClusterTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this import cluster template unauthorized response has a 5xx status code
func (o *ImportClusterTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this import cluster template unauthorized response a status code equal to that given
func (o *ImportClusterTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImportClusterTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplateUnauthorized ", 401)
}

func (o *ImportClusterTemplateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplateUnauthorized ", 401)
}

func (o *ImportClusterTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportClusterTemplateForbidden creates a ImportClusterTemplateForbidden with default headers values
func NewImportClusterTemplateForbidden() *ImportClusterTemplateForbidden {
	return &ImportClusterTemplateForbidden{}
}

/* ImportClusterTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ImportClusterTemplateForbidden struct {
}

// IsSuccess returns true when this import cluster template forbidden response has a 2xx status code
func (o *ImportClusterTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import cluster template forbidden response has a 3xx status code
func (o *ImportClusterTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import cluster template forbidden response has a 4xx status code
func (o *ImportClusterTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this import cluster template forbidden response has a 5xx status code
func (o *ImportClusterTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this import cluster template forbidden response a status code equal to that given
func (o *ImportClusterTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImportClusterTemplateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplateForbidden ", 403)
}

func (o *ImportClusterTemplateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplateForbidden ", 403)
}

func (o *ImportClusterTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportClusterTemplateDefault creates a ImportClusterTemplateDefault with default headers values
func NewImportClusterTemplateDefault(code int) *ImportClusterTemplateDefault {
	return &ImportClusterTemplateDefault{
		_statusCode: code,
	}
}

/* ImportClusterTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ImportClusterTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the import cluster template default response
func (o *ImportClusterTemplateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this import cluster template default response has a 2xx status code
func (o *ImportClusterTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this import cluster template default response has a 3xx status code
func (o *ImportClusterTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this import cluster template default response has a 4xx status code
func (o *ImportClusterTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this import cluster template default response has a 5xx status code
func (o *ImportClusterTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this import cluster template default response a status code equal to that given
func (o *ImportClusterTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ImportClusterTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *ImportClusterTemplateDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates/import][%d] importClusterTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *ImportClusterTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImportClusterTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ImportClusterTemplateBody import cluster template body
swagger:model ImportClusterTemplateBody
*/
type ImportClusterTemplateBody struct {

	// Annotations that can be added to the resource
	Annotations map[string]string `json:"annotations,omitempty"`

	// applications
	Applications []*models.Application `json:"applications"`

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	// Format: date-time
	DeletionTimestamp strfmt.DateTime `json:"deletionTimestamp,omitempty"`

	// ID unique value that identifies the resource generated by the server. Read-Only.
	ID string `json:"id,omitempty"`

	// Name represents human readable name for the resource
	Name string `json:"name,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// user SSH keys
	UserSSHKeys []*models.ClusterTemplateSSHKey `json:"userSshKeys"`

	// cluster
	Cluster *models.ClusterTemplateInfo `json:"cluster,omitempty"`

	// node deployment
	NodeDeployment *models.ClusterTemplateNodeDeployment `json:"nodeDeployment,omitempty"`
}

// Validate validates this import cluster template body
func (o *ImportClusterTemplateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeletionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserSSHKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNodeDeployment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImportClusterTemplateBody) validateApplications(formats strfmt.Registry) error {
	if swag.IsZero(o.Applications) { // not required
		return nil
	}

	for i := 0; i < len(o.Applications); i++ {
		if swag.IsZero(o.Applications[i]) { // not required
			continue
		}

		if o.Applications[i] != nil {
			if err := o.Applications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Body" + "." + "applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Body" + "." + "applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ImportClusterTemplateBody) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("Body"+"."+"creationTimestamp", "body", "date-time", o.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ImportClusterTemplateBody) validateDeletionTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.DeletionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("Body"+"."+"deletionTimestamp", "body", "date-time", o.DeletionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ImportClusterTemplateBody) validateUserSSHKeys(formats strfmt.Registry) error {
	if swag.IsZero(o.UserSSHKeys) { // not required
		return nil
	}

	for i := 0; i < len(o.UserSSHKeys); i++ {
		if swag.IsZero(o.UserSSHKeys[i]) { // not required
			continue
		}

		if o.UserSSHKeys[i] != nil {
			if err := o.UserSSHKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Body" + "." + "userSshKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Body" + "." + "userSshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ImportClusterTemplateBody) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(o.Cluster) { // not required
		return nil
	}

	if o.Cluster != nil {
		if err := o.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *ImportClusterTemplateBody) validateNodeDeployment(formats strfmt.Registry) error {
	if swag.IsZero(o.NodeDeployment) { // not required
		return nil
	}

	if o.NodeDeployment != nil {
		if err := o.NodeDeployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "nodeDeployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "nodeDeployment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this import cluster template body based on the context it is used
func (o *ImportClusterTemplateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUserSSHKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNodeDeployment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImportClusterTemplateBody) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Applications); i++ {

		if o.Applications[i] != nil {
			if err := o.Applications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Body" + "." + "applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Body" + "." + "applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ImportClusterTemplateBody) contextValidateUserSSHKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.UserSSHKeys); i++ {

		if o.UserSSHKeys[i] != nil {
			if err := o.UserSSHKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Body" + "." + "userSshKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Body" + "." + "userSshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ImportClusterTemplateBody) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if o.Cluster != nil {
		if err := o.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *ImportClusterTemplateBody) contextValidateNodeDeployment(ctx context.Context, formats strfmt.Registry) error {

	if o.NodeDeployment != nil {
		if err := o.NodeDeployment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "nodeDeployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Body" + "." + "nodeDeployment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ImportClusterTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImportClusterTemplateBody) UnmarshalBinary(b []byte) error {
	var res ImportClusterTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
