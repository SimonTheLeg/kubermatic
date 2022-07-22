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

// NewUpdateApplicationInstallationParams creates a new UpdateApplicationInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateApplicationInstallationParams() *UpdateApplicationInstallationParams {
	return &UpdateApplicationInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateApplicationInstallationParamsWithTimeout creates a new UpdateApplicationInstallationParams object
// with the ability to set a timeout on a request.
func NewUpdateApplicationInstallationParamsWithTimeout(timeout time.Duration) *UpdateApplicationInstallationParams {
	return &UpdateApplicationInstallationParams{
		timeout: timeout,
	}
}

// NewUpdateApplicationInstallationParamsWithContext creates a new UpdateApplicationInstallationParams object
// with the ability to set a context for a request.
func NewUpdateApplicationInstallationParamsWithContext(ctx context.Context) *UpdateApplicationInstallationParams {
	return &UpdateApplicationInstallationParams{
		Context: ctx,
	}
}

// NewUpdateApplicationInstallationParamsWithHTTPClient creates a new UpdateApplicationInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateApplicationInstallationParamsWithHTTPClient(client *http.Client) *UpdateApplicationInstallationParams {
	return &UpdateApplicationInstallationParams{
		HTTPClient: client,
	}
}

/* UpdateApplicationInstallationParams contains all the parameters to send to the API endpoint
   for the update application installation operation.

   Typically these are written to a http.Request.
*/
type UpdateApplicationInstallationParams struct {

	// Body.
	Body *models.ApplicationInstallationStatus

	// AppinstallName.
	ApplicationInstallationName string

	// ClusterID.
	ClusterID string

	// Namespace.
	Namespace string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update application installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateApplicationInstallationParams) WithDefaults() *UpdateApplicationInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update application installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateApplicationInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update application installation params
func (o *UpdateApplicationInstallationParams) WithTimeout(timeout time.Duration) *UpdateApplicationInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update application installation params
func (o *UpdateApplicationInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update application installation params
func (o *UpdateApplicationInstallationParams) WithContext(ctx context.Context) *UpdateApplicationInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update application installation params
func (o *UpdateApplicationInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update application installation params
func (o *UpdateApplicationInstallationParams) WithHTTPClient(client *http.Client) *UpdateApplicationInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update application installation params
func (o *UpdateApplicationInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update application installation params
func (o *UpdateApplicationInstallationParams) WithBody(body *models.ApplicationInstallationStatus) *UpdateApplicationInstallationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update application installation params
func (o *UpdateApplicationInstallationParams) SetBody(body *models.ApplicationInstallationStatus) {
	o.Body = body
}

// WithApplicationInstallationName adds the appinstallName to the update application installation params
func (o *UpdateApplicationInstallationParams) WithApplicationInstallationName(appinstallName string) *UpdateApplicationInstallationParams {
	o.SetApplicationInstallationName(appinstallName)
	return o
}

// SetApplicationInstallationName adds the appinstallName to the update application installation params
func (o *UpdateApplicationInstallationParams) SetApplicationInstallationName(appinstallName string) {
	o.ApplicationInstallationName = appinstallName
}

// WithClusterID adds the clusterID to the update application installation params
func (o *UpdateApplicationInstallationParams) WithClusterID(clusterID string) *UpdateApplicationInstallationParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update application installation params
func (o *UpdateApplicationInstallationParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNamespace adds the namespace to the update application installation params
func (o *UpdateApplicationInstallationParams) WithNamespace(namespace string) *UpdateApplicationInstallationParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update application installation params
func (o *UpdateApplicationInstallationParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the update application installation params
func (o *UpdateApplicationInstallationParams) WithProjectID(projectID string) *UpdateApplicationInstallationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update application installation params
func (o *UpdateApplicationInstallationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateApplicationInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param appinstall_name
	if err := r.SetPathParam("appinstall_name", o.ApplicationInstallationName); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
