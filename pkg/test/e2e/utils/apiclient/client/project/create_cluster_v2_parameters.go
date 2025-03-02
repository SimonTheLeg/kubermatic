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

// NewCreateClusterV2Params creates a new CreateClusterV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterV2Params() *CreateClusterV2Params {
	return &CreateClusterV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterV2ParamsWithTimeout creates a new CreateClusterV2Params object
// with the ability to set a timeout on a request.
func NewCreateClusterV2ParamsWithTimeout(timeout time.Duration) *CreateClusterV2Params {
	return &CreateClusterV2Params{
		timeout: timeout,
	}
}

// NewCreateClusterV2ParamsWithContext creates a new CreateClusterV2Params object
// with the ability to set a context for a request.
func NewCreateClusterV2ParamsWithContext(ctx context.Context) *CreateClusterV2Params {
	return &CreateClusterV2Params{
		Context: ctx,
	}
}

// NewCreateClusterV2ParamsWithHTTPClient creates a new CreateClusterV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterV2ParamsWithHTTPClient(client *http.Client) *CreateClusterV2Params {
	return &CreateClusterV2Params{
		HTTPClient: client,
	}
}

/*
CreateClusterV2Params contains all the parameters to send to the API endpoint

	for the create cluster v2 operation.

	Typically these are written to a http.Request.
*/
type CreateClusterV2Params struct {

	// Body.
	Body *models.CreateClusterSpec

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterV2Params) WithDefaults() *CreateClusterV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster v2 params
func (o *CreateClusterV2Params) WithTimeout(timeout time.Duration) *CreateClusterV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster v2 params
func (o *CreateClusterV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster v2 params
func (o *CreateClusterV2Params) WithContext(ctx context.Context) *CreateClusterV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster v2 params
func (o *CreateClusterV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster v2 params
func (o *CreateClusterV2Params) WithHTTPClient(client *http.Client) *CreateClusterV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster v2 params
func (o *CreateClusterV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster v2 params
func (o *CreateClusterV2Params) WithBody(body *models.CreateClusterSpec) *CreateClusterV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster v2 params
func (o *CreateClusterV2Params) SetBody(body *models.CreateClusterSpec) {
	o.Body = body
}

// WithProjectID adds the projectID to the create cluster v2 params
func (o *CreateClusterV2Params) WithProjectID(projectID string) *CreateClusterV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create cluster v2 params
func (o *CreateClusterV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
