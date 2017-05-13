package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtensionPermissions extension permissions
// swagger:model ExtensionPermissions
type ExtensionPermissions struct {

	// Admin permission
	Admin *PermissionInfo `json:"admin,omitempty"`

	// International Calling permission
	InternationalCalling *PermissionInfo `json:"internationalCalling,omitempty"`
}

// Validate validates this extension permissions
func (m *ExtensionPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmin(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInternationalCalling(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionPermissions) validateAdmin(formats strfmt.Registry) error {

	if swag.IsZero(m.Admin) { // not required
		return nil
	}

	if m.Admin != nil {

		if err := m.Admin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionPermissions) validateInternationalCalling(formats strfmt.Registry) error {

	if swag.IsZero(m.InternationalCalling) { // not required
		return nil
	}

	if m.InternationalCalling != nil {

		if err := m.InternationalCalling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internationalCalling")
			}
			return err
		}
	}

	return nil
}