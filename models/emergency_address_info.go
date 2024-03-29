package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// EmergencyAddressInfo emergency address info
// swagger:model EmergencyAddressInfo
type EmergencyAddressInfo struct {

	// City name
	City string `json:"city,omitempty"`

	// Country name
	Country string `json:"country,omitempty"`

	// Name of a customer
	CustomerName string `json:"customerName,omitempty"`

	// State/province name
	State string `json:"state,omitempty"`

	// Street address, line 1 - street address, P.O. box, company name, c/o
	Street string `json:"street,omitempty"`

	// Street address, line 2 - apartment, suite, unit, building, floor, etc.
	Street2 string `json:"street2,omitempty"`

	// Zip code
	Zip string `json:"zip,omitempty"`
}

// Validate validates this emergency address info
func (m *EmergencyAddressInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
