package meetings

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

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams object
// with the default values initialized.
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams() *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParamsWithTimeout creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	var (
		accountIDDefault   = string("~")
		extensionIDDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams{
		AccountID:   accountIDDefault,
		ExtensionID: extensionIDDefault,

		timeout: timeout,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParamsWithContext creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParamsWithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,

		Context: ctx,
	}
}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParamsWithHTTPClient creates a new GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	var (
		accountIdDefault   = string("~")
		extensionIdDefault = string("~")
	)
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams{
		AccountID:   accountIdDefault,
		ExtensionID: extensionIdDefault,
		HTTPClient:  client,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 account account ID extension extension ID meeting meeting ID operation typically these are written to a http.Request
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams struct {

	/*AccountID
	  Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session

	*/
	AccountID string
	/*ExtensionID
	  Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session

	*/
	ExtensionID string
	/*MeetingID
	  Internal identifier of a meeting

	*/
	MeetingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) WithContext(ctx context.Context) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) WithAccountID(accountID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithExtensionID adds the extensionID to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) WithExtensionID(extensionID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithMeetingID adds the meetingID to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) WithMeetingID(meetingID string) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams {
	o.SetMeetingID(meetingID)
	return o
}

// SetMeetingID adds the meetingId to the get restapi v10 account account ID extension extension ID meeting meeting ID params
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) SetMeetingID(meetingID string) {
	o.MeetingID = meetingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param meetingId
	if err := r.SetPathParam("meetingId", o.MeetingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
