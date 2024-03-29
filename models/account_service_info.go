package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountServiceInfo account service info
// swagger:model Account.ServiceInfo
type AccountServiceInfo struct {

	// Information on account billing plan
	BillingPlan *BillingPlanInfo `json:"billingPlan,omitempty"`

	// Information on account brand
	Brand *BrandInfo `json:"brand,omitempty"`

	// Information on account service plan
	ServicePlan *ServicePlanInfo `json:"servicePlan,omitempty"`

	// Information on account target service plan
	TargetServicePlan *TargetServicePlanInfo `json:"targetServicePlan,omitempty"`

	// Canonical URI of a service info resource
	URI string `json:"uri,omitempty"`
}

// Validate validates this account service info
func (m *AccountServiceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBrand(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServicePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTargetServicePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountServiceInfo) validateBillingPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingPlan) { // not required
		return nil
	}

	if m.BillingPlan != nil {

		if err := m.BillingPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingPlan")
			}
			return err
		}
	}

	return nil
}

func (m *AccountServiceInfo) validateBrand(formats strfmt.Registry) error {

	if swag.IsZero(m.Brand) { // not required
		return nil
	}

	if m.Brand != nil {

		if err := m.Brand.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("brand")
			}
			return err
		}
	}

	return nil
}

func (m *AccountServiceInfo) validateServicePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicePlan) { // not required
		return nil
	}

	if m.ServicePlan != nil {

		if err := m.ServicePlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicePlan")
			}
			return err
		}
	}

	return nil
}

func (m *AccountServiceInfo) validateTargetServicePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetServicePlan) { // not required
		return nil
	}

	if m.TargetServicePlan != nil {

		if err := m.TargetServicePlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetServicePlan")
			}
			return err
		}
	}

	return nil
}
