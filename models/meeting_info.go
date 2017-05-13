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

// MeetingInfo meeting info
// swagger:model MeetingInfo
type MeetingInfo struct {

	// If 'True' then the meeting can be joined and started by any participant (not host only). Supported for the meetings of 'Scheduled' and 'Recurring' type.
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost,omitempty"`

	// Meeting audio options. Possible values are 'Phone', 'ComputerAudio'
	AudioOptions []string `json:"audioOptions"`

	// Internal identifier of a meeting as retrieved from Zoom
	ID string `json:"id,omitempty"`

	// Links to start/join the meeting
	Links *LinksInfo `json:"links,omitempty"`

	// Type of a meeting
	MeetingType string `json:"meetingType,omitempty"`

	// Password required to join a meeting
	Password string `json:"password,omitempty"`

	// Schedule of a meeting
	Schedule *MeetingScheduleInfo `json:"schedule,omitempty"`

	// Enables starting video when host joins the meeting
	StartHostVideo bool `json:"startHostVideo,omitempty"`

	// Enables starting video when participants join the meeting
	StartParticipantsVideo bool `json:"startParticipantsVideo,omitempty"`

	// Current status of a meeting
	Status string `json:"status,omitempty"`

	// Topic of a meeting
	Topic string `json:"topic,omitempty"`

	// Canonical URI of a meeting resource
	URI string `json:"uri,omitempty"`
}

// Validate validates this meeting info
func (m *MeetingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudioOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeetingType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
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

func (m *MeetingInfo) validateAudioOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.AudioOptions) { // not required
		return nil
	}

	return nil
}

func (m *MeetingInfo) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

var meetingInfoTypeMeetingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Scheduled","Instant","Recurring"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		meetingInfoTypeMeetingTypePropEnum = append(meetingInfoTypeMeetingTypePropEnum, v)
	}
}

const (
	// MeetingInfoMeetingTypeScheduled captures enum value "Scheduled"
	MeetingInfoMeetingTypeScheduled string = "Scheduled"
	// MeetingInfoMeetingTypeInstant captures enum value "Instant"
	MeetingInfoMeetingTypeInstant string = "Instant"
	// MeetingInfoMeetingTypeRecurring captures enum value "Recurring"
	MeetingInfoMeetingTypeRecurring string = "Recurring"
)

// prop value enum
func (m *MeetingInfo) validateMeetingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, meetingInfoTypeMeetingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MeetingInfo) validateMeetingType(formats strfmt.Registry) error {

	if swag.IsZero(m.MeetingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMeetingTypeEnum("meetingType", "body", m.MeetingType); err != nil {
		return err
	}

	return nil
}

func (m *MeetingInfo) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {

		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

var meetingInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotStarted","Started"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		meetingInfoTypeStatusPropEnum = append(meetingInfoTypeStatusPropEnum, v)
	}
}

const (
	// MeetingInfoStatusNotStarted captures enum value "NotStarted"
	MeetingInfoStatusNotStarted string = "NotStarted"
	// MeetingInfoStatusStarted captures enum value "Started"
	MeetingInfoStatusStarted string = "Started"
)

// prop value enum
func (m *MeetingInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, meetingInfoTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MeetingInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}