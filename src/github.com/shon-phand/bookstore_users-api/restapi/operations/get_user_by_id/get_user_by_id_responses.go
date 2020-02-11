// Code generated by go-swagger; DO NOT EDIT.

package get_user_by_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shon-phand/bookstore_users-api/models"
)

// GetUserByIDOKCode is the HTTP code returned for type GetUserByIDOK
const GetUserByIDOKCode int = 200

/*GetUserByIDOK user fetched successfully

swagger:response getUserByIdOK
*/
type GetUserByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserByIDOK creates GetUserByIDOK with default headers values
func NewGetUserByIDOK() *GetUserByIDOK {

	return &GetUserByIDOK{}
}

// WithPayload adds the payload to the get user by Id o k response
func (o *GetUserByIDOK) WithPayload(payload *models.User) *GetUserByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id o k response
func (o *GetUserByIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUserByIDDefault error response

swagger:response getUserByIdDefault
*/
type GetUserByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.RestErr `json:"body,omitempty"`
}

// NewGetUserByIDDefault creates GetUserByIDDefault with default headers values
func NewGetUserByIDDefault(code int) *GetUserByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user by Id default response
func (o *GetUserByIDDefault) WithStatusCode(code int) *GetUserByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user by Id default response
func (o *GetUserByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user by Id default response
func (o *GetUserByIDDefault) WithPayload(payload *models.RestErr) *GetUserByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id default response
func (o *GetUserByIDDefault) SetPayload(payload *models.RestErr) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}