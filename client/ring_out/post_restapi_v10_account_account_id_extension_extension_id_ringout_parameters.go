package ring_out

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

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams object
// with the default values initialized.
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams() *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParamsWithTimeout creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParamsWithTimeout(timeout time.Duration) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParamsWithContext creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParamsWithContext(ctx context.Context) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParamsWithHTTPClient creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParamsWithHTTPClient(client *http.Client) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams contains all the parameters to send to the API endpoint
for the post restapi v10 account account ID extension extension ID ringout operation typically these are written to a http.Request
*/
type PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*Body*/
	Body PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutBody
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) WithTimeout(timeout time.Duration) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) WithContext(ctx context.Context) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) WithHTTPClient(client *http.Client) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) WithAccountID(accountID string) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBody adds the body to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) WithBody(body PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutBody) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) SetBody(body PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutBody) {
	o.Body = body
}

// WithExtensionID adds the extensionID to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) WithExtensionID(extensionID string) *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the post restapi v10 account account ID extension extension ID ringout params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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