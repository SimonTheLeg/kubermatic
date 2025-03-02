// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewUpdateAlertmanagerParams creates a new UpdateAlertmanagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAlertmanagerParams() *UpdateAlertmanagerParams {
	return &UpdateAlertmanagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAlertmanagerParamsWithTimeout creates a new UpdateAlertmanagerParams object
// with the ability to set a timeout on a request.
func NewUpdateAlertmanagerParamsWithTimeout(timeout time.Duration) *UpdateAlertmanagerParams {
	return &UpdateAlertmanagerParams{
		timeout: timeout,
	}
}

// NewUpdateAlertmanagerParamsWithContext creates a new UpdateAlertmanagerParams object
// with the ability to set a context for a request.
func NewUpdateAlertmanagerParamsWithContext(ctx context.Context) *UpdateAlertmanagerParams {
	return &UpdateAlertmanagerParams{
		Context: ctx,
	}
}

// NewUpdateAlertmanagerParamsWithHTTPClient creates a new UpdateAlertmanagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAlertmanagerParamsWithHTTPClient(client *http.Client) *UpdateAlertmanagerParams {
	return &UpdateAlertmanagerParams{
		HTTPClient: client,
	}
}

/*
UpdateAlertmanagerParams contains all the parameters to send to the API endpoint

	for the update alertmanager operation.

	Typically these are written to a http.Request.
*/
type UpdateAlertmanagerParams struct {

	// Body.
	Body *models.Alertmanager

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update alertmanager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAlertmanagerParams) WithDefaults() *UpdateAlertmanagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update alertmanager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAlertmanagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update alertmanager params
func (o *UpdateAlertmanagerParams) WithTimeout(timeout time.Duration) *UpdateAlertmanagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update alertmanager params
func (o *UpdateAlertmanagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update alertmanager params
func (o *UpdateAlertmanagerParams) WithContext(ctx context.Context) *UpdateAlertmanagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update alertmanager params
func (o *UpdateAlertmanagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update alertmanager params
func (o *UpdateAlertmanagerParams) WithHTTPClient(client *http.Client) *UpdateAlertmanagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update alertmanager params
func (o *UpdateAlertmanagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update alertmanager params
func (o *UpdateAlertmanagerParams) WithBody(body *models.Alertmanager) *UpdateAlertmanagerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update alertmanager params
func (o *UpdateAlertmanagerParams) SetBody(body *models.Alertmanager) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update alertmanager params
func (o *UpdateAlertmanagerParams) WithClusterID(clusterID string) *UpdateAlertmanagerParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update alertmanager params
func (o *UpdateAlertmanagerParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the update alertmanager params
func (o *UpdateAlertmanagerParams) WithProjectID(projectID string) *UpdateAlertmanagerParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update alertmanager params
func (o *UpdateAlertmanagerParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAlertmanagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
