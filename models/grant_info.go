package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GrantInfo grant info
// swagger:model GrantInfo
type GrantInfo struct {

	// Specifies if monitoring of other extensions' calls is allowed for the extension. If 'CallMonitoring' feature is disabled for the given extension, the flag is not returned
	CallMonitoring bool `json:"callMonitoring,omitempty"`

	// Specifies if picking up of other extensions' calls is allowed for the extension. If 'Presence' feature is disabled for the given extension, the flag is not returned
	CallPickup bool `json:"callPickup,omitempty"`

	// Extension information
	Extension *GrantInfoExtensionInfo `json:"extension,omitempty"`

	// Canonical URI of a grant
	URI string `json:"uri,omitempty"`
}

// Validate validates this grant info
func (m *GrantInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrantInfo) validateExtension(formats strfmt.Registry) error {

	if swag.IsZero(m.Extension) { // not required
		return nil
	}

	if m.Extension != nil {

		if err := m.Extension.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension")
			}
			return err
		}
	}

	return nil
}
