package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AnsweringRuleScheduleInfo answering rule schedule info
// swagger:model AnsweringRule.ScheduleInfo
type AnsweringRuleScheduleInfo struct {

	// Specific data ranges. If specified, weeklyRanges cannot be specified
	Ranges *RangesInfo `json:"ranges,omitempty"`

	// Weekly schedule. If specified, ranges cannot be specified
	WeeklyRanges *WeeklyScheduleInfo `json:"weeklyRanges,omitempty"`
}

// Validate validates this answering rule schedule info
func (m *AnsweringRuleScheduleInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRanges(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWeeklyRanges(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnsweringRuleScheduleInfo) validateRanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Ranges) { // not required
		return nil
	}

	if m.Ranges != nil {

		if err := m.Ranges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ranges")
			}
			return err
		}
	}

	return nil
}

func (m *AnsweringRuleScheduleInfo) validateWeeklyRanges(formats strfmt.Registry) error {

	if swag.IsZero(m.WeeklyRanges) { // not required
		return nil
	}

	if m.WeeklyRanges != nil {

		if err := m.WeeklyRanges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weeklyRanges")
			}
			return err
		}
	}

	return nil
}
