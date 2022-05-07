// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/huavanthong/microservice-golang/product-api/sdk/models"
)

// ListSingleProductReader is a Reader for the ListSingleProduct structure.
type ListSingleProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSingleProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSingleProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListSingleProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSingleProductOK creates a ListSingleProductOK with default headers values
func NewListSingleProductOK() *ListSingleProductOK {
	return &ListSingleProductOK{}
}

/*ListSingleProductOK handles this case with default header values.

Data structure representing a single product
*/
type ListSingleProductOK struct {
	Payload *models.Product
}

func (o *ListSingleProductOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductOK  %+v", 200, o.Payload)
}

func (o *ListSingleProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *ListSingleProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSingleProductNotFound creates a ListSingleProductNotFound with default headers values
func NewListSingleProductNotFound() *ListSingleProductNotFound {
	return &ListSingleProductNotFound{}
}

/*ListSingleProductNotFound handles this case with default header values.

Generic error message returned as a string
*/
type ListSingleProductNotFound struct {
	Payload *models.GenericError
}

func (o *ListSingleProductNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductNotFound  %+v", 404, o.Payload)
}

func (o *ListSingleProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ListSingleProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
