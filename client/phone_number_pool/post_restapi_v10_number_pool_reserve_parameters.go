package phone_number_pool

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

// NewPostRestapiV10NumberPoolReserveParams creates a new PostRestapiV10NumberPoolReserveParams object
// with the default values initialized.
func NewPostRestapiV10NumberPoolReserveParams() *PostRestapiV10NumberPoolReserveParams {
	var ()
	return &PostRestapiV10NumberPoolReserveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRestapiV10NumberPoolReserveParamsWithTimeout creates a new PostRestapiV10NumberPoolReserveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRestapiV10NumberPoolReserveParamsWithTimeout(timeout time.Duration) *PostRestapiV10NumberPoolReserveParams {
	var ()
	return &PostRestapiV10NumberPoolReserveParams{

		timeout: timeout,
	}
}

// NewPostRestapiV10NumberPoolReserveParamsWithContext creates a new PostRestapiV10NumberPoolReserveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRestapiV10NumberPoolReserveParamsWithContext(ctx context.Context) *PostRestapiV10NumberPoolReserveParams {
	var ()
	return &PostRestapiV10NumberPoolReserveParams{

		Context: ctx,
	}
}

// NewPostRestapiV10NumberPoolReserveParamsWithHTTPClient creates a new PostRestapiV10NumberPoolReserveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRestapiV10NumberPoolReserveParamsWithHTTPClient(client *http.Client) *PostRestapiV10NumberPoolReserveParams {
	var ()
	return &PostRestapiV10NumberPoolReserveParams{
		HTTPClient: client,
	}
}

/*PostRestapiV10NumberPoolReserveParams contains all the parameters to send to the API endpoint
for the post restapi v10 number pool reserve operation typically these are written to a http.Request
*/
type PostRestapiV10NumberPoolReserveParams struct {

	/*Body*/
	Body PostRestapiV10NumberPoolReserveBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) WithTimeout(timeout time.Duration) *PostRestapiV10NumberPoolReserveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) WithContext(ctx context.Context) *PostRestapiV10NumberPoolReserveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) WithHTTPClient(client *http.Client) *PostRestapiV10NumberPoolReserveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) WithBody(body PostRestapiV10NumberPoolReserveBody) *PostRestapiV10NumberPoolReserveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post restapi v10 number pool reserve params
func (o *PostRestapiV10NumberPoolReserveParams) SetBody(body PostRestapiV10NumberPoolReserveBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRestapiV10NumberPoolReserveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
