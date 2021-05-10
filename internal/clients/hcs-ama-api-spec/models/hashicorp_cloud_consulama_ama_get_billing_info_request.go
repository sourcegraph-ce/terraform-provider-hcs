// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaGetBillingInfoRequest See ConsulAMAService.GetBillingInfo.
//
// swagger:model hashicorp.cloud.consulama.ama.GetBillingInfoRequest
type HashicorpCloudConsulamaAmaGetBillingInfoRequest struct {

	// resource_group is the managed resource group of the Managed App.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// subscription_id is the ID of the Azure subscription the Managed App is
	// deployed to. This is the customer's subscription ID.
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama get billing info request
func (m *HashicorpCloudConsulamaAmaGetBillingInfoRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaGetBillingInfoRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaGetBillingInfoRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaGetBillingInfoRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
