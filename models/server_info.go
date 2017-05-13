package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerInfo server info
// swagger:model ServerInfo
type ServerInfo struct {

	// Full API version information: uri, number, release date
	APIVersions []*VersionInfo `json:"apiVersions"`

	// Server revision
	ServerRevision string `json:"serverRevision,omitempty"`

	// Server version
	ServerVersion string `json:"serverVersion,omitempty"`

	// Canonical URI of the API version
	URI string `json:"uri,omitempty"`
}

// Validate validates this server info
func (m *ServerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerInfo) validateAPIVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.APIVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.APIVersions); i++ {

		if swag.IsZero(m.APIVersions[i]) { // not required
			continue
		}

		if m.APIVersions[i] != nil {

			if err := m.APIVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
