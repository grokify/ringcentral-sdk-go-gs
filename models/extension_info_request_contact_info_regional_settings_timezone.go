package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ExtensionInfoRequestContactInfoRegionalSettingsTimezone extension info request contact info regional settings timezone
// swagger:model ExtensionInfo.Request.ContactInfo.RegionalSettings.Timezone
type ExtensionInfoRequestContactInfoRegionalSettingsTimezone struct {

	// Timezone identifier. The default value is "58" (US&Canada)
	ID string `json:"id,omitempty"`
}

// Validate validates this extension info request contact info regional settings timezone
func (m *ExtensionInfoRequestContactInfoRegionalSettingsTimezone) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}