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

// SubscriptionInfo subscription info
// swagger:model SubscriptionInfo
type SubscriptionInfo struct {

	// Subscription creation datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	CreationTime strfmt.DateTime `json:"creationTime,omitempty"`

	// Delivery mode data
	DeliveryMode *DeliveryMode `json:"deliveryMode,omitempty"`

	// Collection of URIs to API resources (message-store/presence/detailed presence)
	EventFilters []string `json:"eventFilters"`

	// Subscription expiration datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	ExpirationTime strfmt.DateTime `json:"expirationTime,omitempty"`

	// Subscription lifetime in seconds. The default value is 900
	ExpiresIn int64 `json:"expiresIn,omitempty"`

	// Internal identifier of a subscription
	ID string `json:"id,omitempty"`

	// Subscription status
	Status string `json:"status,omitempty"`

	// Canonical URI of a subscription
	URI string `json:"uri,omitempty"`
}

// Validate validates this subscription info
func (m *SubscriptionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeliveryMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventFilters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionInfo) validateDeliveryMode(formats strfmt.Registry) error {

	if swag.IsZero(m.DeliveryMode) { // not required
		return nil
	}

	if m.DeliveryMode != nil {

		if err := m.DeliveryMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliveryMode")
			}
			return err
		}
	}

	return nil
}

func (m *SubscriptionInfo) validateEventFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.EventFilters) { // not required
		return nil
	}

	return nil
}

var subscriptionInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Suspended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionInfoTypeStatusPropEnum = append(subscriptionInfoTypeStatusPropEnum, v)
	}
}

const (
	// SubscriptionInfoStatusActive captures enum value "Active"
	SubscriptionInfoStatusActive string = "Active"
	// SubscriptionInfoStatusSuspended captures enum value "Suspended"
	SubscriptionInfoStatusSuspended string = "Suspended"
)

// prop value enum
func (m *SubscriptionInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionInfoTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}
