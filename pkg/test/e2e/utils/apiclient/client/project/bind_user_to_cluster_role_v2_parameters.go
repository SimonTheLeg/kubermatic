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

// NewBindUserToClusterRoleV2Params creates a new BindUserToClusterRoleV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBindUserToClusterRoleV2Params() *BindUserToClusterRoleV2Params {
	return &BindUserToClusterRoleV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewBindUserToClusterRoleV2ParamsWithTimeout creates a new BindUserToClusterRoleV2Params object
// with the ability to set a timeout on a request.
func NewBindUserToClusterRoleV2ParamsWithTimeout(timeout time.Duration) *BindUserToClusterRoleV2Params {
	return &BindUserToClusterRoleV2Params{
		timeout: timeout,
	}
}

// NewBindUserToClusterRoleV2ParamsWithContext creates a new BindUserToClusterRoleV2Params object
// with the ability to set a context for a request.
func NewBindUserToClusterRoleV2ParamsWithContext(ctx context.Context) *BindUserToClusterRoleV2Params {
	return &BindUserToClusterRoleV2Params{
		Context: ctx,
	}
}

// NewBindUserToClusterRoleV2ParamsWithHTTPClient creates a new BindUserToClusterRoleV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewBindUserToClusterRoleV2ParamsWithHTTPClient(client *http.Client) *BindUserToClusterRoleV2Params {
	return &BindUserToClusterRoleV2Params{
		HTTPClient: client,
	}
}

/*
BindUserToClusterRoleV2Params contains all the parameters to send to the API endpoint

	for the bind user to cluster role v2 operation.

	Typically these are written to a http.Request.
*/
type BindUserToClusterRoleV2Params struct {

	// Body.
	Body *models.ClusterRoleUser

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	// RoleID.
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bind user to cluster role v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BindUserToClusterRoleV2Params) WithDefaults() *BindUserToClusterRoleV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bind user to cluster role v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BindUserToClusterRoleV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) WithTimeout(timeout time.Duration) *BindUserToClusterRoleV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) WithContext(ctx context.Context) *BindUserToClusterRoleV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) WithHTTPClient(client *http.Client) *BindUserToClusterRoleV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) WithBody(body *models.ClusterRoleUser) *BindUserToClusterRoleV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) SetBody(body *models.ClusterRoleUser) {
	o.Body = body
}

// WithClusterID adds the clusterID to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) WithClusterID(clusterID string) *BindUserToClusterRoleV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) WithProjectID(projectID string) *BindUserToClusterRoleV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRoleID adds the roleID to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) WithRoleID(roleID string) *BindUserToClusterRoleV2Params {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the bind user to cluster role v2 params
func (o *BindUserToClusterRoleV2Params) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *BindUserToClusterRoleV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
