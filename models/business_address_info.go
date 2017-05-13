package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// BusinessAddressInfo business address info
// swagger:model BusinessAddressInfo
type BusinessAddressInfo struct {

	// Name of a city
	City string `json:"city,omitempty"`

	// Name of a country
	Country string `json:"country,omitempty"`

	// Name of a state/province
	State string `json:"state,omitempty"`

	// Street address
	Street string `json:"street,omitempty"`

	// Zip code
	Zip string `json:"zip,omitempty"`
}

// Validate validates this business address info
func (m *BusinessAddressInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
