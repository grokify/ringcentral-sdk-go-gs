package meetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new meetings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for meetings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID Delete Meeting
*/
func (a *Client) DeleteRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID(params *DeleteRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID",
		Method:             "DELETE",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDMeeting Get Meetings List
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDMeeting(params *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDMeeting",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID Get Meeting
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID(params *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfo Get Meeting Service Info
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfo(params *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfo",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/service-info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoReader{formats: a.formats},
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
PostRestapiV10AccountAccountIDExtensionExtensionIDMeeting Create Meeting
*/
func (a *Client) PostRestapiV10AccountAccountIDExtensionExtensionIDMeeting(params *PostRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRestapiV10AccountAccountIDExtensionExtensionIDMeetingParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRestapiV10AccountAccountIDExtensionExtensionIDMeeting",
		Method:             "POST",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRestapiV10AccountAccountIDExtensionExtensionIDMeetingReader{formats: a.formats},
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
PostRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDEnd End Current Meeting
*/
func (a *Client) PostRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDEnd(params *PostRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDEndParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDEndParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDEnd",
		Method:             "POST",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId}/end",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDEndReader{formats: a.formats},
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
PutRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID Update Meeting
*/
func (a *Client) PutRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID(params *PutRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingID",
		Method:             "PUT",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutRestapiV10AccountAccountIDExtensionExtensionIDMeetingMeetingIDReader{formats: a.formats},
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