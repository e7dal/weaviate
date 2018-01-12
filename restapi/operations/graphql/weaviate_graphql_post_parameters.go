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

package graphql

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateGraphqlPostParams creates a new WeaviateGraphqlPostParams object
// with the default values initialized.
func NewWeaviateGraphqlPostParams() WeaviateGraphqlPostParams {
	var ()
	return WeaviateGraphqlPostParams{}
}

// WeaviateGraphqlPostParams contains all the bound params for the weaviate graphql post operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.graphql.post
type WeaviateGraphqlPostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The GraphQL query request parameters.
	  Required: true
	  In: body
	*/
	Body *models.GraphQLQuery
	/*Takes a snapshot back in time, in case not set it will show the most recent results.
	  In: query
	*/
	Timesnap *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaviateGraphqlPostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.GraphQLQuery
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	qTimesnap, qhkTimesnap, _ := qs.GetOK("timesnap")
	if err := o.bindTimesnap(qTimesnap, qhkTimesnap, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateGraphqlPostParams) bindTimesnap(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("timesnap", "query", "int64", raw)
	}
	o.Timesnap = &value

	return nil
}
