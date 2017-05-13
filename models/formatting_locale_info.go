package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// FormattingLocaleInfo formatting locale info
// swagger:model FormattingLocaleInfo
type FormattingLocaleInfo struct {

	// Internal identifier of a formatting language
	ID string `json:"id,omitempty"`

	// Localization code of a formatting language
	LocaleCode string `json:"localeCode,omitempty"`

	// Official name of a formatting language
	Name string `json:"name,omitempty"`
}

// Validate validates this formatting locale info
func (m *FormattingLocaleInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}