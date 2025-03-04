// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigPortAllOKCode is the HTTP code returned for type GetConfigPortAllOK
const GetConfigPortAllOKCode int = 200

/*
GetConfigPortAllOK OK

swagger:response getConfigPortAllOK
*/
type GetConfigPortAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigPortAllOKBody `json:"body,omitempty"`
}

// NewGetConfigPortAllOK creates GetConfigPortAllOK with default headers values
func NewGetConfigPortAllOK() *GetConfigPortAllOK {

	return &GetConfigPortAllOK{}
}

// WithPayload adds the payload to the get config port all o k response
func (o *GetConfigPortAllOK) WithPayload(payload *GetConfigPortAllOKBody) *GetConfigPortAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config port all o k response
func (o *GetConfigPortAllOK) SetPayload(payload *GetConfigPortAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigPortAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigPortAllUnauthorizedCode is the HTTP code returned for type GetConfigPortAllUnauthorized
const GetConfigPortAllUnauthorizedCode int = 401

/*
GetConfigPortAllUnauthorized Invalid authentication credentials

swagger:response getConfigPortAllUnauthorized
*/
type GetConfigPortAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigPortAllUnauthorized creates GetConfigPortAllUnauthorized with default headers values
func NewGetConfigPortAllUnauthorized() *GetConfigPortAllUnauthorized {

	return &GetConfigPortAllUnauthorized{}
}

// WithPayload adds the payload to the get config port all unauthorized response
func (o *GetConfigPortAllUnauthorized) WithPayload(payload *models.Error) *GetConfigPortAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config port all unauthorized response
func (o *GetConfigPortAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigPortAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigPortAllInternalServerErrorCode is the HTTP code returned for type GetConfigPortAllInternalServerError
const GetConfigPortAllInternalServerErrorCode int = 500

/*
GetConfigPortAllInternalServerError Internal service error

swagger:response getConfigPortAllInternalServerError
*/
type GetConfigPortAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigPortAllInternalServerError creates GetConfigPortAllInternalServerError with default headers values
func NewGetConfigPortAllInternalServerError() *GetConfigPortAllInternalServerError {

	return &GetConfigPortAllInternalServerError{}
}

// WithPayload adds the payload to the get config port all internal server error response
func (o *GetConfigPortAllInternalServerError) WithPayload(payload *models.Error) *GetConfigPortAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config port all internal server error response
func (o *GetConfigPortAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigPortAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigPortAllServiceUnavailableCode is the HTTP code returned for type GetConfigPortAllServiceUnavailable
const GetConfigPortAllServiceUnavailableCode int = 503

/*
GetConfigPortAllServiceUnavailable Maintenance mode

swagger:response getConfigPortAllServiceUnavailable
*/
type GetConfigPortAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigPortAllServiceUnavailable creates GetConfigPortAllServiceUnavailable with default headers values
func NewGetConfigPortAllServiceUnavailable() *GetConfigPortAllServiceUnavailable {

	return &GetConfigPortAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config port all service unavailable response
func (o *GetConfigPortAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigPortAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config port all service unavailable response
func (o *GetConfigPortAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigPortAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
