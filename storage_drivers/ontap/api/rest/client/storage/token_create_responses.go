// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// TokenCreateReader is a Reader for the TokenCreate structure.
type TokenCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTokenCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTokenCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTokenCreateCreated creates a TokenCreateCreated with default headers values
func NewTokenCreateCreated() *TokenCreateCreated {
	return &TokenCreateCreated{}
}

/* TokenCreateCreated describes a response with status code 201, with default header values.

Created
*/
type TokenCreateCreated struct {
	Payload *models.TokenResponse
}

func (o *TokenCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/file/clone/tokens][%d] tokenCreateCreated  %+v", 201, o.Payload)
}
func (o *TokenCreateCreated) GetPayload() *models.TokenResponse {
	return o.Payload
}

func (o *TokenCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenCreateDefault creates a TokenCreateDefault with default headers values
func NewTokenCreateDefault(code int) *TokenCreateDefault {
	return &TokenCreateDefault{
		_statusCode: code,
	}
}

/* TokenCreateDefault describes a response with status code -1, with default header values.

Error
*/
type TokenCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the token create default response
func (o *TokenCreateDefault) Code() int {
	return o._statusCode
}

func (o *TokenCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/file/clone/tokens][%d] token_create default  %+v", o._statusCode, o.Payload)
}
func (o *TokenCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TokenCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
