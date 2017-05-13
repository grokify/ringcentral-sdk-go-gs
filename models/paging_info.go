package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PagingInfo paging info
// swagger:model PagingInfo
type PagingInfo struct {

	// The current page number. 1-indexed, so the first page is 1 by default. May be omitted if result is empty (because non-existent page was specified or perPage=0 was requested)
	Page int64 `json:"page,omitempty"`

	// The zero-based index of the last element on the current page. Omitted if the page is omitted or result is empty
	PageEnd int64 `json:"pageEnd,omitempty"`

	// The zero-based number of the first element on the current page. Omitted if the page is omitted or result is empty
	PageStart int64 `json:"pageStart,omitempty"`

	// Current page size, describes how many items are in each page. Default value is 100. Maximum value is 1000. If perPage value in the request is greater than 1000, the maximum value (1000) is applied
	PerPage int64 `json:"perPage,omitempty"`

	// The total number of elements in a dataset. May be omitted for some resource due to performance reasons
	TotalElements int64 `json:"totalElements,omitempty"`

	// The total number of pages in a dataset. May be omitted for some resources due to performance reasons
	TotalPages int64 `json:"totalPages,omitempty"`
}

// Validate validates this paging info
func (m *PagingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}