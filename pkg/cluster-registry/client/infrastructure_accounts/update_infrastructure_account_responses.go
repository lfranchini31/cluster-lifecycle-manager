// Code generated by go-swagger; DO NOT EDIT.

package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// UpdateInfrastructureAccountReader is a Reader for the UpdateInfrastructureAccount structure.
type UpdateInfrastructureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInfrastructureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInfrastructureAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateInfrastructureAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInfrastructureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInfrastructureAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInfrastructureAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateInfrastructureAccountOK creates a UpdateInfrastructureAccountOK with default headers values
func NewUpdateInfrastructureAccountOK() *UpdateInfrastructureAccountOK {
	return &UpdateInfrastructureAccountOK{}
}

/*UpdateInfrastructureAccountOK handles this case with default header values.

The infrastructure account update request is accepted
*/
type UpdateInfrastructureAccountOK struct {
	Payload *models.InfrastructureAccount
}

func (o *UpdateInfrastructureAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /infrastructure-accounts/{account_id}][%d] updateInfrastructureAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateInfrastructureAccountOK) GetPayload() *models.InfrastructureAccount {
	return o.Payload
}

func (o *UpdateInfrastructureAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfrastructureAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfrastructureAccountUnauthorized creates a UpdateInfrastructureAccountUnauthorized with default headers values
func NewUpdateInfrastructureAccountUnauthorized() *UpdateInfrastructureAccountUnauthorized {
	return &UpdateInfrastructureAccountUnauthorized{}
}

/*UpdateInfrastructureAccountUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateInfrastructureAccountUnauthorized struct {
}

func (o *UpdateInfrastructureAccountUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /infrastructure-accounts/{account_id}][%d] updateInfrastructureAccountUnauthorized ", 401)
}

func (o *UpdateInfrastructureAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInfrastructureAccountForbidden creates a UpdateInfrastructureAccountForbidden with default headers values
func NewUpdateInfrastructureAccountForbidden() *UpdateInfrastructureAccountForbidden {
	return &UpdateInfrastructureAccountForbidden{}
}

/*UpdateInfrastructureAccountForbidden handles this case with default header values.

Forbidden
*/
type UpdateInfrastructureAccountForbidden struct {
}

func (o *UpdateInfrastructureAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /infrastructure-accounts/{account_id}][%d] updateInfrastructureAccountForbidden ", 403)
}

func (o *UpdateInfrastructureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInfrastructureAccountNotFound creates a UpdateInfrastructureAccountNotFound with default headers values
func NewUpdateInfrastructureAccountNotFound() *UpdateInfrastructureAccountNotFound {
	return &UpdateInfrastructureAccountNotFound{}
}

/*UpdateInfrastructureAccountNotFound handles this case with default header values.

InfrastructureAccount not found
*/
type UpdateInfrastructureAccountNotFound struct {
}

func (o *UpdateInfrastructureAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /infrastructure-accounts/{account_id}][%d] updateInfrastructureAccountNotFound ", 404)
}

func (o *UpdateInfrastructureAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInfrastructureAccountInternalServerError creates a UpdateInfrastructureAccountInternalServerError with default headers values
func NewUpdateInfrastructureAccountInternalServerError() *UpdateInfrastructureAccountInternalServerError {
	return &UpdateInfrastructureAccountInternalServerError{}
}

/*UpdateInfrastructureAccountInternalServerError handles this case with default header values.

Unexpected error
*/
type UpdateInfrastructureAccountInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateInfrastructureAccountInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /infrastructure-accounts/{account_id}][%d] updateInfrastructureAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInfrastructureAccountInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInfrastructureAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
