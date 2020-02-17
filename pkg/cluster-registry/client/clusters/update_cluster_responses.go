// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateClusterOK creates a UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

/*UpdateClusterOK handles this case with default header values.

The cluster update request is performed and the updated cluster is returned.
*/
type UpdateClusterOK struct {
	Payload *models.Cluster
}

func (o *UpdateClusterOK) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}][%d] updateClusterOK  %+v", 200, o.Payload)
}

func (o *UpdateClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *UpdateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterUnauthorized creates a UpdateClusterUnauthorized with default headers values
func NewUpdateClusterUnauthorized() *UpdateClusterUnauthorized {
	return &UpdateClusterUnauthorized{}
}

/*UpdateClusterUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateClusterUnauthorized struct {
}

func (o *UpdateClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}][%d] updateClusterUnauthorized ", 401)
}

func (o *UpdateClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClusterForbidden creates a UpdateClusterForbidden with default headers values
func NewUpdateClusterForbidden() *UpdateClusterForbidden {
	return &UpdateClusterForbidden{}
}

/*UpdateClusterForbidden handles this case with default header values.

Forbidden
*/
type UpdateClusterForbidden struct {
}

func (o *UpdateClusterForbidden) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}][%d] updateClusterForbidden ", 403)
}

func (o *UpdateClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClusterNotFound creates a UpdateClusterNotFound with default headers values
func NewUpdateClusterNotFound() *UpdateClusterNotFound {
	return &UpdateClusterNotFound{}
}

/*UpdateClusterNotFound handles this case with default header values.

Cluster not found
*/
type UpdateClusterNotFound struct {
}

func (o *UpdateClusterNotFound) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}][%d] updateClusterNotFound ", 404)
}

func (o *UpdateClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClusterInternalServerError creates a UpdateClusterInternalServerError with default headers values
func NewUpdateClusterInternalServerError() *UpdateClusterInternalServerError {
	return &UpdateClusterInternalServerError{}
}

/*UpdateClusterInternalServerError handles this case with default header values.

Unexpected error
*/
type UpdateClusterInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateClusterInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}][%d] updateClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
