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

// MessageChange message change
// swagger:model MessageChange
type MessageChange struct {

	// The number of new messages. Can be omitted if the value is zero
	NewCount int64 `json:"newCount,omitempty"`

	// Message type
	Type string `json:"type,omitempty"`

	// The number of updated messages. Can be omitted if the value is zero
	UpdatedCount int64 `json:"updatedCount,omitempty"`
}

// Validate validates this message change
func (m *MessageChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var messageChangeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Voicemail","SMS","Fax","Pager"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageChangeTypeTypePropEnum = append(messageChangeTypeTypePropEnum, v)
	}
}

const (
	// MessageChangeTypeVoicemail captures enum value "Voicemail"
	MessageChangeTypeVoicemail string = "Voicemail"
	// MessageChangeTypeSMS captures enum value "SMS"
	MessageChangeTypeSMS string = "SMS"
	// MessageChangeTypeFax captures enum value "Fax"
	MessageChangeTypeFax string = "Fax"
	// MessageChangeTypePager captures enum value "Pager"
	MessageChangeTypePager string = "Pager"
)

// prop value enum
func (m *MessageChange) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, messageChangeTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MessageChange) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}