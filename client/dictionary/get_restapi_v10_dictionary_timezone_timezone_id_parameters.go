package dictionary

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

// NewGetRestapiV10DictionaryTimezoneTimezoneIDParams creates a new GetRestapiV10DictionaryTimezoneTimezoneIDParams object
// with the default values initialized.
func NewGetRestapiV10DictionaryTimezoneTimezoneIDParams() *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	var ()
	return &GetRestapiV10DictionaryTimezoneTimezoneIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10DictionaryTimezoneTimezoneIDParamsWithTimeout creates a new GetRestapiV10DictionaryTimezoneTimezoneIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10DictionaryTimezoneTimezoneIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	var ()
	return &GetRestapiV10DictionaryTimezoneTimezoneIDParams{

		timeout: timeout,
	}
}

// NewGetRestapiV10DictionaryTimezoneTimezoneIDParamsWithContext creates a new GetRestapiV10DictionaryTimezoneTimezoneIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10DictionaryTimezoneTimezoneIDParamsWithContext(ctx context.Context) *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	var ()
	return &GetRestapiV10DictionaryTimezoneTimezoneIDParams{

		Context: ctx,
	}
}

// NewGetRestapiV10DictionaryTimezoneTimezoneIDParamsWithHTTPClient creates a new GetRestapiV10DictionaryTimezoneTimezoneIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10DictionaryTimezoneTimezoneIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	var ()
	return &GetRestapiV10DictionaryTimezoneTimezoneIDParams{
		HTTPClient: client,
	}
}

/*GetRestapiV10DictionaryTimezoneTimezoneIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 dictionary timezone timezone ID operation typically these are written to a http.Request
*/
type GetRestapiV10DictionaryTimezoneTimezoneIDParams struct {

	/*TimezoneID
	  Internal identifier of a timezone

	*/
	TimezoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) WithContext(ctx context.Context) *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimezoneID adds the timezoneID to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) WithTimezoneID(timezoneID string) *GetRestapiV10DictionaryTimezoneTimezoneIDParams {
	o.SetTimezoneID(timezoneID)
	return o
}

// SetTimezoneID adds the timezoneId to the get restapi v10 dictionary timezone timezone ID params
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) SetTimezoneID(timezoneID string) {
	o.TimezoneID = timezoneID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10DictionaryTimezoneTimezoneIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param timezoneId
	if err := r.SetPathParam("timezoneId", o.TimezoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}