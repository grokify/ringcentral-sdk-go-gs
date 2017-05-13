package blocked_numbers

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

	"github.com/grokify/go-ringcentral/models"
)

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams object
// with the default values initialized.
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams() *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParamsWithTimeout creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParamsWithTimeout(timeout time.Duration) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParamsWithContext creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParamsWithContext(ctx context.Context) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParamsWithHTTPClient creates a new PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParamsWithHTTPClient(client *http.Client) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams contains all the parameters to send to the API endpoint
for the put restapi v10 account account ID extension extension ID blocked number blocked number ID operation typically these are written to a http.Request
*/
type PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*BlockedNumberID
	  Internal identifier of a blocked number list entry

	*/
	BlockedNumberID string
	/*Body*/
	Body *models.BlockedNumberInfo
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WithTimeout(timeout time.Duration) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WithContext(ctx context.Context) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WithHTTPClient(client *http.Client) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WithAccountID(accountID string) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBlockedNumberID adds the blockedNumberID to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WithBlockedNumberID(blockedNumberID string) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	o.SetBlockedNumberID(blockedNumberID)
	return o
}

// SetBlockedNumberID adds the blockedNumberId to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) SetBlockedNumberID(blockedNumberID string) {
	o.BlockedNumberID = blockedNumberID
}

// WithBody adds the body to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WithBody(body *models.BlockedNumberInfo) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) SetBody(body *models.BlockedNumberInfo) {
	o.Body = body
}

// WithExtensionID adds the extensionID to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WithExtensionID(extensionID string) *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the put restapi v10 account account ID extension extension ID blocked number blocked number ID params
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDBlockedNumberBlockedNumberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	// path param blockedNumberId
	if err := r.SetPathParam("blockedNumberId", o.BlockedNumberID); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models.BlockedNumberInfo)
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