// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VersionGetEntry version get entry
//
// swagger:model VersionGetEntry
type VersionGetEntry struct {

	// build info
	BuildInfo string `json:"buildInfo,omitempty"`

	// Instance name
	Version string `json:"version,omitempty"`
}

// Validate validates this version get entry
func (m *VersionGetEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this version get entry based on context it is used
func (m *VersionGetEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VersionGetEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionGetEntry) UnmarshalBinary(b []byte) error {
	var res VersionGetEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
