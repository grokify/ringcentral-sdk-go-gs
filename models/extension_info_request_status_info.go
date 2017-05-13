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

// ExtensionInfoRequestStatusInfo extension info request status info
// swagger:model ExtensionInfo.Request.StatusInfo
type ExtensionInfoRequestStatusInfo struct {

	// Required extension status
	Status string `json:"status,omitempty"`

	// Extension status information, only for the 'Disabled' status
	StatusInfo *StatusInfo `json:"statusInfo,omitempty"`
}

// Validate validates this extension info request status info
func (m *ExtensionInfoRequestStatusInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatusInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var extensionInfoRequestStatusInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Disabled","Enabled","NotActivated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extensionInfoRequestStatusInfoTypeStatusPropEnum = append(extensionInfoRequestStatusInfoTypeStatusPropEnum, v)
	}
}

const (
	// ExtensionInfoRequestStatusInfoStatusDisabled captures enum value "Disabled"
	ExtensionInfoRequestStatusInfoStatusDisabled string = "Disabled"
	// ExtensionInfoRequestStatusInfoStatusEnabled captures enum value "Enabled"
	ExtensionInfoRequestStatusInfoStatusEnabled string = "Enabled"
	// ExtensionInfoRequestStatusInfoStatusNotActivated captures enum value "NotActivated"
	ExtensionInfoRequestStatusInfoStatusNotActivated string = "NotActivated"
)

// prop value enum
func (m *ExtensionInfoRequestStatusInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, extensionInfoRequestStatusInfoTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExtensionInfoRequestStatusInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ExtensionInfoRequestStatusInfo) validateStatusInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusInfo) { // not required
		return nil
	}

	if m.StatusInfo != nil {

		if err := m.StatusInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusInfo")
			}
			return err
		}
	}

	return nil
}