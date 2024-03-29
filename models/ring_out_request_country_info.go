package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// RingOutRequestCountryInfo ring out request country info
// swagger:model RingOut.Request.CountryInfo
type RingOutRequestCountryInfo struct {

	// Dialing plan country identifier
	ID string `json:"id,omitempty"`
}

// Validate validates this ring out request country info
func (m *RingOutRequestCountryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
