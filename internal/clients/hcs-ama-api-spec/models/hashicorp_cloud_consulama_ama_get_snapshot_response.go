// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaGetSnapshotResponse See ConsulAMAService.GetSnapshot
//
// swagger:model hashicorp.cloud.consulama.ama.GetSnapshotResponse
type HashicorpCloudConsulamaAmaGetSnapshotResponse struct {

	// snapshot contains the properties of the snapshot.
	Snapshot *HashicorpCloudConsulamaAmaSnapshotProperties `json:"snapshot,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama get snapshot response
func (m *HashicorpCloudConsulamaAmaGetSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaGetSnapshotResponse) validateSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaGetSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaGetSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaGetSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
