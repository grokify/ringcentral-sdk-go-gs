package dictionary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dictionary API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dictionary API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRestapiV10DictionaryCountry Get Country List
*/
func (a *Client) GetRestapiV10DictionaryCountry(params *GetRestapiV10DictionaryCountryParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryCountryParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryCountry",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/country",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryCountryReader{formats: a.formats},
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
GetRestapiV10DictionaryCountryCountryID Get Country by ID
*/
func (a *Client) GetRestapiV10DictionaryCountryCountryID(params *GetRestapiV10DictionaryCountryCountryIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryCountryCountryIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryCountryCountryID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/country/{countryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryCountryCountryIDReader{formats: a.formats},
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
GetRestapiV10DictionaryLanguage Get Supported Language List
*/
func (a *Client) GetRestapiV10DictionaryLanguage(params *GetRestapiV10DictionaryLanguageParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryLanguageParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryLanguage",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/language",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryLanguageReader{formats: a.formats},
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
GetRestapiV10DictionaryLanguageLanguageID Get Language by ID
*/
func (a *Client) GetRestapiV10DictionaryLanguageLanguageID(params *GetRestapiV10DictionaryLanguageLanguageIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryLanguageLanguageIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryLanguageLanguageID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/language/{languageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryLanguageLanguageIDReader{formats: a.formats},
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
GetRestapiV10DictionaryLocation Get Location List
*/
func (a *Client) GetRestapiV10DictionaryLocation(params *GetRestapiV10DictionaryLocationParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryLocationParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryLocation",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/location",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryLocationReader{formats: a.formats},
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
GetRestapiV10DictionaryState Get State/Province List
*/
func (a *Client) GetRestapiV10DictionaryState(params *GetRestapiV10DictionaryStateParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryStateParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryState",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryStateReader{formats: a.formats},
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
GetRestapiV10DictionaryStateStateID Get State/Province by ID
*/
func (a *Client) GetRestapiV10DictionaryStateStateID(params *GetRestapiV10DictionaryStateStateIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryStateStateIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryStateStateID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/state/{stateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryStateStateIDReader{formats: a.formats},
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
GetRestapiV10DictionaryTimezone Get Time Zone List
*/
func (a *Client) GetRestapiV10DictionaryTimezone(params *GetRestapiV10DictionaryTimezoneParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryTimezoneParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryTimezone",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/timezone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryTimezoneReader{formats: a.formats},
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
GetRestapiV10DictionaryTimezoneTimezoneID Get Time Zone by ID
*/
func (a *Client) GetRestapiV10DictionaryTimezoneTimezoneID(params *GetRestapiV10DictionaryTimezoneTimezoneIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestapiV10DictionaryTimezoneTimezoneIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRestapiV10DictionaryTimezoneTimezoneID",
		Method:             "GET",
		PathPattern:        "/restapi/v1.0/dictionary/timezone/{timezoneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestapiV10DictionaryTimezoneTimezoneIDReader{formats: a.formats},
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