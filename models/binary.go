package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Binary binary
// swagger:model Binary
type Binary struct {

	// Required. Binary data.
	Data io.ReadCloser `json:"data,omitempty"`
}

// Validate validates this binary
func (m *Binary) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
