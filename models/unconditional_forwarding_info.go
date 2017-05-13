package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// UnconditionalForwardingInfo unconditional forwarding info
// swagger:model UnconditionalForwardingInfo
type UnconditionalForwardingInfo struct {

	// Phone number to which the call is forwarded
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// Validate validates this unconditional forwarding info
func (m *UnconditionalForwardingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
