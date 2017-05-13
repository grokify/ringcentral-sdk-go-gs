package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// RuleInfoForwardingNumberInfo rule info forwarding number info
// swagger:model RuleInfo.ForwardingNumberInfo
type RuleInfoForwardingNumberInfo struct {

	// Internal identifier of a forwarding number
	ID string `json:"id,omitempty"`

	// Title of a forwarding number
	Label string `json:"label,omitempty"`

	// Phone number to which the call is forwarded
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Link to a forwarding number resource
	URI string `json:"uri,omitempty"`
}

// Validate validates this rule info forwarding number info
func (m *RuleInfoForwardingNumberInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}