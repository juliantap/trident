// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewUnixGroupUsersGetParams creates a new UnixGroupUsersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupUsersGetParams() *UnixGroupUsersGetParams {
	return &UnixGroupUsersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupUsersGetParamsWithTimeout creates a new UnixGroupUsersGetParams object
// with the ability to set a timeout on a request.
func NewUnixGroupUsersGetParamsWithTimeout(timeout time.Duration) *UnixGroupUsersGetParams {
	return &UnixGroupUsersGetParams{
		timeout: timeout,
	}
}

// NewUnixGroupUsersGetParamsWithContext creates a new UnixGroupUsersGetParams object
// with the ability to set a context for a request.
func NewUnixGroupUsersGetParamsWithContext(ctx context.Context) *UnixGroupUsersGetParams {
	return &UnixGroupUsersGetParams{
		Context: ctx,
	}
}

// NewUnixGroupUsersGetParamsWithHTTPClient creates a new UnixGroupUsersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupUsersGetParamsWithHTTPClient(client *http.Client) *UnixGroupUsersGetParams {
	return &UnixGroupUsersGetParams{
		HTTPClient: client,
	}
}

/* UnixGroupUsersGetParams contains all the parameters to send to the API endpoint
   for the unix group users get operation.

   Typically these are written to a http.Request.
*/
type UnixGroupUsersGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Name.

	   UNIX user to be retrieved from the group.
	*/
	NamePathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	/* UnixGroupName.

	   UNIX group name.
	*/
	UnixGroupNamePathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix group users get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUsersGetParams) WithDefaults() *UnixGroupUsersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group users get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUsersGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unix group users get params
func (o *UnixGroupUsersGetParams) WithTimeout(timeout time.Duration) *UnixGroupUsersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group users get params
func (o *UnixGroupUsersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group users get params
func (o *UnixGroupUsersGetParams) WithContext(ctx context.Context) *UnixGroupUsersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group users get params
func (o *UnixGroupUsersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group users get params
func (o *UnixGroupUsersGetParams) WithHTTPClient(client *http.Client) *UnixGroupUsersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group users get params
func (o *UnixGroupUsersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the unix group users get params
func (o *UnixGroupUsersGetParams) WithFieldsQueryParameter(fields []string) *UnixGroupUsersGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the unix group users get params
func (o *UnixGroupUsersGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNamePathParameter adds the name to the unix group users get params
func (o *UnixGroupUsersGetParams) WithNamePathParameter(name string) *UnixGroupUsersGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the unix group users get params
func (o *UnixGroupUsersGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSVMUUIDPathParameter adds the svmUUID to the unix group users get params
func (o *UnixGroupUsersGetParams) WithSVMUUIDPathParameter(svmUUID string) *UnixGroupUsersGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the unix group users get params
func (o *UnixGroupUsersGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithUnixGroupNamePathParameter adds the unixGroupName to the unix group users get params
func (o *UnixGroupUsersGetParams) WithUnixGroupNamePathParameter(unixGroupName string) *UnixGroupUsersGetParams {
	o.SetUnixGroupNamePathParameter(unixGroupName)
	return o
}

// SetUnixGroupNamePathParameter adds the unixGroupName to the unix group users get params
func (o *UnixGroupUsersGetParams) SetUnixGroupNamePathParameter(unixGroupName string) {
	o.UnixGroupNamePathParameter = unixGroupName
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupUsersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	// path param unix_group.name
	if err := r.SetPathParam("unix_group.name", o.UnixGroupNamePathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUnixGroupUsersGet binds the parameter fields
func (o *UnixGroupUsersGetParams) bindParamFields(formats strfmt.Registry) []string {
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
