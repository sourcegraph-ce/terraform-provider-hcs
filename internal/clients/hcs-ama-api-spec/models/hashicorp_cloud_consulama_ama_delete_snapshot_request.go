// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaDeleteSnapshotRequest See ConsulAMAService.DeleteSnapshot
//
// swagger:model hashicorp.cloud.consulama.ama.DeleteSnapshotRequest
type HashicorpCloudConsulamaAmaDeleteSnapshotRequest struct {

	// resource_group is the resource group in which the Consul snapshot exists.
	// This is the AMA instance's managed resource group.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// snapshot_id is the ID of the Consul snapshot to delete:.
	SnapshotID string `json:"snapshotId,omitempty"`

	// subscription_id is the ID of the Azure subscription the Consul snapshot
	// exists in. This is the customer's subscription ID.
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama delete snapshot request
func (m *HashicorpCloudConsulamaAmaDeleteSnapshotRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaDeleteSnapshotRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaDeleteSnapshotRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaDeleteSnapshotRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
