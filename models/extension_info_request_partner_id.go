package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ExtensionInfoRequestPartnerID extension info request partner Id
// swagger:model ExtensionInfo.Request.PartnerId
type ExtensionInfoRequestPartnerID struct {

	// Extension partner identifier
	PartnerID string `json:"partnerId,omitempty"`
}

// Validate validates this extension info request partner Id
func (m *ExtensionInfoRequestPartnerID) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}