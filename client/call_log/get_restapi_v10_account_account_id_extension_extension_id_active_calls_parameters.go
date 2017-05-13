package call_log

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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID active calls operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*Direction
	  The direction for the result records. It is allowed to specify more than one direction. If not specified, both inbound and outbound records are returned. Multiple values are accepted

	*/
	Direction *string
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
	/*Type
	  Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithDirection adds the direction to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithDirection(direction *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithPage adds the page to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithPage(page *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithPerPage(perPage *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithType adds the typeVar to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WithType(typeVar *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get restapi v10 account account ID extension extension ID active calls params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string
		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {
			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}

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
