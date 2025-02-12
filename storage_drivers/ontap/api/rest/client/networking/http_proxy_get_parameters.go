// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewHTTPProxyGetParams creates a new HTTPProxyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHTTPProxyGetParams() *HTTPProxyGetParams {
	return &HTTPProxyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPProxyGetParamsWithTimeout creates a new HTTPProxyGetParams object
// with the ability to set a timeout on a request.
func NewHTTPProxyGetParamsWithTimeout(timeout time.Duration) *HTTPProxyGetParams {
	return &HTTPProxyGetParams{
		timeout: timeout,
	}
}

// NewHTTPProxyGetParamsWithContext creates a new HTTPProxyGetParams object
// with the ability to set a context for a request.
func NewHTTPProxyGetParamsWithContext(ctx context.Context) *HTTPProxyGetParams {
	return &HTTPProxyGetParams{
		Context: ctx,
	}
}

// NewHTTPProxyGetParamsWithHTTPClient creates a new HTTPProxyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewHTTPProxyGetParamsWithHTTPClient(client *http.Client) *HTTPProxyGetParams {
	return &HTTPProxyGetParams{
		HTTPClient: client,
	}
}

/* HTTPProxyGetParams contains all the parameters to send to the API endpoint
   for the http proxy get operation.

   Typically these are written to a http.Request.
*/
type HTTPProxyGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   HTTP proxy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the http proxy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyGetParams) WithDefaults() *HTTPProxyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the http proxy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the http proxy get params
func (o *HTTPProxyGetParams) WithTimeout(timeout time.Duration) *HTTPProxyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the http proxy get params
func (o *HTTPProxyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the http proxy get params
func (o *HTTPProxyGetParams) WithContext(ctx context.Context) *HTTPProxyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the http proxy get params
func (o *HTTPProxyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the http proxy get params
func (o *HTTPProxyGetParams) WithHTTPClient(client *http.Client) *HTTPProxyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the http proxy get params
func (o *HTTPProxyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the http proxy get params
func (o *HTTPProxyGetParams) WithFieldsQueryParameter(fields []string) *HTTPProxyGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the http proxy get params
func (o *HTTPProxyGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the http proxy get params
func (o *HTTPProxyGetParams) WithUUIDPathParameter(uuid string) *HTTPProxyGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the http proxy get params
func (o *HTTPProxyGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPProxyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamHTTPProxyGet binds the parameter fields
func (o *HTTPProxyGetParams) bindParamFields(formats strfmt.Registry) []string {
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
