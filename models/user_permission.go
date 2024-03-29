package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserPermission user permission
// swagger:model UserPermission
type UserPermission struct {

	// Information on a permission granted
	Permission *UserPermissionInfo `json:"permission,omitempty"`

	// List of active scopes for permission
	Scopes []string `json:"scopes"`
}

// Validate validates this user permission
func (m *UserPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermission(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScopes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPermission) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if m.Permission != nil {

		if err := m.Permission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permission")
			}
			return err
		}
	}

	return nil
}

func (m *UserPermission) validateScopes(formats strfmt.Registry) error {

	if swag.IsZero(m.Scopes) { // not required
		return nil
	}

	return nil
}
