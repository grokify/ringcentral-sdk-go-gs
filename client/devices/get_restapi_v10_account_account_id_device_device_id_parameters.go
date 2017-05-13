package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRestapiV10AccountAccountIDDeviceDeviceIDParams creates a new GetRestapiV10AccountAccountIDDeviceDeviceIDParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDDeviceDeviceIDParams() *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDDeviceDeviceIDParams{
		AccountID: accountIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDDeviceDeviceIDParamsWithTimeout creates a new GetRestapiV10AccountAccountIDDeviceDeviceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDDeviceDeviceIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDDeviceDeviceIDParams{
		AccountID: accountIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDDeviceDeviceIDParamsWithContext creates a new GetRestapiV10AccountAccountIDDeviceDeviceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDDeviceDeviceIDParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDDeviceDeviceIDParams{
		AccountID: accountIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDDeviceDeviceIDParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDDeviceDeviceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDDeviceDeviceIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDDeviceDeviceIDParams{
		AccountID:  accountIdDefault,
		HTTPClient: client,
	}
}

/*GetRestapiV10AccountAccountIDDeviceDeviceIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID device device ID operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDDeviceDeviceIDParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*DeviceID
	  Internal identifier of a device

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithDeviceID adds the deviceID to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) WithDeviceID(deviceID string) *GetRestapiV10AccountAccountIDDeviceDeviceIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get restapi v10 account account ID device device ID params
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDDeviceDeviceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
