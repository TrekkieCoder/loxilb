// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigBgpPolicyDefinitionsAllOKCode is the HTTP code returned for type GetConfigBgpPolicyDefinitionsAllOK
const GetConfigBgpPolicyDefinitionsAllOKCode int = 200

/*
GetConfigBgpPolicyDefinitionsAllOK OK

swagger:response getConfigBgpPolicyDefinitionsAllOK
*/
type GetConfigBgpPolicyDefinitionsAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigBgpPolicyDefinitionsAllOKBody `json:"body,omitempty"`
}

// NewGetConfigBgpPolicyDefinitionsAllOK creates GetConfigBgpPolicyDefinitionsAllOK with default headers values
func NewGetConfigBgpPolicyDefinitionsAllOK() *GetConfigBgpPolicyDefinitionsAllOK {

	return &GetConfigBgpPolicyDefinitionsAllOK{}
}

// WithPayload adds the payload to the get config bgp policy definitions all o k response
func (o *GetConfigBgpPolicyDefinitionsAllOK) WithPayload(payload *GetConfigBgpPolicyDefinitionsAllOKBody) *GetConfigBgpPolicyDefinitionsAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp policy definitions all o k response
func (o *GetConfigBgpPolicyDefinitionsAllOK) SetPayload(payload *GetConfigBgpPolicyDefinitionsAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpPolicyDefinitionsAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpPolicyDefinitionsAllUnauthorizedCode is the HTTP code returned for type GetConfigBgpPolicyDefinitionsAllUnauthorized
const GetConfigBgpPolicyDefinitionsAllUnauthorizedCode int = 401

/*
GetConfigBgpPolicyDefinitionsAllUnauthorized Invalid authentication credentials

swagger:response getConfigBgpPolicyDefinitionsAllUnauthorized
*/
type GetConfigBgpPolicyDefinitionsAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpPolicyDefinitionsAllUnauthorized creates GetConfigBgpPolicyDefinitionsAllUnauthorized with default headers values
func NewGetConfigBgpPolicyDefinitionsAllUnauthorized() *GetConfigBgpPolicyDefinitionsAllUnauthorized {

	return &GetConfigBgpPolicyDefinitionsAllUnauthorized{}
}

// WithPayload adds the payload to the get config bgp policy definitions all unauthorized response
func (o *GetConfigBgpPolicyDefinitionsAllUnauthorized) WithPayload(payload *models.Error) *GetConfigBgpPolicyDefinitionsAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp policy definitions all unauthorized response
func (o *GetConfigBgpPolicyDefinitionsAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpPolicyDefinitionsAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpPolicyDefinitionsAllInternalServerErrorCode is the HTTP code returned for type GetConfigBgpPolicyDefinitionsAllInternalServerError
const GetConfigBgpPolicyDefinitionsAllInternalServerErrorCode int = 500

/*
GetConfigBgpPolicyDefinitionsAllInternalServerError Internal service error

swagger:response getConfigBgpPolicyDefinitionsAllInternalServerError
*/
type GetConfigBgpPolicyDefinitionsAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpPolicyDefinitionsAllInternalServerError creates GetConfigBgpPolicyDefinitionsAllInternalServerError with default headers values
func NewGetConfigBgpPolicyDefinitionsAllInternalServerError() *GetConfigBgpPolicyDefinitionsAllInternalServerError {

	return &GetConfigBgpPolicyDefinitionsAllInternalServerError{}
}

// WithPayload adds the payload to the get config bgp policy definitions all internal server error response
func (o *GetConfigBgpPolicyDefinitionsAllInternalServerError) WithPayload(payload *models.Error) *GetConfigBgpPolicyDefinitionsAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp policy definitions all internal server error response
func (o *GetConfigBgpPolicyDefinitionsAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpPolicyDefinitionsAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigBgpPolicyDefinitionsAllServiceUnavailableCode is the HTTP code returned for type GetConfigBgpPolicyDefinitionsAllServiceUnavailable
const GetConfigBgpPolicyDefinitionsAllServiceUnavailableCode int = 503

/*
GetConfigBgpPolicyDefinitionsAllServiceUnavailable Maintanence mode

swagger:response getConfigBgpPolicyDefinitionsAllServiceUnavailable
*/
type GetConfigBgpPolicyDefinitionsAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigBgpPolicyDefinitionsAllServiceUnavailable creates GetConfigBgpPolicyDefinitionsAllServiceUnavailable with default headers values
func NewGetConfigBgpPolicyDefinitionsAllServiceUnavailable() *GetConfigBgpPolicyDefinitionsAllServiceUnavailable {

	return &GetConfigBgpPolicyDefinitionsAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config bgp policy definitions all service unavailable response
func (o *GetConfigBgpPolicyDefinitionsAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigBgpPolicyDefinitionsAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config bgp policy definitions all service unavailable response
func (o *GetConfigBgpPolicyDefinitionsAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigBgpPolicyDefinitionsAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}