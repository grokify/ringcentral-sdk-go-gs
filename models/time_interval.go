package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TimeInterval time interval
// swagger:model TimeInterval
type TimeInterval struct {

	// Time in format hh:mm
	From strfmt.DateTime `json:"from,omitempty"`

	// Time in format hh:mm
	To strfmt.DateTime `json:"to,omitempty"`
}

// Validate validates this time interval
func (m *TimeInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}