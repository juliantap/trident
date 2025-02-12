// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityKeyManagerKeyServersCollectionGetParams creates a new SecurityKeyManagerKeyServersCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerKeyServersCollectionGetParams() *SecurityKeyManagerKeyServersCollectionGetParams {
	return &SecurityKeyManagerKeyServersCollectionGetParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerKeyServersCollectionGetParamsWithTimeout creates a new SecurityKeyManagerKeyServersCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerKeyServersCollectionGetParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersCollectionGetParams {
	return &SecurityKeyManagerKeyServersCollectionGetParams{
		requestTimeout: timeout,
	}
}

// NewSecurityKeyManagerKeyServersCollectionGetParamsWithContext creates a new SecurityKeyManagerKeyServersCollectionGetParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerKeyServersCollectionGetParamsWithContext(ctx context.Context) *SecurityKeyManagerKeyServersCollectionGetParams {
	return &SecurityKeyManagerKeyServersCollectionGetParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerKeyServersCollectionGetParamsWithHTTPClient creates a new SecurityKeyManagerKeyServersCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerKeyServersCollectionGetParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersCollectionGetParams {
	return &SecurityKeyManagerKeyServersCollectionGetParams{
		HTTPClient: client,
	}
}

/* SecurityKeyManagerKeyServersCollectionGetParams contains all the parameters to send to the API endpoint
   for the security key manager key servers collection get operation.

   Typically these are written to a http.Request.
*/
type SecurityKeyManagerKeyServersCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* SecondaryKeyServers.

	   Filter by secondary_key_servers
	*/
	SecondaryKeyServersQueryParameter *string

	/* Server.

	   Filter by server
	*/
	ServerQueryParameter *string

	/* Timeout.

	   Filter by timeout
	*/
	TimeoutQueryParameter *int64

	/* Username.

	   Filter by username
	*/
	UsernameQueryParameter *string

	/* UUID.

	   External key manager UUID
	*/
	UUIDPathParameter string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the security key manager key servers collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithDefaults() *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager key servers collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SecurityKeyManagerKeyServersCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.requestTimeout = o.requestTimeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithRequestTimeout adds the timeout to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithRequestTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithContext(ctx context.Context) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithFieldsQueryParameter(fields []string) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSecondaryKeyServersQueryParameter adds the secondaryKeyServers to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithSecondaryKeyServersQueryParameter(secondaryKeyServers *string) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetSecondaryKeyServersQueryParameter(secondaryKeyServers)
	return o
}

// SetSecondaryKeyServersQueryParameter adds the secondaryKeyServers to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetSecondaryKeyServersQueryParameter(secondaryKeyServers *string) {
	o.SecondaryKeyServersQueryParameter = secondaryKeyServers
}

// WithServerQueryParameter adds the server to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithServerQueryParameter(server *string) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetServerQueryParameter(server)
	return o
}

// SetServerQueryParameter adds the server to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetServerQueryParameter(server *string) {
	o.ServerQueryParameter = server
}

// WithTimeoutQueryParameter adds the timeout to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithTimeoutQueryParameter(timeout *int64) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetTimeoutQueryParameter(timeout)
	return o
}

// SetTimeoutQueryParameter adds the timeout to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetTimeoutQueryParameter(timeout *int64) {
	o.TimeoutQueryParameter = timeout
}

// WithUsernameQueryParameter adds the username to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithUsernameQueryParameter(username *string) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetUsernameQueryParameter(username)
	return o
}

// SetUsernameQueryParameter adds the username to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetUsernameQueryParameter(username *string) {
	o.UsernameQueryParameter = username
}

// WithUUIDPathParameter adds the uuid to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WithUUIDPathParameter(uuid string) *SecurityKeyManagerKeyServersCollectionGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the security key manager key servers collection get params
func (o *SecurityKeyManagerKeyServersCollectionGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerKeyServersCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
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

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SecondaryKeyServersQueryParameter != nil {

		// query param secondary_key_servers
		var qrSecondaryKeyServers string

		if o.SecondaryKeyServersQueryParameter != nil {
			qrSecondaryKeyServers = *o.SecondaryKeyServersQueryParameter
		}
		qSecondaryKeyServers := qrSecondaryKeyServers
		if qSecondaryKeyServers != "" {

			if err := r.SetQueryParam("secondary_key_servers", qSecondaryKeyServers); err != nil {
				return err
			}
		}
	}

	if o.ServerQueryParameter != nil {

		// query param server
		var qrServer string

		if o.ServerQueryParameter != nil {
			qrServer = *o.ServerQueryParameter
		}
		qServer := qrServer
		if qServer != "" {

			if err := r.SetQueryParam("server", qServer); err != nil {
				return err
			}
		}
	}

	if o.TimeoutQueryParameter != nil {

		// query param timeout
		var qrTimeout int64

		if o.TimeoutQueryParameter != nil {
			qrTimeout = *o.TimeoutQueryParameter
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {

			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}
	}

	if o.UsernameQueryParameter != nil {

		// query param username
		var qrUsername string

		if o.UsernameQueryParameter != nil {
			qrUsername = *o.UsernameQueryParameter
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
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

// bindParamSecurityKeyManagerKeyServersCollectionGet binds the parameter fields
func (o *SecurityKeyManagerKeyServersCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSecurityKeyManagerKeyServersCollectionGet binds the parameter order_by
func (o *SecurityKeyManagerKeyServersCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
