package ring_out

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ring out API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ring out API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID Cancel RingOut Call
*/
func (a *Client) DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID(params *DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID",
		Method:             "DELETE",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID Get RingOut Call Status
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID(params *GetRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDReader{formats: a.formats},
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
PostRestapiV10AccountAccountIDExtensionExtensionIDRingout Initiate RingOut Call
*/
func (a *Client) PostRestapiV10AccountAccountIDExtensionExtensionIDRingout(params *PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRestapiV10AccountAccountIDExtensionExtensionIDRingoutParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRestapiV10AccountAccountIDExtensionExtensionIDRingout",
		Method:             "POST",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRestapiV10AccountAccountIDExtensionExtensionIDRingoutReader{formats: a.formats},
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
