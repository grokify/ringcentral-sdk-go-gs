package account

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

// NewGetRestapiV10AccountAccountIDParams creates a new GetRestapiV10AccountAccountIDParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDParams() *GetRestapiV10AccountAccountIDParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDParams{
		AccountID: accountIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDParamsWithTimeout creates a new GetRestapiV10AccountAccountIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDParams{
		AccountID: accountIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDParamsWithContext creates a new GetRestapiV10AccountAccountIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDParams{
		AccountID: accountIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDParams{
		AccountID:  accountIdDefault,
		HTTPClient: client,
	}
}

/*GetRestapiV10AccountAccountIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID params
func (o *GetRestapiV10AccountAccountIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
