package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// RecipientInfo recipient info
// swagger:model RecipientInfo
type RecipientInfo struct {

	// Internal identifier of a recipient extension
	ID string `json:"id,omitempty"`

	// Link to a recipient extension resource
	URI string `json:"uri,omitempty"`
}

// Validate validates this recipient info
func (m *RecipientInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}