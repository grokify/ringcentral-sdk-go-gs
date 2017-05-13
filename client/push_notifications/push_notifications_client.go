package push_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new push notifications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for push notifications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRestapiV10SubscriptionSubscriptionID Cancel Subscription by ID
*/
func (a *Client) DeleteRestapiV10SubscriptionSubscriptionID(params *DeleteRestapiV10SubscriptionSubscriptionIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRestapiV10SubscriptionSubscriptionIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRestapiV10SubscriptionSubscriptionID",
		Method:             "DELETE",
		PathPattern:        "/restapi/v1.0/subscription/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRestapiV10SubscriptionSubscriptionIDReader{formats: a.formats},
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
GetRestapiV10Subscription Get Subscription List
*/
func (a *Client) GetRestapiV10Subscription(params *GetRestapiV10SubscriptionParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10SubscriptionParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10Subscription",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10SubscriptionReader{formats: a.formats},
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
GetRestapiV10SubscriptionSubscriptionID Get Subscription by ID
*/
func (a *Client) GetRestapiV10SubscriptionSubscriptionID(params *GetRestapiV10SubscriptionSubscriptionIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10SubscriptionSubscriptionIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10SubscriptionSubscriptionID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/subscription/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10SubscriptionSubscriptionIDReader{formats: a.formats},
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
PostRestapiV10Subscription Create New Subscription
*/
func (a *Client) PostRestapiV10Subscription(params *PostRestapiV10SubscriptionParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRestapiV10SubscriptionParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRestapiV10Subscription",
		Method:             "POST",
		PathPattern:        "/restapi/v1.0/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRestapiV10SubscriptionReader{formats: a.formats},
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
PutRestapiV10SubscriptionSubscriptionID Update/Renew Subscription by ID
*/
func (a *Client) PutRestapiV10SubscriptionSubscriptionID(params *PutRestapiV10SubscriptionSubscriptionIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutRestapiV10SubscriptionSubscriptionIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutRestapiV10SubscriptionSubscriptionID",
		Method:             "PUT",
		PathPattern:        "/restapi/v1.0/subscription/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutRestapiV10SubscriptionSubscriptionIDReader{formats: a.formats},
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
