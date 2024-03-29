package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TokenInfo token info
// swagger:model TokenInfo
type TokenInfo struct {

	// Access token to pass to subsequent API requests
	AccessToken string `json:"access_token,omitempty"`

	// Unique identifier of a client application passed by the client, or auto-generated by server if not specified in request
	EndpointID string `json:"endpoint_id,omitempty"`

	// Issued access token TTL (time to live), in seconds
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// Extension identifier
	OwnerID string `json:"owner_id,omitempty"`

	// Refresh token to get a new access token, when the issued one expires
	RefreshToken string `json:"refresh_token,omitempty"`

	// Issued refresh token TTL (time to live), in seconds
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in,omitempty"`

	// List of permissions allowed with this access token, white-space separated
	Scope string `json:"scope,omitempty"`

	// Type of token. Use this parameter in Authorization header of requests
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this token info
func (m *TokenInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
