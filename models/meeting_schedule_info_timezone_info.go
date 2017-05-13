package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// MeetingScheduleInfoTimezoneInfo meeting schedule info timezone info
// swagger:model MeetingScheduleInfo.TimezoneInfo
type MeetingScheduleInfoTimezoneInfo struct {

	// Identifier of a timezone
	ID string `json:"id,omitempty"`
}

// Validate validates this meeting schedule info timezone info
func (m *MeetingScheduleInfoTimezoneInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}