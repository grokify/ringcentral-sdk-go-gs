package extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDGrantParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID grant operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string
	/*Page
	  Indicates the page number to retrieve. Only positive number values are allowed. Default value is '1'

	*/
	Page *int64
	/*PerPage
	  Indicates the page size (number of items). If not specified, the value is '100' by default

	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithPage adds the page to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WithPage(page *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WithPerPage(perPage *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get restapi v10 account account ID extension extension ID grant params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDGrantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}