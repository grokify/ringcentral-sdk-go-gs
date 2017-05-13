package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DetailedPresenceEvent detailed presence event
// swagger:model DetailedPresenceEvent
type DetailedPresenceEvent struct {

	// Collection of Active Call Info
	ActiveCalls []*DetailedPresenceEventActiveCallInfo `json:"activeCalls"`

	// If 'True' enables other extensions to see the extension presence status
	AllowSeeMyPresence bool `json:"allowSeeMyPresence,omitempty"`

	// Extended DnD (Do not Disturb) status
	DndStatus string `json:"dndStatus,omitempty"`

	// Internal identifier of an extension. Optional parameter
	ExtensionID *string `json:"extensionId,omitempty"`

	// If 'True' enables the extension user to pick up a monitored line on hold
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold,omitempty"`

	// Aggregated presence status, calculated from a number of sources
	PresenceStatus string `json:"presenceStatus,omitempty"`

	// If 'True' enables to ring extension phone, if any user monitored by this extension is ringing
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall,omitempty"`

	// Order number of a notification to state the chronology
	Sequence int64 `json:"sequence,omitempty"`

	// Telephony presence status. Returned if telephony status is changed. See Telephony Status Values
	TelephonyStatus string `json:"telephonyStatus,omitempty"`

	// Type of call termination. Supported for calls in 'NoCall' status. If the returned termination type is 'Intermediate' it means the call is not actually ended, the connection is established on one of the devices
	TerminationType string `json:"terminationType,omitempty"`

	// User-defined presence status (as previously published by the user)
	UserStatus string `json:"userStatus,omitempty"`
}

// Validate validates this detailed presence event
func (m *DetailedPresenceEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveCalls(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDndStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePresenceStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTelephonyStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTerminationType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedPresenceEvent) validateActiveCalls(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveCalls) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveCalls); i++ {

		if swag.IsZero(m.ActiveCalls[i]) { // not required
			continue
		}

		if m.ActiveCalls[i] != nil {

			if err := m.ActiveCalls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activeCalls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var detailedPresenceEventTypeDndStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TakeAllCalls","DoNotAcceptAnyCalls","DoNotAcceptDepartmentCalls","TakeDepartmentCallsOnly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedPresenceEventTypeDndStatusPropEnum = append(detailedPresenceEventTypeDndStatusPropEnum, v)
	}
}

const (
	// DetailedPresenceEventDndStatusTakeAllCalls captures enum value "TakeAllCalls"
	DetailedPresenceEventDndStatusTakeAllCalls string = "TakeAllCalls"
	// DetailedPresenceEventDndStatusDoNotAcceptAnyCalls captures enum value "DoNotAcceptAnyCalls"
	DetailedPresenceEventDndStatusDoNotAcceptAnyCalls string = "DoNotAcceptAnyCalls"
	// DetailedPresenceEventDndStatusDoNotAcceptDepartmentCalls captures enum value "DoNotAcceptDepartmentCalls"
	DetailedPresenceEventDndStatusDoNotAcceptDepartmentCalls string = "DoNotAcceptDepartmentCalls"
	// DetailedPresenceEventDndStatusTakeDepartmentCallsOnly captures enum value "TakeDepartmentCallsOnly"
	DetailedPresenceEventDndStatusTakeDepartmentCallsOnly string = "TakeDepartmentCallsOnly"
)

// prop value enum
func (m *DetailedPresenceEvent) validateDndStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedPresenceEventTypeDndStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedPresenceEvent) validateDndStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.DndStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateDndStatusEnum("dndStatus", "body", m.DndStatus); err != nil {
		return err
	}

	return nil
}

var detailedPresenceEventTypePresenceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Busy","Available"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedPresenceEventTypePresenceStatusPropEnum = append(detailedPresenceEventTypePresenceStatusPropEnum, v)
	}
}

const (
	// DetailedPresenceEventPresenceStatusOffline captures enum value "Offline"
	DetailedPresenceEventPresenceStatusOffline string = "Offline"
	// DetailedPresenceEventPresenceStatusBusy captures enum value "Busy"
	DetailedPresenceEventPresenceStatusBusy string = "Busy"
	// DetailedPresenceEventPresenceStatusAvailable captures enum value "Available"
	DetailedPresenceEventPresenceStatusAvailable string = "Available"
)

// prop value enum
func (m *DetailedPresenceEvent) validatePresenceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedPresenceEventTypePresenceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedPresenceEvent) validatePresenceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PresenceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePresenceStatusEnum("presenceStatus", "body", m.PresenceStatus); err != nil {
		return err
	}

	return nil
}

var detailedPresenceEventTypeTelephonyStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NoCall","CallConnected","Ringing","OnHold","ParkedCall"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedPresenceEventTypeTelephonyStatusPropEnum = append(detailedPresenceEventTypeTelephonyStatusPropEnum, v)
	}
}

const (
	// DetailedPresenceEventTelephonyStatusNoCall captures enum value "NoCall"
	DetailedPresenceEventTelephonyStatusNoCall string = "NoCall"
	// DetailedPresenceEventTelephonyStatusCallConnected captures enum value "CallConnected"
	DetailedPresenceEventTelephonyStatusCallConnected string = "CallConnected"
	// DetailedPresenceEventTelephonyStatusRinging captures enum value "Ringing"
	DetailedPresenceEventTelephonyStatusRinging string = "Ringing"
	// DetailedPresenceEventTelephonyStatusOnHold captures enum value "OnHold"
	DetailedPresenceEventTelephonyStatusOnHold string = "OnHold"
	// DetailedPresenceEventTelephonyStatusParkedCall captures enum value "ParkedCall"
	DetailedPresenceEventTelephonyStatusParkedCall string = "ParkedCall"
)

// prop value enum
func (m *DetailedPresenceEvent) validateTelephonyStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedPresenceEventTypeTelephonyStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedPresenceEvent) validateTelephonyStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TelephonyStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateTelephonyStatusEnum("telephonyStatus", "body", m.TelephonyStatus); err != nil {
		return err
	}

	return nil
}

var detailedPresenceEventTypeTerminationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["final","intermediate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedPresenceEventTypeTerminationTypePropEnum = append(detailedPresenceEventTypeTerminationTypePropEnum, v)
	}
}

const (
	// DetailedPresenceEventTerminationTypeFinal captures enum value "final"
	DetailedPresenceEventTerminationTypeFinal string = "final"
	// DetailedPresenceEventTerminationTypeIntermediate captures enum value "intermediate"
	DetailedPresenceEventTerminationTypeIntermediate string = "intermediate"
)

// prop value enum
func (m *DetailedPresenceEvent) validateTerminationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedPresenceEventTypeTerminationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedPresenceEvent) validateTerminationType(formats strfmt.Registry) error {

	if swag.IsZero(m.TerminationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTerminationTypeEnum("terminationType", "body", m.TerminationType); err != nil {
		return err
	}

	return nil
}

var detailedPresenceEventTypeUserStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Busy","Available"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedPresenceEventTypeUserStatusPropEnum = append(detailedPresenceEventTypeUserStatusPropEnum, v)
	}
}

const (
	// DetailedPresenceEventUserStatusOffline captures enum value "Offline"
	DetailedPresenceEventUserStatusOffline string = "Offline"
	// DetailedPresenceEventUserStatusBusy captures enum value "Busy"
	DetailedPresenceEventUserStatusBusy string = "Busy"
	// DetailedPresenceEventUserStatusAvailable captures enum value "Available"
	DetailedPresenceEventUserStatusAvailable string = "Available"
)

// prop value enum
func (m *DetailedPresenceEvent) validateUserStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedPresenceEventTypeUserStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedPresenceEvent) validateUserStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.UserStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserStatusEnum("userStatus", "body", m.UserStatus); err != nil {
		return err
	}

	return nil
}