// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Fulcio Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificateRequest certificate request
//
// swagger:model CertificateRequest
type CertificateRequest struct {

	// public key
	// Required: true
	PublicKey *CertificateRequestPublicKey `json:"publicKey"`

	// signed email address
	// Required: true
	// Format: byte
	SignedEmailAddress *strfmt.Base64 `json:"signedEmailAddress"`
}

// Validate validates this certificate request
func (m *CertificateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignedEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateRequest) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicKey")
			}
			return err
		}
	}

	return nil
}

func (m *CertificateRequest) validateSignedEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("signedEmailAddress", "body", m.SignedEmailAddress); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateRequest) UnmarshalBinary(b []byte) error {
	var res CertificateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateRequestPublicKey certificate request public key
//
// swagger:model CertificateRequestPublicKey
type CertificateRequestPublicKey struct {

	// algorithm
	// Required: true
	// Enum: [ecdsa]
	Algorithm *string `json:"algorithm"`

	// content
	// Required: true
	// Format: byte
	Content *strfmt.Base64 `json:"content"`
}

// Validate validates this certificate request public key
func (m *CertificateRequestPublicKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var certificateRequestPublicKeyTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ecdsa"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateRequestPublicKeyTypeAlgorithmPropEnum = append(certificateRequestPublicKeyTypeAlgorithmPropEnum, v)
	}
}

const (

	// CertificateRequestPublicKeyAlgorithmEcdsa captures enum value "ecdsa"
	CertificateRequestPublicKeyAlgorithmEcdsa string = "ecdsa"
)

// prop value enum
func (m *CertificateRequestPublicKey) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateRequestPublicKeyTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateRequestPublicKey) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("publicKey"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("publicKey"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *CertificateRequestPublicKey) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("publicKey"+"."+"content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateRequestPublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateRequestPublicKey) UnmarshalBinary(b []byte) error {
	var res CertificateRequestPublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
