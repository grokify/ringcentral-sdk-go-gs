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

// NewGetRestapiV10AccountAccountIDServiceInfoParams creates a new GetRestapiV10AccountAccountIDServiceInfoParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDServiceInfoParams() *GetRestapiV10AccountAccountIDServiceInfoParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDServiceInfoParams{
		AccountID: accountIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDServiceInfoParamsWithTimeout creates a new GetRestapiV10AccountAccountIDServiceInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDServiceInfoParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDServiceInfoParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDServiceInfoParams{
		AccountID: accountIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDServiceInfoParamsWithContext creates a new GetRestapiV10AccountAccountIDServiceInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDServiceInfoParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDServiceInfoParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDServiceInfoParams{
		AccountID: accountIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDServiceInfoParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDServiceInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDServiceInfoParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDServiceInfoParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDServiceInfoParams{
		AccountID:  accountIdDefault,
		HTTPClient: client,
	}
}

/*GetRestapiV10AccountAccountIDServiceInfoParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID service info operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDServiceInfoParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDServiceInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDServiceInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDServiceInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDServiceInfoParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID service info params
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDServiceInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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