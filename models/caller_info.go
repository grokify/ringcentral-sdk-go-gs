package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// CallerInfo caller info
// swagger:model CallerInfo
type CallerInfo struct {

	// Extension short number (usually 3 or 4 digits). This property is filled when parties communicate by means of short internal numbers, for example when calling to other extension or sending/receiving Company Pager message
	ExtensionNumber string `json:"extensionNumber,omitempty"`

	// Contains party location (city, state) if one can be determined from phoneNumber. This property is filled only when phoneNumber is not empty and server can calculate location information from it (for example, this information is unavailable for US toll-free numbers)
	Location string `json:"location,omitempty"`

	// Symbolic name associated with a party. If the phone does not belong to the known extension, only the location is returned, the name is not determined then
	Name string `json:"name,omitempty"`

	// Phone number of a party. Usually it is a plain number including country and area code like 18661234567. But sometimes it could be returned from database with some formatting applied, for example (866)123-4567. This property is filled in all cases where parties communicate by means of global phone numbers, for example when calling to direct numbers or sending/receiving SMS
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// Validate validates this caller info
func (m *CallerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
