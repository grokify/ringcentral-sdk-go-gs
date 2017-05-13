package ordering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// PostRestapiV10AccountAccountIDOrderReader is a Reader for the PostRestapiV10AccountAccountIDOrder structure.
type PostRestapiV10AccountAccountIDOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRestapiV10AccountAccountIDOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostRestapiV10AccountAccountIDOrderDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostRestapiV10AccountAccountIDOrderDefault creates a PostRestapiV10AccountAccountIDOrderDefault with default headers values
func NewPostRestapiV10AccountAccountIDOrderDefault(code int) *PostRestapiV10AccountAccountIDOrderDefault {
	return &PostRestapiV10AccountAccountIDOrderDefault{
		_statusCode: code,
	}
}

/*PostRestapiV10AccountAccountIDOrderDefault handles this case with default header values.

OK
*/
type PostRestapiV10AccountAccountIDOrderDefault struct {
	_statusCode int

	Payload PostRestapiV10AccountAccountIDOrderDefaultBody
}

// Code gets the status code for the post restapi v10 account account ID order default response
func (o *PostRestapiV10AccountAccountIDOrderDefault) Code() int {
	return o._statusCode
}

func (o *PostRestapiV10AccountAccountIDOrderDefault) Error() string {
	return fmt.Sprintf("[POST /restapi/v1.0/account/{accountId}/order][%d] PostRestapiV10AccountAccountIDOrder default  %+v", o._statusCode, o.Payload)
}

func (o *PostRestapiV10AccountAccountIDOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRestapiV10AccountAccountIDOrderBody post restapi v10 account account ID order body
swagger:model PostRestapiV10AccountAccountIDOrderBody
*/
type PostRestapiV10AccountAccountIDOrderBody struct {

	// List of devices to order
	Devices []*models.DeviceInfo `json:"devices"`
}

/*PostRestapiV10AccountAccountIDOrderDefaultBody post restapi v10 account account ID order default body
swagger:model PostRestapiV10AccountAccountIDOrderDefaultBody
*/
type PostRestapiV10AccountAccountIDOrderDefaultBody struct {

	// List of the ordered devices
	// Required: true
	Devices []*models.DeviceInfo `json:"devices"`
}

// Validate validates this post restapi v10 account account ID order default body
func (o *PostRestapiV10AccountAccountIDOrderDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostRestapiV10AccountAccountIDOrderDefaultBody) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("PostRestapiV10AccountAccountIDOrder default"+"."+"devices", "body", o.Devices); err != nil {
		return err
	}

	for i := 0; i < len(o.Devices); i++ {

		if swag.IsZero(o.Devices[i]) { // not required
			continue
		}

		if o.Devices[i] != nil {

			if err := o.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PostRestapiV10AccountAccountIDOrder default" + "." + "devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}