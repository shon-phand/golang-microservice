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

// User user
// swagger:model User
type User struct {

	// date created
	// Read Only: true
	// Min Length: 1
	DateCreated string `json:"date_created,omitempty"`

	// email
	// Required: true
	// Min Length: 1
	Email *string `json:"email"`

	// first name
	// Required: true
	// Min Length: 1
	FirstName *string `json:"first_name"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// last name
	// Required: true
	// Min Length: 1
	LastName *string `json:"last_name"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.MinLength("date_created", "body", string(m.DateCreated), 1); err != nil {
		return err
	}

	return nil
}

func (m *User) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.MinLength("email", "body", string(*m.Email), 1); err != nil {
		return err
	}

	return nil
}

func (m *User) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MinLength("first_name", "body", string(*m.FirstName), 1); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	if err := validate.MinLength("last_name", "body", string(*m.LastName), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}