package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProfileImageInfo profile image info
// swagger:model ProfileImageInfo
type ProfileImageInfo struct {

	// The type of an image
	ContentType string `json:"contentType,omitempty"`

	// Identifier of an image
	Etag string `json:"etag,omitempty"`

	// The datetime when an image was last updated in ISO 8601 format, for example 2016-03-10T18:07:52.534Z
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// List of URIs to profile images in different dimensions
	Scales []*ImageURI `json:"scales"`

	// Link to a profile image. If an image is not uploaded for an extension, only uri is returned
	URI string `json:"uri,omitempty"`
}

// Validate validates this profile image info
func (m *ProfileImageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScales(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProfileImageInfo) validateScales(formats strfmt.Registry) error {

	if swag.IsZero(m.Scales) { // not required
		return nil
	}

	for i := 0; i < len(m.Scales); i++ {

		if swag.IsZero(m.Scales[i]) { // not required
			continue
		}

		if m.Scales[i] != nil {

			if err := m.Scales[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scales" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
