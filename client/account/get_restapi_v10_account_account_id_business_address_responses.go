package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// GetRestapiV10AccountAccountIDBusinessAddressReader is a Reader for the GetRestapiV10AccountAccountIDBusinessAddress structure.
type GetRestapiV10AccountAccountIDBusinessAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDBusinessAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDBusinessAddressDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDBusinessAddressDefault creates a GetRestapiV10AccountAccountIDBusinessAddressDefault with default headers values
func NewGetRestapiV10AccountAccountIDBusinessAddressDefault(code int) *GetRestapiV10AccountAccountIDBusinessAddressDefault {
	return &GetRestapiV10AccountAccountIDBusinessAddressDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDBusinessAddressDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDBusinessAddressDefault struct {
	_statusCode int

	Payload GetRestapiV10AccountAccountIDBusinessAddressDefaultBody
}

// Code gets the status code for the get restapi v10 account account ID business address default response
func (o *GetRestapiV10AccountAccountIDBusinessAddressDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDBusinessAddressDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/business-address][%d] GetRestapiV10AccountAccountIDBusinessAddress default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDBusinessAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRestapiV10AccountAccountIDBusinessAddressDefaultBody get restapi v10 account account ID business address default body
swagger:model GetRestapiV10AccountAccountIDBusinessAddressDefaultBody
*/
type GetRestapiV10AccountAccountIDBusinessAddressDefaultBody struct {

	// Company business address
	// Required: true
	BusinessAddress *models.BusinessAddressInfo `json:"businessAddress"`

	// Company business name
	// Required: true
	Company *string `json:"company"`

	// Company business email address
	// Required: true
	Email *string `json:"email"`

	// Canonical URI of the business address resource
	// Required: true
	URI *string `json:"uri"`
}

// Validate validates this get restapi v10 account account ID business address default body
func (o *GetRestapiV10AccountAccountIDBusinessAddressDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBusinessAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCompany(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateURI(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRestapiV10AccountAccountIDBusinessAddressDefaultBody) validateBusinessAddress(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDBusinessAddress default"+"."+"businessAddress", "body", o.BusinessAddress); err != nil {
		return err
	}

	if o.BusinessAddress != nil {

		if err := o.BusinessAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GetRestapiV10AccountAccountIDBusinessAddress default" + "." + "businessAddress")
			}
			return err
		}
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDBusinessAddressDefaultBody) validateCompany(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDBusinessAddress default"+"."+"company", "body", o.Company); err != nil {
		return err
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDBusinessAddressDefaultBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDBusinessAddress default"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDBusinessAddressDefaultBody) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDBusinessAddress default"+"."+"uri", "body", o.URI); err != nil {
		return err
	}

	return nil
}