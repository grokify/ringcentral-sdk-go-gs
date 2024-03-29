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

// LegInfo leg info
// swagger:model LegInfo
type LegInfo struct {

	// Action description of the call operation
	Action string `json:"action,omitempty"`

	// Call direction
	Direction string `json:"direction,omitempty"`

	// Call duration in seconds
	Duration int64 `json:"duration,omitempty"`

	// Information on extension
	Extension *LegInfoExtensionInfo `json:"extension,omitempty"`

	// Caller information
	From *CallerInfo `json:"from,omitempty"`

	// Leg type
	LegType string `json:"legType,omitempty"`

	// Call recording data. Returned if the call is recorded
	Recording *RecordingInfo `json:"recording,omitempty"`

	// Status description of the call operation
	Result string `json:"result,omitempty"`

	// The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// Callee information
	To *CallerInfo `json:"to,omitempty"`

	// Call transport
	Transport string `json:"transport,omitempty"`

	// Call type
	Type string `json:"type,omitempty"`
}

// Validate validates this leg info
func (m *LegInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExtension(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecording(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransport(formats); err != nil {
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

var legInfoTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Phone Call","Phone Login","Incoming Fax","Accept Call","FindMe","FollowMe","Outgoing Fax","Call Return","Calling Card","Ring Directly","RingOut Web","VoIP Call","RingOut PC","RingMe","Transfer","411 Info","Emergency","E911 Update","Support","RingOut Mobile"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		legInfoTypeActionPropEnum = append(legInfoTypeActionPropEnum, v)
	}
}

const (
	// LegInfoActionUnknown captures enum value "Unknown"
	LegInfoActionUnknown string = "Unknown"
	// LegInfoActionPhoneCall captures enum value "Phone Call"
	LegInfoActionPhoneCall string = "Phone Call"
	// LegInfoActionPhoneLogin captures enum value "Phone Login"
	LegInfoActionPhoneLogin string = "Phone Login"
	// LegInfoActionIncomingFax captures enum value "Incoming Fax"
	LegInfoActionIncomingFax string = "Incoming Fax"
	// LegInfoActionAcceptCall captures enum value "Accept Call"
	LegInfoActionAcceptCall string = "Accept Call"
	// LegInfoActionFindMe captures enum value "FindMe"
	LegInfoActionFindMe string = "FindMe"
	// LegInfoActionFollowMe captures enum value "FollowMe"
	LegInfoActionFollowMe string = "FollowMe"
	// LegInfoActionOutgoingFax captures enum value "Outgoing Fax"
	LegInfoActionOutgoingFax string = "Outgoing Fax"
	// LegInfoActionCallReturn captures enum value "Call Return"
	LegInfoActionCallReturn string = "Call Return"
	// LegInfoActionCallingCard captures enum value "Calling Card"
	LegInfoActionCallingCard string = "Calling Card"
	// LegInfoActionRingDirectly captures enum value "Ring Directly"
	LegInfoActionRingDirectly string = "Ring Directly"
	// LegInfoActionRingOutWeb captures enum value "RingOut Web"
	LegInfoActionRingOutWeb string = "RingOut Web"
	// LegInfoActionVoIPCall captures enum value "VoIP Call"
	LegInfoActionVoIPCall string = "VoIP Call"
	// LegInfoActionRingOutPC captures enum value "RingOut PC"
	LegInfoActionRingOutPC string = "RingOut PC"
	// LegInfoActionRingMe captures enum value "RingMe"
	LegInfoActionRingMe string = "RingMe"
	// LegInfoActionTransfer captures enum value "Transfer"
	LegInfoActionTransfer string = "Transfer"
	// LegInfoActionNr411Info captures enum value "411 Info"
	LegInfoActionNr411Info string = "411 Info"
	// LegInfoActionEmergency captures enum value "Emergency"
	LegInfoActionEmergency string = "Emergency"
	// LegInfoActionE911Update captures enum value "E911 Update"
	LegInfoActionE911Update string = "E911 Update"
	// LegInfoActionSupport captures enum value "Support"
	LegInfoActionSupport string = "Support"
	// LegInfoActionRingOutMobile captures enum value "RingOut Mobile"
	LegInfoActionRingOutMobile string = "RingOut Mobile"
)

// prop value enum
func (m *LegInfo) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, legInfoTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LegInfo) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var legInfoTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Inbound","Outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		legInfoTypeDirectionPropEnum = append(legInfoTypeDirectionPropEnum, v)
	}
}

const (
	// LegInfoDirectionInbound captures enum value "Inbound"
	LegInfoDirectionInbound string = "Inbound"
	// LegInfoDirectionOutbound captures enum value "Outbound"
	LegInfoDirectionOutbound string = "Outbound"
)

// prop value enum
func (m *LegInfo) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, legInfoTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LegInfo) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *LegInfo) validateExtension(formats strfmt.Registry) error {

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

func (m *LegInfo) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {

		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *LegInfo) validateRecording(formats strfmt.Registry) error {

	if swag.IsZero(m.Recording) { // not required
		return nil
	}

	if m.Recording != nil {

		if err := m.Recording.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recording")
			}
			return err
		}
	}

	return nil
}

var legInfoTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","ResultInProgress","Missed","Call accepted","Voicemail","Rejected","Reply","Received","Receive Error","Fax on Demand","Partial Receive","Blocked","Call connected","No Answer","International Disabled","Busy","Send Error","Sent","No fax machine","ResultEmpty","Account","Suspended","Call Failed","Call Failure","Internal Error","IP Phone offline","Restricted Number","Wrong Number","Stopped","Hang up","Poor Line Quality","Partially Sent","International Restriction","Abandoned","Declined","Fax Receipt Error","Fax Send Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		legInfoTypeResultPropEnum = append(legInfoTypeResultPropEnum, v)
	}
}

const (
	// LegInfoResultUnknown captures enum value "Unknown"
	LegInfoResultUnknown string = "Unknown"
	// LegInfoResultResultInProgress captures enum value "ResultInProgress"
	LegInfoResultResultInProgress string = "ResultInProgress"
	// LegInfoResultMissed captures enum value "Missed"
	LegInfoResultMissed string = "Missed"
	// LegInfoResultCallAccepted captures enum value "Call accepted"
	LegInfoResultCallAccepted string = "Call accepted"
	// LegInfoResultVoicemail captures enum value "Voicemail"
	LegInfoResultVoicemail string = "Voicemail"
	// LegInfoResultRejected captures enum value "Rejected"
	LegInfoResultRejected string = "Rejected"
	// LegInfoResultReply captures enum value "Reply"
	LegInfoResultReply string = "Reply"
	// LegInfoResultReceived captures enum value "Received"
	LegInfoResultReceived string = "Received"
	// LegInfoResultReceiveError captures enum value "Receive Error"
	LegInfoResultReceiveError string = "Receive Error"
	// LegInfoResultFaxOnDemand captures enum value "Fax on Demand"
	LegInfoResultFaxOnDemand string = "Fax on Demand"
	// LegInfoResultPartialReceive captures enum value "Partial Receive"
	LegInfoResultPartialReceive string = "Partial Receive"
	// LegInfoResultBlocked captures enum value "Blocked"
	LegInfoResultBlocked string = "Blocked"
	// LegInfoResultCallConnected captures enum value "Call connected"
	LegInfoResultCallConnected string = "Call connected"
	// LegInfoResultNoAnswer captures enum value "No Answer"
	LegInfoResultNoAnswer string = "No Answer"
	// LegInfoResultInternationalDisabled captures enum value "International Disabled"
	LegInfoResultInternationalDisabled string = "International Disabled"
	// LegInfoResultBusy captures enum value "Busy"
	LegInfoResultBusy string = "Busy"
	// LegInfoResultSendError captures enum value "Send Error"
	LegInfoResultSendError string = "Send Error"
	// LegInfoResultSent captures enum value "Sent"
	LegInfoResultSent string = "Sent"
	// LegInfoResultNoFaxMachine captures enum value "No fax machine"
	LegInfoResultNoFaxMachine string = "No fax machine"
	// LegInfoResultResultEmpty captures enum value "ResultEmpty"
	LegInfoResultResultEmpty string = "ResultEmpty"
	// LegInfoResultAccount captures enum value "Account"
	LegInfoResultAccount string = "Account"
	// LegInfoResultSuspended captures enum value "Suspended"
	LegInfoResultSuspended string = "Suspended"
	// LegInfoResultCallFailed captures enum value "Call Failed"
	LegInfoResultCallFailed string = "Call Failed"
	// LegInfoResultCallFailure captures enum value "Call Failure"
	LegInfoResultCallFailure string = "Call Failure"
	// LegInfoResultInternalError captures enum value "Internal Error"
	LegInfoResultInternalError string = "Internal Error"
	// LegInfoResultIPPhoneOffline captures enum value "IP Phone offline"
	LegInfoResultIPPhoneOffline string = "IP Phone offline"
	// LegInfoResultRestrictedNumber captures enum value "Restricted Number"
	LegInfoResultRestrictedNumber string = "Restricted Number"
	// LegInfoResultWrongNumber captures enum value "Wrong Number"
	LegInfoResultWrongNumber string = "Wrong Number"
	// LegInfoResultStopped captures enum value "Stopped"
	LegInfoResultStopped string = "Stopped"
	// LegInfoResultHangUp captures enum value "Hang up"
	LegInfoResultHangUp string = "Hang up"
	// LegInfoResultPoorLineQuality captures enum value "Poor Line Quality"
	LegInfoResultPoorLineQuality string = "Poor Line Quality"
	// LegInfoResultPartiallySent captures enum value "Partially Sent"
	LegInfoResultPartiallySent string = "Partially Sent"
	// LegInfoResultInternationalRestriction captures enum value "International Restriction"
	LegInfoResultInternationalRestriction string = "International Restriction"
	// LegInfoResultAbandoned captures enum value "Abandoned"
	LegInfoResultAbandoned string = "Abandoned"
	// LegInfoResultDeclined captures enum value "Declined"
	LegInfoResultDeclined string = "Declined"
	// LegInfoResultFaxReceiptError captures enum value "Fax Receipt Error"
	LegInfoResultFaxReceiptError string = "Fax Receipt Error"
	// LegInfoResultFaxSendError captures enum value "Fax Send Error"
	LegInfoResultFaxSendError string = "Fax Send Error"
)

// prop value enum
func (m *LegInfo) validateResultEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, legInfoTypeResultPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LegInfo) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

func (m *LegInfo) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if m.To != nil {

		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

var legInfoTypeTransportPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PSTN","VoIP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		legInfoTypeTransportPropEnum = append(legInfoTypeTransportPropEnum, v)
	}
}

const (
	// LegInfoTransportPSTN captures enum value "PSTN"
	LegInfoTransportPSTN string = "PSTN"
	// LegInfoTransportVoIP captures enum value "VoIP"
	LegInfoTransportVoIP string = "VoIP"
)

// prop value enum
func (m *LegInfo) validateTransportEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, legInfoTypeTransportPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LegInfo) validateTransport(formats strfmt.Registry) error {

	if swag.IsZero(m.Transport) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransportEnum("transport", "body", m.Transport); err != nil {
		return err
	}

	return nil
}

var legInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Voice","Fax"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		legInfoTypeTypePropEnum = append(legInfoTypeTypePropEnum, v)
	}
}

const (
	// LegInfoTypeVoice captures enum value "Voice"
	LegInfoTypeVoice string = "Voice"
	// LegInfoTypeFax captures enum value "Fax"
	LegInfoTypeFax string = "Fax"
)

// prop value enum
func (m *LegInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, legInfoTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LegInfo) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
