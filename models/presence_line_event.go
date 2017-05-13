package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PresenceLineEvent presence line event
// swagger:model PresenceLineEvent
type PresenceLineEvent struct {

	// Extension information
	Extension *PresenceLineEventExtensionInfo `json:"extension,omitempty"`

	// Order number of a notification to state the chronology
	Sequence int64 `json:"sequence,omitempty"`
}

// Validate validates this presence line event
func (m *PresenceLineEvent) Validate(formats strfmt.Registry) error {
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

func (m *PresenceLineEvent) validateExtension(formats strfmt.Registry) error {

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
