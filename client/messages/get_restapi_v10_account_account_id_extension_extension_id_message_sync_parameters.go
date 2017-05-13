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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID message sync operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ConversationID
	  Conversation identifier for the resulting messages. Meaningful for SMS and Pager messages only.

	*/
	ConversationID *int64
	/*DateFrom
	  The start datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours

	*/
	DateFrom *strfmt.DateTime
	/*DateTo
	  The end datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time

	*/
	DateTo *strfmt.DateTime
	/*Direction
	  Direction for the resulting messages. If not specified, both inbound and outbound messages are returned. Multiple values are accepted

	*/
	Direction *string
	/*DistinctConversations
	  If 'True', then the latest messages per every conversation ID are returned

	*/
	DistinctConversations *bool
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string
	/*MessageType
	  Type for the resulting messages. If not specified, all types of messages are returned. Multiple values are accepted

	*/
	MessageType *string
	/*RecordCount
	  Limits the number of records to be returned (works in combination with dateFrom and dateTo if specified)

	*/
	RecordCount *int64
	/*SyncToken
	  Value of syncToken property of last sync request response

	*/
	SyncToken *string
	/*SyncType
	  Type of message synchronization

	*/
	SyncType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithConversationID adds the conversationID to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithConversationID(conversationID *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetConversationID(conversationID *int64) {
	o.ConversationID = conversationID
}

// WithDateFrom adds the dateFrom to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithDateFrom(dateFrom *strfmt.DateTime) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetDateFrom(dateFrom *strfmt.DateTime) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithDateTo(dateTo *strfmt.DateTime) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetDateTo(dateTo *strfmt.DateTime) {
	o.DateTo = dateTo
}

// WithDirection adds the direction to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithDirection(direction *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithDistinctConversations adds the distinctConversations to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithDistinctConversations(distinctConversations *bool) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetDistinctConversations(distinctConversations)
	return o
}

// SetDistinctConversations adds the distinctConversations to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetDistinctConversations(distinctConversations *bool) {
	o.DistinctConversations = distinctConversations
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithMessageType adds the messageType to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithMessageType(messageType *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetMessageType(messageType)
	return o
}

// SetMessageType adds the messageType to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetMessageType(messageType *string) {
	o.MessageType = messageType
}

// WithRecordCount adds the recordCount to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithRecordCount(recordCount *int64) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetRecordCount(recordCount)
	return o
}

// SetRecordCount adds the recordCount to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetRecordCount(recordCount *int64) {
	o.RecordCount = recordCount
}

// WithSyncToken adds the syncToken to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithSyncToken(syncToken *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetSyncToken(syncToken)
	return o
}

// SetSyncToken adds the syncToken to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetSyncToken(syncToken *string) {
	o.SyncToken = syncToken
}

// WithSyncType adds the syncType to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WithSyncType(syncType *string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams {
	o.SetSyncType(syncType)
	return o
}

// SetSyncType adds the syncType to the get restapi v10 account account ID extension extension ID message sync params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) SetSyncType(syncType *string) {
	o.SyncType = syncType
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMessageSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DateFrom != nil {

		// query param dateFrom
		var qrDateFrom strfmt.DateTime
		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := qrDateFrom.String()
		if qDateFrom != "" {
			if err := r.SetQueryParam("dateFrom", qDateFrom); err != nil {
				return err
			}
		}

	}

	if o.DateTo != nil {

		// query param dateTo
		var qrDateTo strfmt.DateTime
		if o.DateTo != nil {
			qrDateTo = *o.DateTo
		}
		qDateTo := qrDateTo.String()
		if qDateTo != "" {
			if err := r.SetQueryParam("dateTo", qDateTo); err != nil {
				return err
			}
		}

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

	if o.DistinctConversations != nil {

		// query param distinctConversations
		var qrDistinctConversations bool
		if o.DistinctConversations != nil {
			qrDistinctConversations = *o.DistinctConversations
		}
		qDistinctConversations := swag.FormatBool(qrDistinctConversations)
		if qDistinctConversations != "" {
			if err := r.SetQueryParam("distinctConversations", qDistinctConversations); err != nil {
				return err
			}
		}

	}

	// path param extensionId
	if err := r.SetPathParam("extensionId", o.ExtensionID); err != nil {
		return err
	}

	if o.MessageType != nil {

		// query param messageType
		var qrMessageType string
		if o.MessageType != nil {
			qrMessageType = *o.MessageType
		}
		qMessageType := qrMessageType
		if qMessageType != "" {
			if err := r.SetQueryParam("messageType", qMessageType); err != nil {
				return err
			}
		}

	}

	if o.RecordCount != nil {

		// query param recordCount
		var qrRecordCount int64
		if o.RecordCount != nil {
			qrRecordCount = *o.RecordCount
		}
		qRecordCount := swag.FormatInt64(qrRecordCount)
		if qRecordCount != "" {
			if err := r.SetQueryParam("recordCount", qRecordCount); err != nil {
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
