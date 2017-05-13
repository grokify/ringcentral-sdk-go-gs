package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PhoneLinesInfoPhoneNumberInfo phone lines info phone number info
// swagger:model PhoneLinesInfo.PhoneNumberInfo
type PhoneLinesInfoPhoneNumberInfo struct {

	// Brief information on a phone number country
	Country *CountryInfo `json:"country,omitempty"`

	// Internal identifier of a phone number
	ID string `json:"id,omitempty"`

	// Location (City, State). Filled for local US numbers
	Location string `json:"location,omitempty"`

	// Payment type. 'External' is returned for forwarded numbers which are not terminated in the RingCentral phone system
	PaymentType string `json:"paymentType,omitempty"`

	// Phone number
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Status of a phone number. If the value is 'Normal', the phone number is ready to be used. Otherwise it is an external number not yet ported to RingCentral
	Status string `json:"status,omitempty"`

	// Phone number type
	Type string `json:"type,omitempty"`

	// Usage type of the phone number
	UsageType string `json:"usageType,omitempty"`
}

// Validate validates this phone lines info phone number info
func (m *PhoneLinesInfoPhoneNumberInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsageType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhoneLinesInfoPhoneNumberInfo) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if m.Country != nil {

		if err := m.Country.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("country")
			}
			return err
		}
	}

	return nil
}

var phoneLinesInfoPhoneNumberInfoTypePaymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["External","TollFree","Local"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		phoneLinesInfoPhoneNumberInfoTypePaymentTypePropEnum = append(phoneLinesInfoPhoneNumberInfoTypePaymentTypePropEnum, v)
	}
}

const (
	// PhoneLinesInfoPhoneNumberInfoPaymentTypeExternal captures enum value "External"
	PhoneLinesInfoPhoneNumberInfoPaymentTypeExternal string = "External"
	// PhoneLinesInfoPhoneNumberInfoPaymentTypeTollFree captures enum value "TollFree"
	PhoneLinesInfoPhoneNumberInfoPaymentTypeTollFree string = "TollFree"
	// PhoneLinesInfoPhoneNumberInfoPaymentTypeLocal captures enum value "Local"
	PhoneLinesInfoPhoneNumberInfoPaymentTypeLocal string = "Local"
)

// prop value enum
func (m *PhoneLinesInfoPhoneNumberInfo) validatePaymentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, phoneLinesInfoPhoneNumberInfoTypePaymentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PhoneLinesInfoPhoneNumberInfo) validatePaymentType(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentTypeEnum("paymentType", "body", m.PaymentType); err != nil {
		return err
	}

	return nil
}

var phoneLinesInfoPhoneNumberInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VoiceFax","FaxOnly","VoiceOnly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		phoneLinesInfoPhoneNumberInfoTypeTypePropEnum = append(phoneLinesInfoPhoneNumberInfoTypeTypePropEnum, v)
	}
}

const (
	// PhoneLinesInfoPhoneNumberInfoTypeVoiceFax captures enum value "VoiceFax"
	PhoneLinesInfoPhoneNumberInfoTypeVoiceFax string = "VoiceFax"
	// PhoneLinesInfoPhoneNumberInfoTypeFaxOnly captures enum value "FaxOnly"
	PhoneLinesInfoPhoneNumberInfoTypeFaxOnly string = "FaxOnly"
	// PhoneLinesInfoPhoneNumberInfoTypeVoiceOnly captures enum value "VoiceOnly"
	PhoneLinesInfoPhoneNumberInfoTypeVoiceOnly string = "VoiceOnly"
)

// prop value enum
func (m *PhoneLinesInfoPhoneNumberInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, phoneLinesInfoPhoneNumberInfoTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PhoneLinesInfoPhoneNumberInfo) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var phoneLinesInfoPhoneNumberInfoTypeUsageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MainCompanyNumber","AdditionalCompanyNumber","CompanyNumber","DirectNumber","CompanyFaxNumber","ForwardedNumber"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		phoneLinesInfoPhoneNumberInfoTypeUsageTypePropEnum = append(phoneLinesInfoPhoneNumberInfoTypeUsageTypePropEnum, v)
	}
}

const (
	// PhoneLinesInfoPhoneNumberInfoUsageTypeMainCompanyNumber captures enum value "MainCompanyNumber"
	PhoneLinesInfoPhoneNumberInfoUsageTypeMainCompanyNumber string = "MainCompanyNumber"
	// PhoneLinesInfoPhoneNumberInfoUsageTypeAdditionalCompanyNumber captures enum value "AdditionalCompanyNumber"
	PhoneLinesInfoPhoneNumberInfoUsageTypeAdditionalCompanyNumber string = "AdditionalCompanyNumber"
	// PhoneLinesInfoPhoneNumberInfoUsageTypeCompanyNumber captures enum value "CompanyNumber"
	PhoneLinesInfoPhoneNumberInfoUsageTypeCompanyNumber string = "CompanyNumber"
	// PhoneLinesInfoPhoneNumberInfoUsageTypeDirectNumber captures enum value "DirectNumber"
	PhoneLinesInfoPhoneNumberInfoUsageTypeDirectNumber string = "DirectNumber"
	// PhoneLinesInfoPhoneNumberInfoUsageTypeCompanyFaxNumber captures enum value "CompanyFaxNumber"
	PhoneLinesInfoPhoneNumberInfoUsageTypeCompanyFaxNumber string = "CompanyFaxNumber"
	// PhoneLinesInfoPhoneNumberInfoUsageTypeForwardedNumber captures enum value "ForwardedNumber"
	PhoneLinesInfoPhoneNumberInfoUsageTypeForwardedNumber string = "ForwardedNumber"
)

// prop value enum
func (m *PhoneLinesInfoPhoneNumberInfo) validateUsageTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, phoneLinesInfoPhoneNumberInfoTypeUsageTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PhoneLinesInfoPhoneNumberInfo) validateUsageType(formats strfmt.Registry) error {

	if swag.IsZero(m.UsageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsageTypeEnum("usageType", "body", m.UsageType); err != nil {
		return err
	}

	return nil
}