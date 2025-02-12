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
	"github.com/go-openapi/swag"
)

// NewSnapshotPolicyScheduleGetParams creates a new SnapshotPolicyScheduleGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotPolicyScheduleGetParams() *SnapshotPolicyScheduleGetParams {
	return &SnapshotPolicyScheduleGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotPolicyScheduleGetParamsWithTimeout creates a new SnapshotPolicyScheduleGetParams object
// with the ability to set a timeout on a request.
func NewSnapshotPolicyScheduleGetParamsWithTimeout(timeout time.Duration) *SnapshotPolicyScheduleGetParams {
	return &SnapshotPolicyScheduleGetParams{
		timeout: timeout,
	}
}

// NewSnapshotPolicyScheduleGetParamsWithContext creates a new SnapshotPolicyScheduleGetParams object
// with the ability to set a context for a request.
func NewSnapshotPolicyScheduleGetParamsWithContext(ctx context.Context) *SnapshotPolicyScheduleGetParams {
	return &SnapshotPolicyScheduleGetParams{
		Context: ctx,
	}
}

// NewSnapshotPolicyScheduleGetParamsWithHTTPClient creates a new SnapshotPolicyScheduleGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotPolicyScheduleGetParamsWithHTTPClient(client *http.Client) *SnapshotPolicyScheduleGetParams {
	return &SnapshotPolicyScheduleGetParams{
		HTTPClient: client,
	}
}

/* SnapshotPolicyScheduleGetParams contains all the parameters to send to the API endpoint
   for the snapshot policy schedule get operation.

   Typically these are written to a http.Request.
*/
type SnapshotPolicyScheduleGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* ScheduleUUID.

	   Snapshot copy policy schedule ID
	*/
	ScheduleUUIDPathParameter string

	/* SnapshotPolicyUUID.

	   Snapshot copy policy UUID
	*/
	SnapshotPolicyUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot policy schedule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleGetParams) WithDefaults() *SnapshotPolicyScheduleGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot policy schedule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) WithTimeout(timeout time.Duration) *SnapshotPolicyScheduleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) WithContext(ctx context.Context) *SnapshotPolicyScheduleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) WithHTTPClient(client *http.Client) *SnapshotPolicyScheduleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) WithFieldsQueryParameter(fields []string) *SnapshotPolicyScheduleGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithScheduleUUIDPathParameter adds the scheduleUUID to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) WithScheduleUUIDPathParameter(scheduleUUID string) *SnapshotPolicyScheduleGetParams {
	o.SetScheduleUUIDPathParameter(scheduleUUID)
	return o
}

// SetScheduleUUIDPathParameter adds the scheduleUuid to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) SetScheduleUUIDPathParameter(scheduleUUID string) {
	o.ScheduleUUIDPathParameter = scheduleUUID
}

// WithSnapshotPolicyUUIDPathParameter adds the snapshotPolicyUUID to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) WithSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID string) *SnapshotPolicyScheduleGetParams {
	o.SetSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID)
	return o
}

// SetSnapshotPolicyUUIDPathParameter adds the snapshotPolicyUuid to the snapshot policy schedule get params
func (o *SnapshotPolicyScheduleGetParams) SetSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID string) {
	o.SnapshotPolicyUUIDPathParameter = snapshotPolicyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotPolicyScheduleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param schedule.uuid
	if err := r.SetPathParam("schedule.uuid", o.ScheduleUUIDPathParameter); err != nil {
		return err
	}

	// path param snapshot_policy.uuid
	if err := r.SetPathParam("snapshot_policy.uuid", o.SnapshotPolicyUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnapshotPolicyScheduleGet binds the parameter fields
func (o *SnapshotPolicyScheduleGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
