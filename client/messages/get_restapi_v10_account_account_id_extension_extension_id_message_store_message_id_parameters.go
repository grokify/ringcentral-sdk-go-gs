package messages

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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID message store message ID operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string
	/*MessageID
	  Internal identifier of a message

	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithMessageID adds the messageID to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithMessageID(messageID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the get restapi v10 account account ID extension extension ID message store message ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param messageId
	if err := r.SetPathParam("messageId", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
