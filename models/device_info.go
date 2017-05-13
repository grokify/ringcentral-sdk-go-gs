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

// DeviceInfo device info
// swagger:model DeviceInfo
type DeviceInfo struct {

	// Box billing identifier of a device. Applicable only for HardPhones. It is an alternative way to identify the device to be ordered. Either "model" structure, or "boxBillingId" must be specified for HardPhone
	BoxBillingID int64 `json:"boxBillingId,omitempty"`

	// PC name for softphone
	ComputerName string `json:"computerName,omitempty"`

	// Address for emergency cases. The same emergency address is assigned to all the numbers of one device
	EmergencyServiceAddress *EmergencyAddressInfo `json:"emergencyServiceAddress,omitempty"`

	// This attribute can be omitted for unassigned devices
	Extension *DeviceInfoExtensionInfo `json:"extension,omitempty"`

	// Internal identifier of a device
	ID string `json:"id,omitempty"`

	// HardPhone model information
	Model *ModelInfo `json:"model,omitempty"`

	// Device name. Mandatory if ordering "SoftPhone" or "OtherPhone". Optional for "HardPhone". If not specified for HardPhone, then device "model" name is used as device "name"
	Name string `json:"name,omitempty"`

	// Phone lines information
	PhoneLines *PhoneLinesInfo `json:"phoneLines,omitempty"`

	// Serial number for HardPhone (is returned only when the phone is shipped and provisioned); endpoint_id for softphone and mobile applications
	Serial string `json:"serial,omitempty"`

	// Shipping information, according to which devices (in case of "HardPhone") or e911 stickers (in case of "SoftPhone" and "OtherPhone") will be delivered to the customer
	Shipping *ShippingInfo `json:"shipping,omitempty"`

	// Device identification number (stock keeping unit) in the format TP-ID [-AT-AC], where TP is device type (HP for RC HardPhone, DV for all other devices including softphone); ID - device model ID; AT -addon type ID; AC - addon count (if any). For example 'HP-56-2-2'
	Sku string `json:"sku,omitempty"`

	// Device type. The default value is 'HardPhone'
	Type string `json:"type,omitempty"`

	// Canonical URI of a device
	URI string `json:"uri,omitempty"`
}

// Validate validates this device info
func (m *DeviceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmergencyServiceAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExtension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhoneLines(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShipping(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceInfo) validateEmergencyServiceAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.EmergencyServiceAddress) { // not required
		return nil
	}

	if m.EmergencyServiceAddress != nil {

		if err := m.EmergencyServiceAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emergencyServiceAddress")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfo) validateExtension(formats strfmt.Registry) error {

	if swag.IsZero(m.Extension) { // not required
		return nil
	}

	if m.Extension != nil {

		if err := m.Extension.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfo) validateModel(formats strfmt.Registry) error {

	if swag.IsZero(m.Model) { // not required
		return nil
	}

	if m.Model != nil {

		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfo) validatePhoneLines(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneLines) { // not required
		return nil
	}

	if m.PhoneLines != nil {

		if err := m.PhoneLines.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneLines")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInfo) validateShipping(formats strfmt.Registry) error {

	if swag.IsZero(m.Shipping) { // not required
		return nil
	}

	if m.Shipping != nil {

		if err := m.Shipping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipping")
			}
			return err
		}
	}

	return nil
}

var deviceInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SoftPhone","OtherPhone","HardPhone"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceInfoTypeTypePropEnum = append(deviceInfoTypeTypePropEnum, v)
	}
}

const (
	// DeviceInfoTypeSoftPhone captures enum value "SoftPhone"
	DeviceInfoTypeSoftPhone string = "SoftPhone"
	// DeviceInfoTypeOtherPhone captures enum value "OtherPhone"
	DeviceInfoTypeOtherPhone string = "OtherPhone"
	// DeviceInfoTypeHardPhone captures enum value "HardPhone"
	DeviceInfoTypeHardPhone string = "HardPhone"
)

// prop value enum
func (m *DeviceInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceInfoTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceInfo) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}