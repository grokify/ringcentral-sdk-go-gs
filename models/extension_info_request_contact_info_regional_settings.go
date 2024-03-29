package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtensionInfoRequestContactInfoRegionalSettings extension info request contact info regional settings
// swagger:model ExtensionInfo.Request.ContactInfo.RegionalSettings
type ExtensionInfoRequestContactInfoRegionalSettings struct {

	// Formatting language preferences for numbers, dates and currencies
	FormattingLocale *ExtensionInfoRequestContactInfoRegionalSettingsFormattingLocale `json:"formattingLocale,omitempty"`

	// Information on language used for telephony greetings
	GreetingLanguage *ExtensionInfoRequestContactInfoRegionalSettingsGreetingLanguage `json:"greetingLanguage,omitempty"`

	// User interface language data
	Language *ExtensionInfoRequestContactInfoRegionalSettingsLanguage `json:"language,omitempty"`

	// Timezone data
	Timezone *ExtensionInfoRequestContactInfoRegionalSettingsTimezone `json:"timezone,omitempty"`
}

// Validate validates this extension info request contact info regional settings
func (m *ExtensionInfoRequestContactInfoRegionalSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormattingLocale(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGreetingLanguage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionInfoRequestContactInfoRegionalSettings) validateFormattingLocale(formats strfmt.Registry) error {

	if swag.IsZero(m.FormattingLocale) { // not required
		return nil
	}

	if m.FormattingLocale != nil {

		if err := m.FormattingLocale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formattingLocale")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionInfoRequestContactInfoRegionalSettings) validateGreetingLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.GreetingLanguage) { // not required
		return nil
	}

	if m.GreetingLanguage != nil {

		if err := m.GreetingLanguage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("greetingLanguage")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionInfoRequestContactInfoRegionalSettings) validateLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.Language) { // not required
		return nil
	}

	if m.Language != nil {

		if err := m.Language.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionInfoRequestContactInfoRegionalSettings) validateTimezone(formats strfmt.Registry) error {

	if swag.IsZero(m.Timezone) { // not required
		return nil
	}

	if m.Timezone != nil {

		if err := m.Timezone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timezone")
			}
			return err
		}
	}

	return nil
}
