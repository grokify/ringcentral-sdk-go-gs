package answering_rules

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

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams object
// with the default values initialized.
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams() *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParamsWithTimeout creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParamsWithTimeout(timeout time.Duration) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParamsWithContext creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParamsWithContext(ctx context.Context) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParamsWithHTTPClient creates a new PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParamsWithHTTPClient(client *http.Client) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams contains all the parameters to send to the API endpoint
for the post restapi v10 account account ID extension extension ID greeting operation typically these are written to a http.Request
*/
type PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*Body*/
	Body PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingBody
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) WithTimeout(timeout time.Duration) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) WithContext(ctx context.Context) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) WithHTTPClient(client *http.Client) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) WithAccountID(accountID string) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBody adds the body to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) WithBody(body PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingBody) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) SetBody(body PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingBody) {
	o.Body = body
}

// WithExtensionID adds the extensionID to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) WithExtensionID(extensionID string) *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the post restapi v10 account account ID extension extension ID greeting params
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDGreetingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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