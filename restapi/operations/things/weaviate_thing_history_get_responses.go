/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @CreativeSofwFdn / yourfriends@weaviate.com
 */

package things

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingHistoryGetOKCode is the HTTP code returned for type WeaviateThingHistoryGetOK
const WeaviateThingHistoryGetOKCode int = 200

/*WeaviateThingHistoryGetOK Successful response.

swagger:response weaviateThingHistoryGetOK
*/
type WeaviateThingHistoryGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.ThingGetHistoryResponse `json:"body,omitempty"`
}

// NewWeaviateThingHistoryGetOK creates WeaviateThingHistoryGetOK with default headers values
func NewWeaviateThingHistoryGetOK() *WeaviateThingHistoryGetOK {
	return &WeaviateThingHistoryGetOK{}
}

// WithPayload adds the payload to the weaviate thing history get o k response
func (o *WeaviateThingHistoryGetOK) WithPayload(payload *models.ThingGetHistoryResponse) *WeaviateThingHistoryGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate thing history get o k response
func (o *WeaviateThingHistoryGetOK) SetPayload(payload *models.ThingGetHistoryResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingHistoryGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingHistoryGetUnauthorizedCode is the HTTP code returned for type WeaviateThingHistoryGetUnauthorized
const WeaviateThingHistoryGetUnauthorizedCode int = 401

/*WeaviateThingHistoryGetUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingHistoryGetUnauthorized
*/
type WeaviateThingHistoryGetUnauthorized struct {
}

// NewWeaviateThingHistoryGetUnauthorized creates WeaviateThingHistoryGetUnauthorized with default headers values
func NewWeaviateThingHistoryGetUnauthorized() *WeaviateThingHistoryGetUnauthorized {
	return &WeaviateThingHistoryGetUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingHistoryGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateThingHistoryGetForbiddenCode is the HTTP code returned for type WeaviateThingHistoryGetForbidden
const WeaviateThingHistoryGetForbiddenCode int = 403

/*WeaviateThingHistoryGetForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingHistoryGetForbidden
*/
type WeaviateThingHistoryGetForbidden struct {
}

// NewWeaviateThingHistoryGetForbidden creates WeaviateThingHistoryGetForbidden with default headers values
func NewWeaviateThingHistoryGetForbidden() *WeaviateThingHistoryGetForbidden {
	return &WeaviateThingHistoryGetForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingHistoryGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateThingHistoryGetNotFoundCode is the HTTP code returned for type WeaviateThingHistoryGetNotFound
const WeaviateThingHistoryGetNotFoundCode int = 404

/*WeaviateThingHistoryGetNotFound Successful query result but no resource was found.

swagger:response weaviateThingHistoryGetNotFound
*/
type WeaviateThingHistoryGetNotFound struct {
}

// NewWeaviateThingHistoryGetNotFound creates WeaviateThingHistoryGetNotFound with default headers values
func NewWeaviateThingHistoryGetNotFound() *WeaviateThingHistoryGetNotFound {
	return &WeaviateThingHistoryGetNotFound{}
}

// WriteResponse to the client
func (o *WeaviateThingHistoryGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateThingHistoryGetNotImplementedCode is the HTTP code returned for type WeaviateThingHistoryGetNotImplemented
const WeaviateThingHistoryGetNotImplementedCode int = 501

/*WeaviateThingHistoryGetNotImplemented Not (yet) implemented.

swagger:response weaviateThingHistoryGetNotImplemented
*/
type WeaviateThingHistoryGetNotImplemented struct {
}

// NewWeaviateThingHistoryGetNotImplemented creates WeaviateThingHistoryGetNotImplemented with default headers values
func NewWeaviateThingHistoryGetNotImplemented() *WeaviateThingHistoryGetNotImplemented {
	return &WeaviateThingHistoryGetNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateThingHistoryGetNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}