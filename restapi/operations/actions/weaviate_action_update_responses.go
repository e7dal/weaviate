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

package actions

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionUpdateAcceptedCode is the HTTP code returned for type WeaviateActionUpdateAccepted
const WeaviateActionUpdateAcceptedCode int = 202

/*WeaviateActionUpdateAccepted Successfully received.

swagger:response weaviateActionUpdateAccepted
*/
type WeaviateActionUpdateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ActionGetResponse `json:"body,omitempty"`
}

// NewWeaviateActionUpdateAccepted creates WeaviateActionUpdateAccepted with default headers values
func NewWeaviateActionUpdateAccepted() *WeaviateActionUpdateAccepted {
	return &WeaviateActionUpdateAccepted{}
}

// WithPayload adds the payload to the weaviate action update accepted response
func (o *WeaviateActionUpdateAccepted) WithPayload(payload *models.ActionGetResponse) *WeaviateActionUpdateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate action update accepted response
func (o *WeaviateActionUpdateAccepted) SetPayload(payload *models.ActionGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionUpdateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionUpdateUnauthorizedCode is the HTTP code returned for type WeaviateActionUpdateUnauthorized
const WeaviateActionUpdateUnauthorizedCode int = 401

/*WeaviateActionUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionUpdateUnauthorized
*/
type WeaviateActionUpdateUnauthorized struct {
}

// NewWeaviateActionUpdateUnauthorized creates WeaviateActionUpdateUnauthorized with default headers values
func NewWeaviateActionUpdateUnauthorized() *WeaviateActionUpdateUnauthorized {
	return &WeaviateActionUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateActionUpdateForbiddenCode is the HTTP code returned for type WeaviateActionUpdateForbidden
const WeaviateActionUpdateForbiddenCode int = 403

/*WeaviateActionUpdateForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionUpdateForbidden
*/
type WeaviateActionUpdateForbidden struct {
}

// NewWeaviateActionUpdateForbidden creates WeaviateActionUpdateForbidden with default headers values
func NewWeaviateActionUpdateForbidden() *WeaviateActionUpdateForbidden {
	return &WeaviateActionUpdateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateActionUpdateNotFoundCode is the HTTP code returned for type WeaviateActionUpdateNotFound
const WeaviateActionUpdateNotFoundCode int = 404

/*WeaviateActionUpdateNotFound Successful query result but no resource was found.

swagger:response weaviateActionUpdateNotFound
*/
type WeaviateActionUpdateNotFound struct {
}

// NewWeaviateActionUpdateNotFound creates WeaviateActionUpdateNotFound with default headers values
func NewWeaviateActionUpdateNotFound() *WeaviateActionUpdateNotFound {
	return &WeaviateActionUpdateNotFound{}
}

// WriteResponse to the client
func (o *WeaviateActionUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateActionUpdateUnprocessableEntityCode is the HTTP code returned for type WeaviateActionUpdateUnprocessableEntity
const WeaviateActionUpdateUnprocessableEntityCode int = 422

/*WeaviateActionUpdateUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateActionUpdateUnprocessableEntity
*/
type WeaviateActionUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionUpdateUnprocessableEntity creates WeaviateActionUpdateUnprocessableEntity with default headers values
func NewWeaviateActionUpdateUnprocessableEntity() *WeaviateActionUpdateUnprocessableEntity {
	return &WeaviateActionUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate action update unprocessable entity response
func (o *WeaviateActionUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateActionUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate action update unprocessable entity response
func (o *WeaviateActionUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionUpdateNotImplementedCode is the HTTP code returned for type WeaviateActionUpdateNotImplemented
const WeaviateActionUpdateNotImplementedCode int = 501

/*WeaviateActionUpdateNotImplemented Not (yet) implemented.

swagger:response weaviateActionUpdateNotImplemented
*/
type WeaviateActionUpdateNotImplemented struct {
}

// NewWeaviateActionUpdateNotImplemented creates WeaviateActionUpdateNotImplemented with default headers values
func NewWeaviateActionUpdateNotImplemented() *WeaviateActionUpdateNotImplemented {
	return &WeaviateActionUpdateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateActionUpdateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
