// Code generated by go-swagger; DO NOT EDIT.

package storage

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
)

// NewVolumeEfficiencyPolicyDeleteParams creates a new VolumeEfficiencyPolicyDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeEfficiencyPolicyDeleteParams() *VolumeEfficiencyPolicyDeleteParams {
	return &VolumeEfficiencyPolicyDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeEfficiencyPolicyDeleteParamsWithTimeout creates a new VolumeEfficiencyPolicyDeleteParams object
// with the ability to set a timeout on a request.
func NewVolumeEfficiencyPolicyDeleteParamsWithTimeout(timeout time.Duration) *VolumeEfficiencyPolicyDeleteParams {
	return &VolumeEfficiencyPolicyDeleteParams{
		timeout: timeout,
	}
}

// NewVolumeEfficiencyPolicyDeleteParamsWithContext creates a new VolumeEfficiencyPolicyDeleteParams object
// with the ability to set a context for a request.
func NewVolumeEfficiencyPolicyDeleteParamsWithContext(ctx context.Context) *VolumeEfficiencyPolicyDeleteParams {
	return &VolumeEfficiencyPolicyDeleteParams{
		Context: ctx,
	}
}

// NewVolumeEfficiencyPolicyDeleteParamsWithHTTPClient creates a new VolumeEfficiencyPolicyDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeEfficiencyPolicyDeleteParamsWithHTTPClient(client *http.Client) *VolumeEfficiencyPolicyDeleteParams {
	return &VolumeEfficiencyPolicyDeleteParams{
		HTTPClient: client,
	}
}

/* VolumeEfficiencyPolicyDeleteParams contains all the parameters to send to the API endpoint
   for the volume efficiency policy delete operation.

   Typically these are written to a http.Request.
*/
type VolumeEfficiencyPolicyDeleteParams struct {

	/* UUID.

	   Volume efficiency policy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume efficiency policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeEfficiencyPolicyDeleteParams) WithDefaults() *VolumeEfficiencyPolicyDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume efficiency policy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeEfficiencyPolicyDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) WithTimeout(timeout time.Duration) *VolumeEfficiencyPolicyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) WithContext(ctx context.Context) *VolumeEfficiencyPolicyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) WithHTTPClient(client *http.Client) *VolumeEfficiencyPolicyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUIDPathParameter adds the uuid to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) WithUUIDPathParameter(uuid string) *VolumeEfficiencyPolicyDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the volume efficiency policy delete params
func (o *VolumeEfficiencyPolicyDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeEfficiencyPolicyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
