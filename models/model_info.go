package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ModelInfo model info
// swagger:model ModelInfo
type ModelInfo struct {

	// Addons description
	Addons []*AddonInfo `json:"addons"`

	// Device model identifier. Mandatory when ordering a HardPhone if boxBillingId is not used for ordering
	ID string `json:"id,omitempty"`

	// Device name
	Name string `json:"name,omitempty"`
}

// Validate validates this model info
func (m *ModelInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddons(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelInfo) validateAddons(formats strfmt.Registry) error {

	if swag.IsZero(m.Addons) { // not required
		return nil
	}

	for i := 0; i < len(m.Addons); i++ {

		if swag.IsZero(m.Addons[i]) { // not required
			continue
		}

		if m.Addons[i] != nil {

			if err := m.Addons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
