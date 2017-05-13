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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID device operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID device params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	// path param extensionId
	if err := r.SetPathParam("extensionId", o.ExtensionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}