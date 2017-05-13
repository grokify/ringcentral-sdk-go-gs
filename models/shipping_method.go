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

// ShippingMethod shipping method
// swagger:model ShippingMethod
type ShippingMethod struct {

	// Method identifier. The default value is "1" (Ground)
	ID string `json:"id,omitempty"`

	// Method name, corresponding to the identifier
	Name string `json:"name,omitempty"`
}

// Validate validates this shipping method
func (m *ShippingMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var shippingMethodTypeIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","2","3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shippingMethodTypeIDPropEnum = append(shippingMethodTypeIDPropEnum, v)
	}
}

const (
	// ShippingMethodIDNr1 captures enum value "1"
	ShippingMethodIDNr1 string = "1"
	// ShippingMethodIDNr2 captures enum value "2"
	ShippingMethodIDNr2 string = "2"
	// ShippingMethodIDNr3 captures enum value "3"
	ShippingMethodIDNr3 string = "3"
)

// prop value enum
func (m *ShippingMethod) validateIDEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, shippingMethodTypeIDPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ShippingMethod) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDEnum("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var shippingMethodTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ground","2 Day","Overnight"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shippingMethodTypeNamePropEnum = append(shippingMethodTypeNamePropEnum, v)
	}
}

const (
	// ShippingMethodNameGround captures enum value "Ground"
	ShippingMethodNameGround string = "Ground"
	// ShippingMethodNameNr2Day captures enum value "2 Day"
	ShippingMethodNameNr2Day string = "2 Day"
	// ShippingMethodNameOvernight captures enum value "Overnight"
	ShippingMethodNameOvernight string = "Overnight"
)

// prop value enum
func (m *ShippingMethod) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, shippingMethodTypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ShippingMethod) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
