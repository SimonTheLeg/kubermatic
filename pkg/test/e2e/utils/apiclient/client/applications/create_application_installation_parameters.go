// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// NewCreateApplicationInstallationParams creates a new CreateApplicationInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateApplicationInstallationParams() *CreateApplicationInstallationParams {
	return &CreateApplicationInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApplicationInstallationParamsWithTimeout creates a new CreateApplicationInstallationParams object
// with the ability to set a timeout on a request.
func NewCreateApplicationInstallationParamsWithTimeout(timeout time.Duration) *CreateApplicationInstallationParams {
	return &CreateApplicationInstallationParams{
		timeout: timeout,
	}
}

// NewCreateApplicationInstallationParamsWithContext creates a new CreateApplicationInstallationParams object
// with the ability to set a context for a request.
func NewCreateApplicationInstallationParamsWithContext(ctx context.Context) *CreateApplicationInstallationParams {
	return &CreateApplicationInstallationParams{
		Context: ctx,
	}
}

// NewCreateApplicationInstallationParamsWithHTTPClient creates a new CreateApplicationInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateApplicationInstallationParamsWithHTTPClient(client *http.Client) *CreateApplicationInstallationParams {
	return &CreateApplicationInstallationParams{
		HTTPClient: client,
	}
}

/* CreateApplicationInstallationParams contains all the parameters to send to the API endpoint
   for the create application installation operation.

   Typically these are written to a http.Request.
*/
type CreateApplicationInstallationParams struct {

	// Body.
	Body *models.ApplicationInstallation

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create application installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateApplicationInstallationParams) WithDefaults() *CreateApplicationInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create application installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateApplicationInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create application installation params
func (o *CreateApplicationInstallationParams) WithTimeout(timeout time.Duration) *CreateApplicationInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create application installation params
func (o *CreateApplicationInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create application installation params
func (o *CreateApplicationInstallationParams) WithContext(ctx context.Context) *CreateApplicationInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create application installation params
func (o *CreateApplicationInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create application installation params
func (o *CreateApplicationInstallationParams) WithHTTPClient(client *http.Client) *CreateApplicationInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create application installation params
func (o *CreateApplicationInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create application installation params
func (o *CreateApplicationInstallationParams) WithBody(body *models.ApplicationInstallation) *CreateApplicationInstallationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create application installation params
func (o *CreateApplicationInstallationParams) SetBody(body *models.ApplicationInstallation) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create application installation params
func (o *CreateApplicationInstallationParams) WithClusterID(clusterID string) *CreateApplicationInstallationParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create application installation params
func (o *CreateApplicationInstallationParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create application installation params
func (o *CreateApplicationInstallationParams) WithProjectID(projectID string) *CreateApplicationInstallationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create application installation params
func (o *CreateApplicationInstallationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApplicationInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
