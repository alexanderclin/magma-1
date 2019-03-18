// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GenericCommandParams generic command params
// swagger:model generic_command_params
type GenericCommandParams struct {

	// command
	// Required: true
	// Min Length: 1
	Command *string `json:"command"`

	// params
	Params map[string]interface{} `json:"params,omitempty"`
}

// Validate validates this generic command params
func (m *GenericCommandParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericCommandParams) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	if err := validate.MinLength("command", "body", string(*m.Command), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericCommandParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericCommandParams) UnmarshalBinary(b []byte) error {
	var res GenericCommandParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
