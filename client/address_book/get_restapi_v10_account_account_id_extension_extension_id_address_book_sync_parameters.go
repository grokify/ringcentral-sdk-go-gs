package address_book

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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID address book sync operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string
	/*PageID
	  Internal identifier of a page. It can be obtained from the 'nextPageId' parameter passed in response body

	*/
	PageID *int64
	/*PerPage
	  Number of records per page to be returned. The max number of records is 250, which is also the default. For FSync — if the number of records exceeds the parameter value (either specified or default), all of the pages can be retrieved in several requests. For ISync — if the number of records exceeds the page size, the number of incoming changes to this number is limited

	*/
	PerPage *int64
	/*SyncToken
	  Value of syncToken property of the last sync request response

	*/
	SyncToken *string
	/*SyncType
	  Type of synchronization. The default value is 'FSync'

	*/
	SyncType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithPageID adds the pageID to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithPageID(pageID *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetPageID(pageID)
	return o
}

// SetPageID adds the pageId to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetPageID(pageID *int64) {
	o.PageID = pageID
}

// WithPerPage adds the perPage to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithPerPage(perPage *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSyncToken adds the syncToken to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithSyncToken(syncToken *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetSyncToken(syncToken)
	return o
}

// SetSyncToken adds the syncToken to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetSyncToken(syncToken *string) {
	o.SyncToken = syncToken
}

// WithSyncType adds the syncType to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WithSyncType(syncType *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams {
	o.SetSyncType(syncType)
	return o
}

// SetSyncType adds the syncType to the get restapi v10 account account ID extension extension ID address book sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) SetSyncType(syncType *string) {
	o.SyncType = syncType
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PageID != nil {

		// query param pageId
		var qrPageID int64
		if o.PageID != nil {
			qrPageID = *o.PageID
		}
		qPageID := swag.FormatInt64(qrPageID)
		if qPageID != "" {
			if err := r.SetQueryParam("pageId", qPageID); err != nil {
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

	if o.SyncToken != nil {

		// query param syncToken
		var qrSyncToken string
		if o.SyncToken != nil {
			qrSyncToken = *o.SyncToken
		}
		qSyncToken := qrSyncToken
		if qSyncToken != "" {
			if err := r.SetQueryParam("syncToken", qSyncToken); err != nil {
				return err
			}
		}

	}

	if o.SyncType != nil {

		// query param syncType
		var qrSyncType string
		if o.SyncType != nil {
			qrSyncType = *o.SyncType
		}
		qSyncType := qrSyncType
		if qSyncType != "" {
			if err := r.SetQueryParam("syncType", qSyncType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
