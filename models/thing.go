/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @CreativeSofwFdn / yourfriends@weaviate.com
 */

package models

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Thing thing
// swagger:model Thing
type Thing struct {
	ThingCreate

	ThingAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Thing) UnmarshalJSON(raw []byte) error {

	var aO0 ThingCreate
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ThingCreate = aO0

	var aO1 ThingAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ThingAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Thing) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ThingCreate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ThingAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this thing
func (m *Thing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ThingCreate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.ThingAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Thing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thing) UnmarshalBinary(b []byte) error {
	var res Thing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
