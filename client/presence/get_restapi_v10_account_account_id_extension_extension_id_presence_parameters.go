package presence

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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID presence operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams struct {

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

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID presence params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDPresenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
