package meetings

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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID meeting operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams struct {

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

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID meeting params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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