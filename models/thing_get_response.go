/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

package models

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingGetResponse thing get response
// swagger:model ThingGetResponse

type ThingGetResponse struct {
	Thing

	// thing Id
	ThingID strfmt.UUID `json:"thingId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ThingGetResponse) UnmarshalJSON(raw []byte) error {

	var aO0 Thing
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Thing = aO0

	var data struct {
		ThingID strfmt.UUID `json:"thingId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.ThingID = data.ThingID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ThingGetResponse) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Thing)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		ThingID strfmt.UUID `json:"thingId,omitempty"`
	}

	data.ThingID = m.ThingID

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this thing get response
func (m *ThingGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Thing.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ThingGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingGetResponse) UnmarshalBinary(b []byte) error {
	var res ThingGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
