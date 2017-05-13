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

// CallLogInfo call log info
// swagger:model CallLogInfo
type CallLogInfo struct {

	// Action description of the call operation
	Action string `json:"action,omitempty"`

	// Call direction
	Direction string `json:"direction,omitempty"`

	// Call duration in seconds
	Duration int64 `json:"duration,omitempty"`

	// Caller information
	From *CallerInfo `json:"from,omitempty"`

	// Internal identifier of a cal log record
	ID string `json:"id,omitempty"`

	// Call recording data. Returned if the call is recorded
	Recording *RecordingInfo `json:"recording,omitempty"`

	// Status description of the call operation
	Result string `json:"result,omitempty"`

	// Internal identifier of a call session
	SessionID string `json:"sessionId,omitempty"`

	// The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// Callee information
	To *CallerInfo `json:"to,omitempty"`

	// Call type
	Type string `json:"type,omitempty"`

	// Canonical URI of a call log record
	URI string `json:"uri,omitempty"`
}

// Validate validates this call log info
func (m *CallLogInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
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

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var callLogInfoTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Phone Call","Phone Login","Incoming Fax","Accept Call","FindMe","FollowMe","Outgoing Fax","Call Return","Calling Card","Ring Directly","RingOut Web","VoIP Call","RingOut PC","RingMe","Transfer","411 Info","Emergency","E911 Update","Support","RingOut Mobile"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callLogInfoTypeActionPropEnum = append(callLogInfoTypeActionPropEnum, v)
	}
}

const (
	// CallLogInfoActionUnknown captures enum value "Unknown"
	CallLogInfoActionUnknown string = "Unknown"
	// CallLogInfoActionPhoneCall captures enum value "Phone Call"
	CallLogInfoActionPhoneCall string = "Phone Call"
	// CallLogInfoActionPhoneLogin captures enum value "Phone Login"
	CallLogInfoActionPhoneLogin string = "Phone Login"
	// CallLogInfoActionIncomingFax captures enum value "Incoming Fax"
	CallLogInfoActionIncomingFax string = "Incoming Fax"
	// CallLogInfoActionAcceptCall captures enum value "Accept Call"
	CallLogInfoActionAcceptCall string = "Accept Call"
	// CallLogInfoActionFindMe captures enum value "FindMe"
	CallLogInfoActionFindMe string = "FindMe"
	// CallLogInfoActionFollowMe captures enum value "FollowMe"
	CallLogInfoActionFollowMe string = "FollowMe"
	// CallLogInfoActionOutgoingFax captures enum value "Outgoing Fax"
	CallLogInfoActionOutgoingFax string = "Outgoing Fax"
	// CallLogInfoActionCallReturn captures enum value "Call Return"
	CallLogInfoActionCallReturn string = "Call Return"
	// CallLogInfoActionCallingCard captures enum value "Calling Card"
	CallLogInfoActionCallingCard string = "Calling Card"
	// CallLogInfoActionRingDirectly captures enum value "Ring Directly"
	CallLogInfoActionRingDirectly string = "Ring Directly"
	// CallLogInfoActionRingOutWeb captures enum value "RingOut Web"
	CallLogInfoActionRingOutWeb string = "RingOut Web"
	// CallLogInfoActionVoIPCall captures enum value "VoIP Call"
	CallLogInfoActionVoIPCall string = "VoIP Call"
	// CallLogInfoActionRingOutPC captures enum value "RingOut PC"
	CallLogInfoActionRingOutPC string = "RingOut PC"
	// CallLogInfoActionRingMe captures enum value "RingMe"
	CallLogInfoActionRingMe string = "RingMe"
	// CallLogInfoActionTransfer captures enum value "Transfer"
	CallLogInfoActionTransfer string = "Transfer"
	// CallLogInfoActionNr411Info captures enum value "411 Info"
	CallLogInfoActionNr411Info string = "411 Info"
	// CallLogInfoActionEmergency captures enum value "Emergency"
	CallLogInfoActionEmergency string = "Emergency"
	// CallLogInfoActionE911Update captures enum value "E911 Update"
	CallLogInfoActionE911Update string = "E911 Update"
	// CallLogInfoActionSupport captures enum value "Support"
	CallLogInfoActionSupport string = "Support"
	// CallLogInfoActionRingOutMobile captures enum value "RingOut Mobile"
	CallLogInfoActionRingOutMobile string = "RingOut Mobile"
)

// prop value enum
func (m *CallLogInfo) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callLogInfoTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallLogInfo) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var callLogInfoTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Inbound","Outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callLogInfoTypeDirectionPropEnum = append(callLogInfoTypeDirectionPropEnum, v)
	}
}

const (
	// CallLogInfoDirectionInbound captures enum value "Inbound"
	CallLogInfoDirectionInbound string = "Inbound"
	// CallLogInfoDirectionOutbound captures enum value "Outbound"
	CallLogInfoDirectionOutbound string = "Outbound"
)

// prop value enum
func (m *CallLogInfo) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callLogInfoTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallLogInfo) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *CallLogInfo) validateFrom(formats strfmt.Registry) error {

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

func (m *CallLogInfo) validateRecording(formats strfmt.Registry) error {

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

var callLogInfoTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","ResultInProgress","Missed","Call accepted","Voicemail","Rejected","Reply","Received","Receive Error","Fax on Demand","Partial Receive","Blocked","Call connected","No Answer","International Disabled","Busy","Send Error","Sent","No fax machine","ResultEmpty","Account","Suspended","Call Failed","Call Failure","Internal Error","IP Phone offline","Restricted Number","Wrong Number","Stopped","Hang up","Poor Line Quality","Partially Sent","International Restriction","Abandoned","Declined","Fax Receipt Error","Fax Send Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callLogInfoTypeResultPropEnum = append(callLogInfoTypeResultPropEnum, v)
	}
}

const (
	// CallLogInfoResultUnknown captures enum value "Unknown"
	CallLogInfoResultUnknown string = "Unknown"
	// CallLogInfoResultResultInProgress captures enum value "ResultInProgress"
	CallLogInfoResultResultInProgress string = "ResultInProgress"
	// CallLogInfoResultMissed captures enum value "Missed"
	CallLogInfoResultMissed string = "Missed"
	// CallLogInfoResultCallAccepted captures enum value "Call accepted"
	CallLogInfoResultCallAccepted string = "Call accepted"
	// CallLogInfoResultVoicemail captures enum value "Voicemail"
	CallLogInfoResultVoicemail string = "Voicemail"
	// CallLogInfoResultRejected captures enum value "Rejected"
	CallLogInfoResultRejected string = "Rejected"
	// CallLogInfoResultReply captures enum value "Reply"
	CallLogInfoResultReply string = "Reply"
	// CallLogInfoResultReceived captures enum value "Received"
	CallLogInfoResultReceived string = "Received"
	// CallLogInfoResultReceiveError captures enum value "Receive Error"
	CallLogInfoResultReceiveError string = "Receive Error"
	// CallLogInfoResultFaxOnDemand captures enum value "Fax on Demand"
	CallLogInfoResultFaxOnDemand string = "Fax on Demand"
	// CallLogInfoResultPartialReceive captures enum value "Partial Receive"
	CallLogInfoResultPartialReceive string = "Partial Receive"
	// CallLogInfoResultBlocked captures enum value "Blocked"
	CallLogInfoResultBlocked string = "Blocked"
	// CallLogInfoResultCallConnected captures enum value "Call connected"
	CallLogInfoResultCallConnected string = "Call connected"
	// CallLogInfoResultNoAnswer captures enum value "No Answer"
	CallLogInfoResultNoAnswer string = "No Answer"
	// CallLogInfoResultInternationalDisabled captures enum value "International Disabled"
	CallLogInfoResultInternationalDisabled string = "International Disabled"
	// CallLogInfoResultBusy captures enum value "Busy"
	CallLogInfoResultBusy string = "Busy"
	// CallLogInfoResultSendError captures enum value "Send Error"
	CallLogInfoResultSendError string = "Send Error"
	// CallLogInfoResultSent captures enum value "Sent"
	CallLogInfoResultSent string = "Sent"
	// CallLogInfoResultNoFaxMachine captures enum value "No fax machine"
	CallLogInfoResultNoFaxMachine string = "No fax machine"
	// CallLogInfoResultResultEmpty captures enum value "ResultEmpty"
	CallLogInfoResultResultEmpty string = "ResultEmpty"
	// CallLogInfoResultAccount captures enum value "Account"
	CallLogInfoResultAccount string = "Account"
	// CallLogInfoResultSuspended captures enum value "Suspended"
	CallLogInfoResultSuspended string = "Suspended"
	// CallLogInfoResultCallFailed captures enum value "Call Failed"
	CallLogInfoResultCallFailed string = "Call Failed"
	// CallLogInfoResultCallFailure captures enum value "Call Failure"
	CallLogInfoResultCallFailure string = "Call Failure"
	// CallLogInfoResultInternalError captures enum value "Internal Error"
	CallLogInfoResultInternalError string = "Internal Error"
	// CallLogInfoResultIPPhoneOffline captures enum value "IP Phone offline"
	CallLogInfoResultIPPhoneOffline string = "IP Phone offline"
	// CallLogInfoResultRestrictedNumber captures enum value "Restricted Number"
	CallLogInfoResultRestrictedNumber string = "Restricted Number"
	// CallLogInfoResultWrongNumber captures enum value "Wrong Number"
	CallLogInfoResultWrongNumber string = "Wrong Number"
	// CallLogInfoResultStopped captures enum value "Stopped"
	CallLogInfoResultStopped string = "Stopped"
	// CallLogInfoResultHangUp captures enum value "Hang up"
	CallLogInfoResultHangUp string = "Hang up"
	// CallLogInfoResultPoorLineQuality captures enum value "Poor Line Quality"
	CallLogInfoResultPoorLineQuality string = "Poor Line Quality"
	// CallLogInfoResultPartiallySent captures enum value "Partially Sent"
	CallLogInfoResultPartiallySent string = "Partially Sent"
	// CallLogInfoResultInternationalRestriction captures enum value "International Restriction"
	CallLogInfoResultInternationalRestriction string = "International Restriction"
	// CallLogInfoResultAbandoned captures enum value "Abandoned"
	CallLogInfoResultAbandoned string = "Abandoned"
	// CallLogInfoResultDeclined captures enum value "Declined"
	CallLogInfoResultDeclined string = "Declined"
	// CallLogInfoResultFaxReceiptError captures enum value "Fax Receipt Error"
	CallLogInfoResultFaxReceiptError string = "Fax Receipt Error"
	// CallLogInfoResultFaxSendError captures enum value "Fax Send Error"
	CallLogInfoResultFaxSendError string = "Fax Send Error"
)

// prop value enum
func (m *CallLogInfo) validateResultEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callLogInfoTypeResultPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallLogInfo) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

func (m *CallLogInfo) validateTo(formats strfmt.Registry) error {

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

var callLogInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Voice","Fax"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callLogInfoTypeTypePropEnum = append(callLogInfoTypeTypePropEnum, v)
	}
}

const (
	// CallLogInfoTypeVoice captures enum value "Voice"
	CallLogInfoTypeVoice string = "Voice"
	// CallLogInfoTypeFax captures enum value "Fax"
	CallLogInfoTypeFax string = "Fax"
)

// prop value enum
func (m *CallLogInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, callLogInfoTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CallLogInfo) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
