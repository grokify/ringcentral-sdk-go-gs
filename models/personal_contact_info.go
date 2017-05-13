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

// PersonalContactInfo personal contact info
// swagger:model PersonalContactInfo
type PersonalContactInfo struct {

	// Assistant phone of a personal contact
	AssistantPhone string `json:"assistantPhone,omitempty"`

	// This property has a special meaning only on Address Book Sync (e.g. a contact can be 'Deleted'). For simple contact list reading it has always the default value - 'Alive'
	Availability string `json:"availability,omitempty"`

	// Date of birth of a personal contact in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	Birthday strfmt.DateTime `json:"birthday,omitempty"`

	// Business address of a personal contact
	BusinessAddress *ContactAddressInfo `json:"businessAddress,omitempty"`

	// Business fax of a personal contact
	BusinessFax string `json:"businessFax,omitempty"`

	// Business phone of a personal contact
	BusinessPhone string `json:"businessPhone,omitempty"`

	// The 2-d business phone of a personal contact
	BusinessPhone2 string `json:"businessPhone2,omitempty"`

	// Callback phone of a personal contact
	CallbackPhone string `json:"callbackPhone,omitempty"`

	// Car phone of a personal contact
	CarPhone string `json:"carPhone,omitempty"`

	// Company name of a personal contact
	Company string `json:"company,omitempty"`

	// Company phone of a personal contact
	CompanyPhone string `json:"companyPhone,omitempty"`

	// Email of a personal contact
	Email string `json:"email,omitempty"`

	// The 2-d email of a personal contact
	Email2 string `json:"email2,omitempty"`

	// The 3-d email of a personal contact
	Email3 string `json:"email3,omitempty"`

	// First name of a personal contact
	FirstName string `json:"firstName,omitempty"`

	// Home address of a personal contact
	HomeAddress *ContactAddressInfo `json:"homeAddress,omitempty"`

	// Home phone of a personal contact
	HomePhone string `json:"homePhone,omitempty"`

	// The 2-d home phone of a personal contact
	HomePhone2 string `json:"homePhone2,omitempty"`

	// Standard resource properties ID
	ID string `json:"id,omitempty"`

	// Job title of a personal contact
	JobTitle string `json:"jobTitle,omitempty"`

	// Last name of a personal contact
	LastName string `json:"lastName,omitempty"`

	// Middle name of a personal contact
	MiddleName string `json:"middleName,omitempty"`

	// Mobile phone of a personal contact
	MobilePhone string `json:"mobilePhone,omitempty"`

	// Nick name of a personal contact
	NickName string `json:"nickName,omitempty"`

	// Notes of a personal contact
	Notes string `json:"notes,omitempty"`

	// Other address of a personal contact
	OtherAddress *ContactAddressInfo `json:"otherAddress,omitempty"`

	// Other fax of a personal contact
	OtherFax string `json:"otherFax,omitempty"`

	// Other phone of a personal contact
	OtherPhone string `json:"otherPhone,omitempty"`

	// Canonical URI
	URL string `json:"url,omitempty"`

	// Web page of a personal contact
	WebPage string `json:"webPage,omitempty"`
}

// Validate validates this personal contact info
func (m *PersonalContactInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBusinessAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHomeAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOtherAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var personalContactInfoTypeAvailabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Alive","Deleted","Purged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		personalContactInfoTypeAvailabilityPropEnum = append(personalContactInfoTypeAvailabilityPropEnum, v)
	}
}

const (
	// PersonalContactInfoAvailabilityAlive captures enum value "Alive"
	PersonalContactInfoAvailabilityAlive string = "Alive"
	// PersonalContactInfoAvailabilityDeleted captures enum value "Deleted"
	PersonalContactInfoAvailabilityDeleted string = "Deleted"
	// PersonalContactInfoAvailabilityPurged captures enum value "Purged"
	PersonalContactInfoAvailabilityPurged string = "Purged"
)

// prop value enum
func (m *PersonalContactInfo) validateAvailabilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, personalContactInfoTypeAvailabilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PersonalContactInfo) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailabilityEnum("availability", "body", m.Availability); err != nil {
		return err
	}

	return nil
}

func (m *PersonalContactInfo) validateBusinessAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.BusinessAddress) { // not required
		return nil
	}

	if m.BusinessAddress != nil {

		if err := m.BusinessAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessAddress")
			}
			return err
		}
	}

	return nil
}

func (m *PersonalContactInfo) validateHomeAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.HomeAddress) { // not required
		return nil
	}

	if m.HomeAddress != nil {

		if err := m.HomeAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("homeAddress")
			}
			return err
		}
	}

	return nil
}

func (m *PersonalContactInfo) validateOtherAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.OtherAddress) { // not required
		return nil
	}

	if m.OtherAddress != nil {

		if err := m.OtherAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otherAddress")
			}
			return err
		}
	}

	return nil
}
