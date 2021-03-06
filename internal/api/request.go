// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Request request
//
// swagger:model Request
type Request struct {

	// tree
	// Required: true
	Tree *Tree `json:"tree"`
}

// Validate validates this request
func (m *Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTree(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Request) validateTree(formats strfmt.Registry) error {

	if err := validate.Required("tree", "body", m.Tree); err != nil {
		return err
	}

	if m.Tree != nil {
		if err := m.Tree.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tree")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tree")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this request based on the context it is used
func (m *Request) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTree(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Request) contextValidateTree(ctx context.Context, formats strfmt.Registry) error {

	if m.Tree != nil {
		if err := m.Tree.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tree")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tree")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Request) UnmarshalBinary(b []byte) error {
	var res Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
