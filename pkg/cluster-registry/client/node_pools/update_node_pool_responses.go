// Code generated by go-swagger; DO NOT EDIT.

package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// UpdateNodePoolReader is a Reader for the UpdateNodePool structure.
type UpdateNodePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNodePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNodePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNodePoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateNodePoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateNodePoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateNodePoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNodePoolOK creates a UpdateNodePoolOK with default headers values
func NewUpdateNodePoolOK() *UpdateNodePoolOK {
	return &UpdateNodePoolOK{}
}

/*UpdateNodePoolOK handles this case with default header values.

The node pool update request is accepted.
*/
type UpdateNodePoolOK struct {
	Payload *models.NodePool
}

func (o *UpdateNodePoolOK) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] updateNodePoolOK  %+v", 200, o.Payload)
}

func (o *UpdateNodePoolOK) GetPayload() *models.NodePool {
	return o.Payload
}

func (o *UpdateNodePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodePool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNodePoolBadRequest creates a UpdateNodePoolBadRequest with default headers values
func NewUpdateNodePoolBadRequest() *UpdateNodePoolBadRequest {
	return &UpdateNodePoolBadRequest{}
}

/*UpdateNodePoolBadRequest handles this case with default header values.

Invalid request
*/
type UpdateNodePoolBadRequest struct {
	Payload *models.Error
}

func (o *UpdateNodePoolBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] updateNodePoolBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNodePoolBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateNodePoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNodePoolUnauthorized creates a UpdateNodePoolUnauthorized with default headers values
func NewUpdateNodePoolUnauthorized() *UpdateNodePoolUnauthorized {
	return &UpdateNodePoolUnauthorized{}
}

/*UpdateNodePoolUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateNodePoolUnauthorized struct {
}

func (o *UpdateNodePoolUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] updateNodePoolUnauthorized ", 401)
}

func (o *UpdateNodePoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNodePoolForbidden creates a UpdateNodePoolForbidden with default headers values
func NewUpdateNodePoolForbidden() *UpdateNodePoolForbidden {
	return &UpdateNodePoolForbidden{}
}

/*UpdateNodePoolForbidden handles this case with default header values.

Forbidden
*/
type UpdateNodePoolForbidden struct {
}

func (o *UpdateNodePoolForbidden) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] updateNodePoolForbidden ", 403)
}

func (o *UpdateNodePoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNodePoolInternalServerError creates a UpdateNodePoolInternalServerError with default headers values
func NewUpdateNodePoolInternalServerError() *UpdateNodePoolInternalServerError {
	return &UpdateNodePoolInternalServerError{}
}

/*UpdateNodePoolInternalServerError handles this case with default header values.

Unexpected error
*/
type UpdateNodePoolInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateNodePoolInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] updateNodePoolInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateNodePoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateNodePoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
