package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// AnsweringRuleInfoCallerInfo answering rule info caller info
// swagger:model AnsweringRuleInfo.CallerInfo
type AnsweringRuleInfoCallerInfo struct {

	// Phone number of a caller
	CallerID string `json:"callerId,omitempty"`

	// Contact name of a caller
	Name string `json:"name,omitempty"`
}

// Validate validates this answering rule info caller info
func (m *AnsweringRuleInfoCallerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}