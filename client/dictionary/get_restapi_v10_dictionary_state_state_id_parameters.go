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

// NewGetRestapiV10DictionaryStateStateIDParams creates a new GetRestapiV10DictionaryStateStateIDParams object
// with the default values initialized.
func NewGetRestapiV10DictionaryStateStateIDParams() *GetRestapiV10DictionaryStateStateIDParams {
	var ()
	return &GetRestapiV10DictionaryStateStateIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10DictionaryStateStateIDParamsWithTimeout creates a new GetRestapiV10DictionaryStateStateIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10DictionaryStateStateIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10DictionaryStateStateIDParams {
	var ()
	return &GetRestapiV10DictionaryStateStateIDParams{

		timeout: timeout,
	}
}

// NewGetRestapiV10DictionaryStateStateIDParamsWithContext creates a new GetRestapiV10DictionaryStateStateIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10DictionaryStateStateIDParamsWithContext(ctx context.Context) *GetRestapiV10DictionaryStateStateIDParams {
	var ()
	return &GetRestapiV10DictionaryStateStateIDParams{

		Context: ctx,
	}
}

// NewGetRestapiV10DictionaryStateStateIDParamsWithHTTPClient creates a new GetRestapiV10DictionaryStateStateIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10DictionaryStateStateIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10DictionaryStateStateIDParams {
	var ()
	return &GetRestapiV10DictionaryStateStateIDParams{
		HTTPClient: client,
	}
}

/*GetRestapiV10DictionaryStateStateIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 dictionary state state ID operation typically these are written to a http.Request
*/
type GetRestapiV10DictionaryStateStateIDParams struct {

	/*StateID
	  Internal identifier of a state

	*/
	StateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10DictionaryStateStateIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) WithContext(ctx context.Context) *GetRestapiV10DictionaryStateStateIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10DictionaryStateStateIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStateID adds the stateID to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) WithStateID(stateID string) *GetRestapiV10DictionaryStateStateIDParams {
	o.SetStateID(stateID)
	return o
}

// SetStateID adds the stateId to the get restapi v10 dictionary state state ID params
func (o *GetRestapiV10DictionaryStateStateIDParams) SetStateID(stateID string) {
	o.StateID = stateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10DictionaryStateStateIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stateId
	if err := r.SetPathParam("stateId", o.StateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
