package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InstantMessageNotification instant message notification
// swagger:model InstantMessageNotification
type InstantMessageNotification struct {

	// Notification payload body
	Body *InstantMessageEvent `json:"body,omitempty"`

	// Event filter URI
	Event string `json:"event,omitempty"`

	// Internal identifier of a subscription
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// Datetime of sending a notification in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// Universally unique identifier of a notification
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this instant message notification
func (m *InstantMessageNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstantMessageNotification) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {

		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			}
			return err
		}
	}

	return nil
}
