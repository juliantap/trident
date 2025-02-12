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

// NewShelfGetParams creates a new ShelfGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShelfGetParams() *ShelfGetParams {
	return &ShelfGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShelfGetParamsWithTimeout creates a new ShelfGetParams object
// with the ability to set a timeout on a request.
func NewShelfGetParamsWithTimeout(timeout time.Duration) *ShelfGetParams {
	return &ShelfGetParams{
		timeout: timeout,
	}
}

// NewShelfGetParamsWithContext creates a new ShelfGetParams object
// with the ability to set a context for a request.
func NewShelfGetParamsWithContext(ctx context.Context) *ShelfGetParams {
	return &ShelfGetParams{
		Context: ctx,
	}
}

// NewShelfGetParamsWithHTTPClient creates a new ShelfGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewShelfGetParamsWithHTTPClient(client *http.Client) *ShelfGetParams {
	return &ShelfGetParams{
		HTTPClient: client,
	}
}

/* ShelfGetParams contains all the parameters to send to the API endpoint
   for the shelf get operation.

   Typically these are written to a http.Request.
*/
type ShelfGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UID.

	   Shelf UID
	*/
	UIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the shelf get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShelfGetParams) WithDefaults() *ShelfGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the shelf get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShelfGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the shelf get params
func (o *ShelfGetParams) WithTimeout(timeout time.Duration) *ShelfGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shelf get params
func (o *ShelfGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shelf get params
func (o *ShelfGetParams) WithContext(ctx context.Context) *ShelfGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shelf get params
func (o *ShelfGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shelf get params
func (o *ShelfGetParams) WithHTTPClient(client *http.Client) *ShelfGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shelf get params
func (o *ShelfGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the shelf get params
func (o *ShelfGetParams) WithFieldsQueryParameter(fields []string) *ShelfGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the shelf get params
func (o *ShelfGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUIDPathParameter adds the uid to the shelf get params
func (o *ShelfGetParams) WithUIDPathParameter(uid string) *ShelfGetParams {
	o.SetUIDPathParameter(uid)
	return o
}

// SetUIDPathParameter adds the uid to the shelf get params
func (o *ShelfGetParams) SetUIDPathParameter(uid string) {
	o.UIDPathParameter = uid
}

// WriteToRequest writes these params to a swagger request
func (o *ShelfGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uid
	if err := r.SetPathParam("uid", o.UIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamShelfGet binds the parameter fields
func (o *ShelfGetParams) bindParamFields(formats strfmt.Registry) []string {
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
