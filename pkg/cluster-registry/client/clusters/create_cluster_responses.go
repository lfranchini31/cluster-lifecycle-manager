// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// CreateClusterReader is a Reader for the CreateCluster structure.
type CreateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateClusterCreated creates a CreateClusterCreated with default headers values
func NewCreateClusterCreated() *CreateClusterCreated {
	return &CreateClusterCreated{}
}

/*CreateClusterCreated handles this case with default header values.

The cluster creation request is accepted
*/
type CreateClusterCreated struct {
	Payload *models.Cluster
}

func (o *CreateClusterCreated) Error() string {
	return fmt.Sprintf("[POST /kubernetes-clusters][%d] createClusterCreated  %+v", 201, o.Payload)
}

func (o *CreateClusterCreated) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *CreateClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterBadRequest creates a CreateClusterBadRequest with default headers values
func NewCreateClusterBadRequest() *CreateClusterBadRequest {
	return &CreateClusterBadRequest{}
}

/*CreateClusterBadRequest handles this case with default header values.

Invalid request
*/
type CreateClusterBadRequest struct {
	Payload *models.Error
}

func (o *CreateClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /kubernetes-clusters][%d] createClusterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterUnauthorized creates a CreateClusterUnauthorized with default headers values
func NewCreateClusterUnauthorized() *CreateClusterUnauthorized {
	return &CreateClusterUnauthorized{}
}

/*CreateClusterUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateClusterUnauthorized struct {
}

func (o *CreateClusterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /kubernetes-clusters][%d] createClusterUnauthorized ", 401)
}

func (o *CreateClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterForbidden creates a CreateClusterForbidden with default headers values
func NewCreateClusterForbidden() *CreateClusterForbidden {
	return &CreateClusterForbidden{}
}

/*CreateClusterForbidden handles this case with default header values.

Forbidden
*/
type CreateClusterForbidden struct {
}

func (o *CreateClusterForbidden) Error() string {
	return fmt.Sprintf("[POST /kubernetes-clusters][%d] createClusterForbidden ", 403)
}

func (o *CreateClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterConflict creates a CreateClusterConflict with default headers values
func NewCreateClusterConflict() *CreateClusterConflict {
	return &CreateClusterConflict{}
}

/*CreateClusterConflict handles this case with default header values.

Conflict, already existing
*/
type CreateClusterConflict struct {
}

func (o *CreateClusterConflict) Error() string {
	return fmt.Sprintf("[POST /kubernetes-clusters][%d] createClusterConflict ", 409)
}

func (o *CreateClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterInternalServerError creates a CreateClusterInternalServerError with default headers values
func NewCreateClusterInternalServerError() *CreateClusterInternalServerError {
	return &CreateClusterInternalServerError{}
}

/*CreateClusterInternalServerError handles this case with default header values.

Unexpected error
*/
type CreateClusterInternalServerError struct {
	Payload *models.Error
}

func (o *CreateClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /kubernetes-clusters][%d] createClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
