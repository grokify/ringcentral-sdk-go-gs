package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GroupInfo group info
// swagger:model GroupInfo
type GroupInfo struct {

	// Amount of contacts in a group
	ContactsCount int64 `json:"contactsCount,omitempty"`

	// Name of a group
	GroupName string `json:"groupName,omitempty"`

	// Internal identifier of a group
	ID string `json:"id,omitempty"`

	// Notes for a group
	Notes string `json:"notes,omitempty"`

	// Canonical URI of a group
	URI string `json:"uri,omitempty"`
}

// Validate validates this group info
func (m *GroupInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
