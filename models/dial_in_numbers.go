package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DialInNumbers dial in numbers
// swagger:model DialInNumbers
type DialInNumbers struct {

	// Country resource corresponding to the dial-in number
	Country *DialInNumbersCountryInfo `json:"country,omitempty"`

	// Phone number of the dial-in number formatted according to the extension home country
	FormattedNumber string `json:"formattedNumber,omitempty"`

	// Optional field in case the dial-in number is associated with a particular location
	Location string `json:"location,omitempty"`

	// Phone number of the dial-in number for the meeting in e.164 format
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// Validate validates this dial in numbers
func (m *DialInNumbers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DialInNumbers) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if m.Country != nil {

		if err := m.Country.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("country")
			}
			return err
		}
	}

	return nil
}
