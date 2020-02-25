// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// DeleteLoadBalancersIDReader is a Reader for the DeleteLoadBalancersID structure.
type DeleteLoadBalancersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLoadBalancersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteLoadBalancersIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteLoadBalancersIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLoadBalancersIDNoContent creates a DeleteLoadBalancersIDNoContent with default headers values
func NewDeleteLoadBalancersIDNoContent() *DeleteLoadBalancersIDNoContent {
	return &DeleteLoadBalancersIDNoContent{}
}

/*DeleteLoadBalancersIDNoContent handles this case with default header values.

The load balancer deletion request was accepted.
*/
type DeleteLoadBalancersIDNoContent struct {
}

func (o *DeleteLoadBalancersIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /load_balancers/{id}][%d] deleteLoadBalancersIdNoContent ", 204)
}

func (o *DeleteLoadBalancersIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLoadBalancersIDNotFound creates a DeleteLoadBalancersIDNotFound with default headers values
func NewDeleteLoadBalancersIDNotFound() *DeleteLoadBalancersIDNotFound {
	return &DeleteLoadBalancersIDNotFound{}
}

/*DeleteLoadBalancersIDNotFound handles this case with default header values.

A load balancer with the specified identifier could not be found.
*/
type DeleteLoadBalancersIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteLoadBalancersIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /load_balancers/{id}][%d] deleteLoadBalancersIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLoadBalancersIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}