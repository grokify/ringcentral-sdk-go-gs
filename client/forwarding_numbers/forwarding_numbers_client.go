package forwarding_numbers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new forwarding numbers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for forwarding numbers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumber Get Forwarding Numbers
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumber(params *GetRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumberParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumberParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumber",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PostRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumber Add New Forwarding Number
*/
func (a *Client) PostRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumber(params *PostRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumberParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumberParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumber",
		Method:             "POST",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRestapiV10AccountAccountIDExtensionExtensionIDForwardingNumberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}