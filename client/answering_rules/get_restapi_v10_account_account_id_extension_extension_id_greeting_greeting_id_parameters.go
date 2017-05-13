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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID greeting greeting ID operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string
	/*GreetingID
	  Internal identifier of a greeting

	*/
	GreetingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithGreetingID adds the greetingID to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) WithGreetingID(greetingID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams {
	o.SetGreetingID(greetingID)
	return o
}

// SetGreetingID adds the greetingId to the get restapi v10 account account ID extension extension ID greeting greeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) SetGreetingID(greetingID string) {
	o.GreetingID = greetingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGreetingGreetingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param greetingId
	if err := r.SetPathParam("greetingId", o.GreetingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
