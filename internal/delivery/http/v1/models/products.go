// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Products products
//
// swagger:model Products
type Products struct {

	// count
	Count int64 `json:"count,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// need recepi
	NeedRecepi bool `json:"need_recepi,omitempty"`

	// position
	Position string `json:"position,omitempty"`

	// price
	Price int64 `json:"price,omitempty"`
}

// Validate validates this products
func (m *Products) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this products based on context it is used
func (m *Products) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Products) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Products) UnmarshalBinary(b []byte) error {
	var res Products
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
