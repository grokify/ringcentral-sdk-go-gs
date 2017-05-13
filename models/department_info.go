package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// DepartmentInfo department info
// swagger:model DepartmentInfo
type DepartmentInfo struct {

	// Number of a department extension
	ExtensionNumber string `json:"extensionNumber,omitempty"`

	// Internal identifier of a department extension
	ID string `json:"id,omitempty"`

	// Canonical URI of a department extension
	URI string `json:"uri,omitempty"`
}

// Validate validates this department info
func (m *DepartmentInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}