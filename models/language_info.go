package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// LanguageInfo language info
// swagger:model LanguageInfo
type LanguageInfo struct {

	// Indicates whether a language is available as formatting locale
	FormattingLocale bool `json:"formattingLocale,omitempty"`

	// Indicates whether a language is available as greeting language
	Greeting bool `json:"greeting,omitempty"`

	// Internal identifier of a language
	ID string `json:"id,omitempty"`

	// Localization code of a language
	LocaleCode string `json:"localeCode,omitempty"`

	// Official name of a language
	Name string `json:"name,omitempty"`

	// Indicates whether a language is available as UI language
	UI bool `json:"ui,omitempty"`

	// Canonical URI of a language
	URI string `json:"uri,omitempty"`
}

// Validate validates this language info
func (m *LanguageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
