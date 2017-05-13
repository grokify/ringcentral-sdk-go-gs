package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// SIPData s IP data
// swagger:model SIPData
type SIPData struct {

	// Sender data
	FromTag string `json:"fromTag,omitempty"`

	// Local address URL
	LocalURI string `json:"localUri,omitempty"`

	// Remote address URL
	RemoteURI string `json:"remoteUri,omitempty"`

	// Recipient data
	ToTag string `json:"toTag,omitempty"`
}

// Validate validates this s IP data
func (m *SIPData) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}