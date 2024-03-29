package conferencing

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

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams object
// with the default values initialized.
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams() *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParamsWithTimeout creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParamsWithTimeout(timeout time.Duration) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParamsWithContext creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParamsWithContext(ctx context.Context) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParamsWithHTTPClient creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParamsWithHTTPClient(client *http.Client) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams contains all the parameters to send to the API endpoint
for the put restapi v10 account account ID extension extension ID conferencing operation typically these are written to a http.Request
*/
type PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*Body*/
	Body PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingBody
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) WithTimeout(timeout time.Duration) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) WithContext(ctx context.Context) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) WithHTTPClient(client *http.Client) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) WithAccountID(accountID string) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBody adds the body to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) WithBody(body PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingBody) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) SetBody(body PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingBody) {
	o.Body = body
}

// WithExtensionID adds the extensionID to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) WithExtensionID(extensionID string) *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the put restapi v10 account account ID extension extension ID conferencing params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDConferencingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
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
