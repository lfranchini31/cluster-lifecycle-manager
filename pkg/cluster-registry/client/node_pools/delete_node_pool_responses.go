// Code generated by go-swagger; DO NOT EDIT.

package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// DeleteNodePoolReader is a Reader for the DeleteNodePool structure.
type DeleteNodePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNodePoolNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNodePoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNodePoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNodePoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNodePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNodePoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNodePoolNoContent creates a DeleteNodePoolNoContent with default headers values
func NewDeleteNodePoolNoContent() *DeleteNodePoolNoContent {
	return &DeleteNodePoolNoContent{}
}

/*DeleteNodePoolNoContent handles this case with default header values.

Node pool deleted.
*/
type DeleteNodePoolNoContent struct {
}

func (o *DeleteNodePoolNoContent) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] deleteNodePoolNoContent ", 204)
}

func (o *DeleteNodePoolNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolBadRequest creates a DeleteNodePoolBadRequest with default headers values
func NewDeleteNodePoolBadRequest() *DeleteNodePoolBadRequest {
	return &DeleteNodePoolBadRequest{}
}

/*DeleteNodePoolBadRequest handles this case with default header values.

Invalid request
*/
type DeleteNodePoolBadRequest struct {
	Payload *models.Error
}

func (o *DeleteNodePoolBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] deleteNodePoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNodePoolBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNodePoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodePoolUnauthorized creates a DeleteNodePoolUnauthorized with default headers values
func NewDeleteNodePoolUnauthorized() *DeleteNodePoolUnauthorized {
	return &DeleteNodePoolUnauthorized{}
}

/*DeleteNodePoolUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNodePoolUnauthorized struct {
}

func (o *DeleteNodePoolUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] deleteNodePoolUnauthorized ", 401)
}

func (o *DeleteNodePoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolForbidden creates a DeleteNodePoolForbidden with default headers values
func NewDeleteNodePoolForbidden() *DeleteNodePoolForbidden {
	return &DeleteNodePoolForbidden{}
}

/*DeleteNodePoolForbidden handles this case with default header values.

Forbidden
*/
type DeleteNodePoolForbidden struct {
}

func (o *DeleteNodePoolForbidden) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] deleteNodePoolForbidden ", 403)
}

func (o *DeleteNodePoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolNotFound creates a DeleteNodePoolNotFound with default headers values
func NewDeleteNodePoolNotFound() *DeleteNodePoolNotFound {
	return &DeleteNodePoolNotFound{}
}

/*DeleteNodePoolNotFound handles this case with default header values.

Node pool not found
*/
type DeleteNodePoolNotFound struct {
}

func (o *DeleteNodePoolNotFound) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] deleteNodePoolNotFound ", 404)
}

func (o *DeleteNodePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolInternalServerError creates a DeleteNodePoolInternalServerError with default headers values
func NewDeleteNodePoolInternalServerError() *DeleteNodePoolInternalServerError {
	return &DeleteNodePoolInternalServerError{}
}

/*DeleteNodePoolInternalServerError handles this case with default header values.

Unexpected error
*/
type DeleteNodePoolInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteNodePoolInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}][%d] deleteNodePoolInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNodePoolInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNodePoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
