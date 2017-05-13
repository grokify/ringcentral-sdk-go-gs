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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams creates a new DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized.
func NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams() *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithTimeout creates a new DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithTimeout(timeout time.Duration) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithContext creates a new DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithContext(ctx context.Context) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithHTTPClient creates a new DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParamsWithHTTPClient(client *http.Client) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams contains all the parameters to send to the API endpoint
for the delete restapi v10 account account ID extension extension ID message store message ID operation typically these are written to a http.Request
*/
type DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ConversationID
	  Internal identifier of a message thread

	*/
	ConversationID *int64
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string
	/*MessageID
	  Internal identifier of a message

	*/
	MessageID string
	/*Purge
	  If the value is 'True', then the message is purged immediately with all the attachments. The default value is 'False'

	*/
	Purge *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithTimeout(timeout time.Duration) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithContext(ctx context.Context) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithHTTPClient(client *http.Client) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithAccountID(accountID string) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithConversationID adds the conversationID to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithConversationID(conversationID *int64) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetConversationID(conversationID *int64) {
	o.ConversationID = conversationID
}

// WithExtensionID adds the extensionID to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithExtensionID(extensionID string) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithMessageID adds the messageID to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithMessageID(messageID string) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WithPurge adds the purge to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WithPurge(purge *bool) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams {
	o.SetPurge(purge)
	return o
}

// SetPurge adds the purge to the delete restapi v10 account account ID extension extension ID message store message ID params
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) SetPurge(purge *bool) {
	o.Purge = purge
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if o.ConversationID != nil {

		// query param conversationId
		var qrConversationID int64
		if o.ConversationID != nil {
			qrConversationID = *o.ConversationID
		}
		qConversationID := swag.FormatInt64(qrConversationID)
		if qConversationID != "" {
			if err := r.SetQueryParam("conversationId", qConversationID); err != nil {
				return err
			}
		}

	}

	// path param extensionId
	if err := r.SetPathParam("extensionId", o.ExtensionID); err != nil {
		return err
	}

	// path param messageId
	if err := r.SetPathParam("messageId", o.MessageID); err != nil {
		return err
	}

	if o.Purge != nil {

		// query param purge
		var qrPurge bool
		if o.Purge != nil {
			qrPurge = *o.Purge
		}
		qPurge := swag.FormatBool(qrPurge)
		if qPurge != "" {
			if err := r.SetQueryParam("purge", qPurge); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
