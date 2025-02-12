// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeEncryptionSupport Indicates whether volume encryption is supported in the cluster.
//
//
// swagger:model volume_encryption_support
type VolumeEncryptionSupport struct {

	// Code corresponding to the status message. Returns a 0 if volume encryption is supported in all nodes of the cluster.
	// Example: 346758
	Code int64 `json:"code,omitempty"`

	// Reason for not supporting volume encryption.
	// Example: No platform support for volume encryption in following nodes - node1, node2.
	Message string `json:"message,omitempty"`

	// Set to true when volume encryption support is available on all nodes of the cluster.
	Supported bool `json:"supported,omitempty"`
}

// Validate validates this volume encryption support
func (m *VolumeEncryptionSupport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume encryption support based on context it is used
func (m *VolumeEncryptionSupport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeEncryptionSupport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEncryptionSupport) UnmarshalBinary(b []byte) error {
	var res VolumeEncryptionSupport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
