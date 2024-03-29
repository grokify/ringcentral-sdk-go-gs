package address_book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new address book API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for address book API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID Delete Contact by ID
*/
func (a *Client) DeleteRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID(params *DeleteRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID",
		Method:             "DELETE",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContact Get Contact List
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContact(params *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContact",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID Get Contact by ID
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID(params *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookGroup Get Contact Group List
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookGroup(params *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookGroupParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookGroupParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookGroup",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookGroupReader{formats: a.formats},
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
GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSync Contacts Synchronization
*/
func (a *Client) GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSync(params *GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSync",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book-sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10AccountAccountIDExtensionExtensionIDAddressBookSyncReader{formats: a.formats},
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
PostRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContact Create New Contact
*/
func (a *Client) PostRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContact(params *PostRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContact",
		Method:             "POST",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactReader{formats: a.formats},
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
PutRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID Update Contact by ID
*/
func (a *Client) PutRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID(params *PutRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactID",
		Method:             "PUT",
		PathPattern:        "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutRestapiV10AccountAccountIDExtensionExtensionIDAddressBookContactContactIDReader{formats: a.formats},
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
