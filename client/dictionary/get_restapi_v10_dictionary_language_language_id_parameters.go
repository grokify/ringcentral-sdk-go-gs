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

// NewGetRestapiV10DictionaryLanguageLanguageIDParams creates a new GetRestapiV10DictionaryLanguageLanguageIDParams object
// with the default values initialized.
func NewGetRestapiV10DictionaryLanguageLanguageIDParams() *GetRestapiV10DictionaryLanguageLanguageIDParams {
	var ()
	return &GetRestapiV10DictionaryLanguageLanguageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestapiV10DictionaryLanguageLanguageIDParamsWithTimeout creates a new GetRestapiV10DictionaryLanguageLanguageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestapiV10DictionaryLanguageLanguageIDParamsWithTimeout(timeout time.Duration) *GetRestapiV10DictionaryLanguageLanguageIDParams {
	var ()
	return &GetRestapiV10DictionaryLanguageLanguageIDParams{

		timeout: timeout,
	}
}

// NewGetRestapiV10DictionaryLanguageLanguageIDParamsWithContext creates a new GetRestapiV10DictionaryLanguageLanguageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestapiV10DictionaryLanguageLanguageIDParamsWithContext(ctx context.Context) *GetRestapiV10DictionaryLanguageLanguageIDParams {
	var ()
	return &GetRestapiV10DictionaryLanguageLanguageIDParams{

		Context: ctx,
	}
}

// NewGetRestapiV10DictionaryLanguageLanguageIDParamsWithHTTPClient creates a new GetRestapiV10DictionaryLanguageLanguageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestapiV10DictionaryLanguageLanguageIDParamsWithHTTPClient(client *http.Client) *GetRestapiV10DictionaryLanguageLanguageIDParams {
	var ()
	return &GetRestapiV10DictionaryLanguageLanguageIDParams{
		HTTPClient: client,
	}
}

/*GetRestapiV10DictionaryLanguageLanguageIDParams contains all the parameters to send to the API endpoint
for the get restapi v10 dictionary language language ID operation typically these are written to a http.Request
*/
type GetRestapiV10DictionaryLanguageLanguageIDParams struct {

	/*LanguageID
	  Internal identifier of a language

	*/
	LanguageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) WithTimeout(timeout time.Duration) *GetRestapiV10DictionaryLanguageLanguageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) WithContext(ctx context.Context) *GetRestapiV10DictionaryLanguageLanguageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) WithHTTPClient(client *http.Client) *GetRestapiV10DictionaryLanguageLanguageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguageID adds the languageID to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) WithLanguageID(languageID string) *GetRestapiV10DictionaryLanguageLanguageIDParams {
	o.SetLanguageID(languageID)
	return o
}

// SetLanguageID adds the languageId to the get restapi v10 dictionary language language ID params
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) SetLanguageID(languageID string) {
	o.LanguageID = languageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestapiV10DictionaryLanguageLanguageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param languageId
	if err := r.SetPathParam("languageId", o.LanguageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
