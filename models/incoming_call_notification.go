package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// IncomingCallNotification incoming call notification
// swagger:model IncomingCallNotification
type IncomingCallNotification struct {

	// Calling action, for example 'StartRing'
	Action string `json:"action,omitempty"`

	// Event filter URI
	Event string `json:"event,omitempty"`

	// Internal identifier of an extension
	ExtensionID *string `json:"extensionId,omitempty"`

	// Phone number of a caller
	From string `json:"from,omitempty"`

	// Caller name
	FromName string `json:"fromName,omitempty"`

	// File containing recorded caller name
	RecURL string `json:"recUrl,omitempty"`

	// Identifier of a server
	ServerID string `json:"serverId,omitempty"`

	// Identifier of a call session
	SessionID string `json:"sessionId,omitempty"`

	// Unique identifier of a session
	Sid string `json:"sid,omitempty"`

	// User data
	SrvLvl string `json:"srvLvl,omitempty"`

	// User data
	SrvLvlExt string `json:"srvLvlExt,omitempty"`

	// Internal identifier of a subscription
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// The datetime of a call action in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// Phone number of a callee
	To string `json:"to,omitempty"`

	// Callee name
	ToName string `json:"toName,omitempty"`

	// SIP proxy registration name
	ToURL string `json:"toUrl,omitempty"`

	// Universally unique identifier of a notification
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this incoming call notification
func (m *IncomingCallNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
