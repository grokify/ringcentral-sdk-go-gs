package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RegionalSettings regional settings
// swagger:model RegionalSettings
type RegionalSettings struct {

	// Formatting language preferences for numbers, dates and currencies
	FormattingLocale *FormattingLocaleInfo `json:"formattingLocale,omitempty"`

	// Information on language used for telephony greetings
	GreetingLanguage *GreetingLanguageInfo `json:"greetingLanguage,omitempty"`

	// Extension country information
	HomeCountry *CountryInfo `json:"homeCountry,omitempty"`

	// User interface language data
	Language *LanguageInfo `json:"language,omitempty"`

	// Extension timezone information
	Timezone *TimezoneInfo `json:"timezone,omitempty"`
}

// Validate validates this regional settings
func (m *RegionalSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormattingLocale(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGreetingLanguage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHomeCountry(formats); err != nil {
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

func (m *RegionalSettings) validateFormattingLocale(formats strfmt.Registry) error {

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

func (m *RegionalSettings) validateGreetingLanguage(formats strfmt.Registry) error {

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

func (m *RegionalSettings) validateHomeCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.HomeCountry) { // not required
		return nil
	}

	if m.HomeCountry != nil {

		if err := m.HomeCountry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("homeCountry")
			}
			return err
		}
	}

	return nil
}

func (m *RegionalSettings) validateLanguage(formats strfmt.Registry) error {

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

func (m *RegionalSettings) validateTimezone(formats strfmt.Registry) error {

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
