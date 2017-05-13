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

// GrantInfoExtensionInfo grant info extension info
// swagger:model GrantInfo.ExtensionInfo
type GrantInfoExtensionInfo struct {

	// Extension short number (usually 3 or 4 digits)
	ExtensionNumber string `json:"extensionNumber,omitempty"`

	// Internal identifier of an extension
	ID string `json:"id,omitempty"`

	// Extension type
	Type string `json:"type,omitempty"`

	// Canonical URI of an extension
	URI string `json:"uri,omitempty"`
}

// Validate validates this grant info extension info
func (m *GrantInfoExtensionInfo) Validate(formats strfmt.Registry) error {
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

var grantInfoExtensionInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["User","Fax User","VirtualUser","DigitalUser","Department","Announcement","Voicemail","SharedLinesGroup","PagingOnly","IvrMenu","ApplicationExtension","Park Location"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		grantInfoExtensionInfoTypeTypePropEnum = append(grantInfoExtensionInfoTypeTypePropEnum, v)
	}
}

const (
	// GrantInfoExtensionInfoTypeUser captures enum value "User"
	GrantInfoExtensionInfoTypeUser string = "User"
	// GrantInfoExtensionInfoTypeFaxUser captures enum value "Fax User"
	GrantInfoExtensionInfoTypeFaxUser string = "Fax User"
	// GrantInfoExtensionInfoTypeVirtualUser captures enum value "VirtualUser"
	GrantInfoExtensionInfoTypeVirtualUser string = "VirtualUser"
	// GrantInfoExtensionInfoTypeDigitalUser captures enum value "DigitalUser"
	GrantInfoExtensionInfoTypeDigitalUser string = "DigitalUser"
	// GrantInfoExtensionInfoTypeDepartment captures enum value "Department"
	GrantInfoExtensionInfoTypeDepartment string = "Department"
	// GrantInfoExtensionInfoTypeAnnouncement captures enum value "Announcement"
	GrantInfoExtensionInfoTypeAnnouncement string = "Announcement"
	// GrantInfoExtensionInfoTypeVoicemail captures enum value "Voicemail"
	GrantInfoExtensionInfoTypeVoicemail string = "Voicemail"
	// GrantInfoExtensionInfoTypeSharedLinesGroup captures enum value "SharedLinesGroup"
	GrantInfoExtensionInfoTypeSharedLinesGroup string = "SharedLinesGroup"
	// GrantInfoExtensionInfoTypePagingOnly captures enum value "PagingOnly"
	GrantInfoExtensionInfoTypePagingOnly string = "PagingOnly"
	// GrantInfoExtensionInfoTypeIvrMenu captures enum value "IvrMenu"
	GrantInfoExtensionInfoTypeIvrMenu string = "IvrMenu"
	// GrantInfoExtensionInfoTypeApplicationExtension captures enum value "ApplicationExtension"
	GrantInfoExtensionInfoTypeApplicationExtension string = "ApplicationExtension"
	// GrantInfoExtensionInfoTypeParkLocation captures enum value "Park Location"
	GrantInfoExtensionInfoTypeParkLocation string = "Park Location"
)

// prop value enum
func (m *GrantInfoExtensionInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, grantInfoExtensionInfoTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GrantInfoExtensionInfo) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
