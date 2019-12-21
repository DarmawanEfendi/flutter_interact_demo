// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gihub.com/DarmawanEfendi/flutter_interact_demo/models"
)

// DeleteProfileFavouriteOKCode is the HTTP code returned for type DeleteProfileFavouriteOK
const DeleteProfileFavouriteOKCode int = 200

/*DeleteProfileFavouriteOK successful operation

swagger:response deleteProfileFavouriteOK
*/
type DeleteProfileFavouriteOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProfileResponse `json:"body,omitempty"`
}

// NewDeleteProfileFavouriteOK creates DeleteProfileFavouriteOK with default headers values
func NewDeleteProfileFavouriteOK() *DeleteProfileFavouriteOK {

	return &DeleteProfileFavouriteOK{}
}

// WithPayload adds the payload to the delete profile favourite o k response
func (o *DeleteProfileFavouriteOK) WithPayload(payload *models.ProfileResponse) *DeleteProfileFavouriteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete profile favourite o k response
func (o *DeleteProfileFavouriteOK) SetPayload(payload *models.ProfileResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProfileFavouriteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProfileFavouriteBadRequestCode is the HTTP code returned for type DeleteProfileFavouriteBadRequest
const DeleteProfileFavouriteBadRequestCode int = 400

/*DeleteProfileFavouriteBadRequest bad request operation

swagger:response deleteProfileFavouriteBadRequest
*/
type DeleteProfileFavouriteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralErrorResponse `json:"body,omitempty"`
}

// NewDeleteProfileFavouriteBadRequest creates DeleteProfileFavouriteBadRequest with default headers values
func NewDeleteProfileFavouriteBadRequest() *DeleteProfileFavouriteBadRequest {

	return &DeleteProfileFavouriteBadRequest{}
}

// WithPayload adds the payload to the delete profile favourite bad request response
func (o *DeleteProfileFavouriteBadRequest) WithPayload(payload *models.GeneralErrorResponse) *DeleteProfileFavouriteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete profile favourite bad request response
func (o *DeleteProfileFavouriteBadRequest) SetPayload(payload *models.GeneralErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProfileFavouriteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProfileFavouriteNotFoundCode is the HTTP code returned for type DeleteProfileFavouriteNotFound
const DeleteProfileFavouriteNotFoundCode int = 404

/*DeleteProfileFavouriteNotFound not found operation

swagger:response deleteProfileFavouriteNotFound
*/
type DeleteProfileFavouriteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralErrorResponse `json:"body,omitempty"`
}

// NewDeleteProfileFavouriteNotFound creates DeleteProfileFavouriteNotFound with default headers values
func NewDeleteProfileFavouriteNotFound() *DeleteProfileFavouriteNotFound {

	return &DeleteProfileFavouriteNotFound{}
}

// WithPayload adds the payload to the delete profile favourite not found response
func (o *DeleteProfileFavouriteNotFound) WithPayload(payload *models.GeneralErrorResponse) *DeleteProfileFavouriteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete profile favourite not found response
func (o *DeleteProfileFavouriteNotFound) SetPayload(payload *models.GeneralErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProfileFavouriteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProfileFavouriteInternalServerErrorCode is the HTTP code returned for type DeleteProfileFavouriteInternalServerError
const DeleteProfileFavouriteInternalServerErrorCode int = 500

/*DeleteProfileFavouriteInternalServerError internal server error operation

swagger:response deleteProfileFavouriteInternalServerError
*/
type DeleteProfileFavouriteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralErrorResponse `json:"body,omitempty"`
}

// NewDeleteProfileFavouriteInternalServerError creates DeleteProfileFavouriteInternalServerError with default headers values
func NewDeleteProfileFavouriteInternalServerError() *DeleteProfileFavouriteInternalServerError {

	return &DeleteProfileFavouriteInternalServerError{}
}

// WithPayload adds the payload to the delete profile favourite internal server error response
func (o *DeleteProfileFavouriteInternalServerError) WithPayload(payload *models.GeneralErrorResponse) *DeleteProfileFavouriteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete profile favourite internal server error response
func (o *DeleteProfileFavouriteInternalServerError) SetPayload(payload *models.GeneralErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProfileFavouriteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
