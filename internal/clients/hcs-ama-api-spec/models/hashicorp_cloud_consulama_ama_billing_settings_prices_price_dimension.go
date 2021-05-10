// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension PriceDimension represents the prices of one billing dimension.
//
// swagger:model hashicorp.cloud.consulama.ama.BillingSettings.Prices.PriceDimension
type HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension struct {

	// minimum is the minimum number of billed units of this dimension.
	Minimum int32 `json:"minimum,omitempty"`

	// name is the dimension's name.
	Name string `json:"name,omitempty"`

	// tiers is the list of tiers to use for progressive discounting.
	//
	// Example:
	//     label: 1-50,     unit_price: 0.022
	//     label: 51-250,   unit_price: 0.018
	//     label: 250-∞,    unit_price: 0.014
	Tiers []*HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimensionTier `json:"tiers"`
}

// Validate validates this hashicorp cloud consulama ama billing settings prices price dimension
func (m *HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTiers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension) validateTiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Tiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Tiers); i++ {
		if swag.IsZero(m.Tiers[i]) { // not required
			continue
		}

		if m.Tiers[i] != nil {
			if err := m.Tiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
