package client_application

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

// NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParams creates a new PutRestapiV10ClientInfoCustomDataCustomDataKeyParams object
// with the default values initialized.
func NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParams() *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	var ()
	return &PutRestapiV10ClientInfoCustomDataCustomDataKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParamsWithTimeout creates a new PutRestapiV10ClientInfoCustomDataCustomDataKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParamsWithTimeout(timeout time.Duration) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	var ()
	return &PutRestapiV10ClientInfoCustomDataCustomDataKeyParams{

		timeout: timeout,
	}
}

// NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParamsWithContext creates a new PutRestapiV10ClientInfoCustomDataCustomDataKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParamsWithContext(ctx context.Context) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	var ()
	return &PutRestapiV10ClientInfoCustomDataCustomDataKeyParams{

		Context: ctx,
	}
}

// NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParamsWithHTTPClient creates a new PutRestapiV10ClientInfoCustomDataCustomDataKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRestapiV10ClientInfoCustomDataCustomDataKeyParamsWithHTTPClient(client *http.Client) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	var ()
	return &PutRestapiV10ClientInfoCustomDataCustomDataKeyParams{
		HTTPClient: client,
	}
}

/*PutRestapiV10ClientInfoCustomDataCustomDataKeyParams contains all the parameters to send to the API endpoint
for the put restapi v10 client info custom data custom data key operation typically these are written to a http.Request
*/
type PutRestapiV10ClientInfoCustomDataCustomDataKeyParams struct {

	/*Body*/
	Body PutRestapiV10ClientInfoCustomDataCustomDataKeyBody
	/*CustomDataKey
	  Custom data access key. The number of unique custom data keys is limited to 100 keys per extension, summarized for all the applications. For example, if you have created 50 custom data keys under the Android mobile client application for the particular extension, then logged in the iOS application and created another 50 keys, the web client application won't be allowed to create any custom data key for that extension

	*/
	CustomDataKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) WithTimeout(timeout time.Duration) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) WithContext(ctx context.Context) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) WithHTTPClient(client *http.Client) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) WithBody(body PutRestapiV10ClientInfoCustomDataCustomDataKeyBody) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) SetBody(body PutRestapiV10ClientInfoCustomDataCustomDataKeyBody) {
	o.Body = body
}

// WithCustomDataKey adds the customDataKey to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) WithCustomDataKey(customDataKey string) *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams {
	o.SetCustomDataKey(customDataKey)
	return o
}

// SetCustomDataKey adds the customDataKey to the put restapi v10 client info custom data custom data key params
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) SetCustomDataKey(customDataKey string) {
	o.CustomDataKey = customDataKey
}

// WriteToRequest writes these params to a swagger request
func (o *PutRestapiV10ClientInfoCustomDataCustomDataKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param customDataKey
	if err := r.SetPathParam("customDataKey", o.CustomDataKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
