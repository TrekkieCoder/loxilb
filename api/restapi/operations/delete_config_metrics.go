// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigMetricsHandlerFunc turns a function with the right signature into a delete config metrics handler
type DeleteConfigMetricsHandlerFunc func(DeleteConfigMetricsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigMetricsHandlerFunc) Handle(params DeleteConfigMetricsParams) middleware.Responder {
	return fn(params)
}

// DeleteConfigMetricsHandler interface for that can handle valid delete config metrics params
type DeleteConfigMetricsHandler interface {
	Handle(DeleteConfigMetricsParams) middleware.Responder
}

// NewDeleteConfigMetrics creates a new http.Handler for the delete config metrics operation
func NewDeleteConfigMetrics(ctx *middleware.Context, handler DeleteConfigMetricsHandler) *DeleteConfigMetrics {
	return &DeleteConfigMetrics{Context: ctx, Handler: handler}
}

/*
	DeleteConfigMetrics swagger:route DELETE /config/metrics deleteConfigMetrics

turn off prometheus option
*/
type DeleteConfigMetrics struct {
	Context *middleware.Context
	Handler DeleteConfigMetricsHandler
}

func (o *DeleteConfigMetrics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigMetricsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
