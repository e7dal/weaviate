/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package places


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeavePlacesHandleInvitationHandlerFunc turns a function with the right signature into a weave places handle invitation handler
type WeavePlacesHandleInvitationHandlerFunc func(WeavePlacesHandleInvitationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeavePlacesHandleInvitationHandlerFunc) Handle(params WeavePlacesHandleInvitationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeavePlacesHandleInvitationHandler interface for that can handle valid weave places handle invitation params
type WeavePlacesHandleInvitationHandler interface {
	Handle(WeavePlacesHandleInvitationParams, interface{}) middleware.Responder
}

// NewWeavePlacesHandleInvitation creates a new http.Handler for the weave places handle invitation operation
func NewWeavePlacesHandleInvitation(ctx *middleware.Context, handler WeavePlacesHandleInvitationHandler) *WeavePlacesHandleInvitation {
	return &WeavePlacesHandleInvitation{Context: ctx, Handler: handler}
}

/*WeavePlacesHandleInvitation swagger:route POST /places/{placeId}/handleInvitation places weavePlacesHandleInvitation

Accepts or declines a pending place invitation.

*/
type WeavePlacesHandleInvitation struct {
	Context *middleware.Context
	Handler WeavePlacesHandleInvitationHandler
}

func (o *WeavePlacesHandleInvitation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeavePlacesHandleInvitationParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}