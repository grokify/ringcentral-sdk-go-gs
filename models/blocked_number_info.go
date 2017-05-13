package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// BlockedNumberInfo blocked number info
// swagger:model BlockedNumberInfo
type BlockedNumberInfo struct {

	// Standard resource properties ID and canonical URI, see the section called “Resource Identification Properties”
	ID string `json:"id,omitempty"`

	// Name assigned by a user to a blocked phone number
	Name string `json:"name,omitempty"`

	// Phone number to be blocked
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Canonical URI of a blocked number resource
	URI string `json:"uri,omitempty"`
}

// Validate validates this blocked number info
func (m *BlockedNumberInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}