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

// NewGetRestapiV10AccountAccountIDExtensionParams creates a new GetRestapiV10AccountAccountIDExtensionParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionParams() *GetRestapiV10AccountAccountIDExtensionParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionParams{
		AccountID: accountIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionParams {
	var (
		accountIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionParams{
		AccountID: accountIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionParams{
		AccountID: accountIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionParams {
	var (
		accountIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionParams{
		AccountID:  accountIdDefault,
		HTTPClient: client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*Page
	  Indicates the page number to retrieve. Only positive number values are allowed. Default value is '1'

	*/
	Page *int64
	/*PerPage
	  Indicates the page size (number of items). If not specified, the value is '100' by default.

	*/
	PerPage *int64
	/*Status
	  Extension current state. Multiple values are supported. If 'Unassigned' is specified, then extensions without extensionNumber are returned. If not specified, then all extensions are returned

	*/
	Status *string
	/*Type
	  Extension type. Multiple values are supported

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithPage adds the page to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithPage(page *int64) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithPerPage(perPage *int64) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStatus adds the status to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithStatus(status *string) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetStatus(status *string) {
	o.Status = status
}

// WithType adds the typeVar to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) WithType(typeVar *string) *GetRestapiV10AccountAccountIDExtensionParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get restapi v10 account account ID extension params
func (o *GetRestapiV10AccountAccountIDExtensionParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
