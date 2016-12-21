package subscriptions


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveSubscriptionsListHandlerFunc turns a function with the right signature into a weave subscriptions list handler
type WeaveSubscriptionsListHandlerFunc func(WeaveSubscriptionsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveSubscriptionsListHandlerFunc) Handle(params WeaveSubscriptionsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveSubscriptionsListHandler interface for that can handle valid weave subscriptions list params
type WeaveSubscriptionsListHandler interface {
	Handle(WeaveSubscriptionsListParams, interface{}) middleware.Responder
}

// NewWeaveSubscriptionsList creates a new http.Handler for the weave subscriptions list operation
func NewWeaveSubscriptionsList(ctx *middleware.Context, handler WeaveSubscriptionsListHandler) *WeaveSubscriptionsList {
	return &WeaveSubscriptionsList{Context: ctx, Handler: handler}
}

/*WeaveSubscriptionsList swagger:route GET /subscriptions subscriptions weaveSubscriptionsList

Lists subscriptions.

*/
type WeaveSubscriptionsList struct {
	Context *middleware.Context
	Handler WeaveSubscriptionsListHandler
}

func (o *WeaveSubscriptionsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveSubscriptionsListParams()

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