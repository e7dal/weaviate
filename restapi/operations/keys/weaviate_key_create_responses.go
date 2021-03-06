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

package keys

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateKeyCreateOKCode is the HTTP code returned for type WeaviateKeyCreateOK
const WeaviateKeyCreateOKCode int = 200

/*WeaviateKeyCreateOK Successfully created.

swagger:response weaviateKeyCreateOK
*/
type WeaviateKeyCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.KeyTokenGetResponse `json:"body,omitempty"`
}

// NewWeaviateKeyCreateOK creates WeaviateKeyCreateOK with default headers values
func NewWeaviateKeyCreateOK() *WeaviateKeyCreateOK {
	return &WeaviateKeyCreateOK{}
}

// WithPayload adds the payload to the weaviate key create o k response
func (o *WeaviateKeyCreateOK) WithPayload(payload *models.KeyTokenGetResponse) *WeaviateKeyCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate key create o k response
func (o *WeaviateKeyCreateOK) SetPayload(payload *models.KeyTokenGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateKeyCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateKeyCreateUnauthorizedCode is the HTTP code returned for type WeaviateKeyCreateUnauthorized
const WeaviateKeyCreateUnauthorizedCode int = 401

/*WeaviateKeyCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateKeyCreateUnauthorized
*/
type WeaviateKeyCreateUnauthorized struct {
}

// NewWeaviateKeyCreateUnauthorized creates WeaviateKeyCreateUnauthorized with default headers values
func NewWeaviateKeyCreateUnauthorized() *WeaviateKeyCreateUnauthorized {
	return &WeaviateKeyCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateKeyCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateKeyCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateKeyCreateUnprocessableEntity
const WeaviateKeyCreateUnprocessableEntityCode int = 422

/*WeaviateKeyCreateUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateKeyCreateUnprocessableEntity
*/
type WeaviateKeyCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateKeyCreateUnprocessableEntity creates WeaviateKeyCreateUnprocessableEntity with default headers values
func NewWeaviateKeyCreateUnprocessableEntity() *WeaviateKeyCreateUnprocessableEntity {
	return &WeaviateKeyCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate key create unprocessable entity response
func (o *WeaviateKeyCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateKeyCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate key create unprocessable entity response
func (o *WeaviateKeyCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateKeyCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateKeyCreateNotImplementedCode is the HTTP code returned for type WeaviateKeyCreateNotImplemented
const WeaviateKeyCreateNotImplementedCode int = 501

/*WeaviateKeyCreateNotImplemented Not (yet) implemented.

swagger:response weaviateKeyCreateNotImplemented
*/
type WeaviateKeyCreateNotImplemented struct {
}

// NewWeaviateKeyCreateNotImplemented creates WeaviateKeyCreateNotImplemented with default headers values
func NewWeaviateKeyCreateNotImplemented() *WeaviateKeyCreateNotImplemented {
	return &WeaviateKeyCreateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateKeyCreateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
