package authorization_profile

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

// GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheck structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault struct {
	_statusCode int

	Payload GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID authz profile check default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile/check][%d] GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody get restapi v10 account account ID extension extension ID authz profile check default body
swagger:model GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody struct {

	// Information on a permission checked. Returned if successful is 'True'
	// Required: true
	Details *models.PermissionDetailsInfo `json:"details"`

	// List of active scopes for permission. Returned if successful is 'True'
	// Required: true
	Scopes []string `json:"scopes"`

	// Specifies if check result is successful or not
	// Required: true
	Successful *bool `json:"successful"`

	// Canonical URI of a permission resource
	// Required: true
	URI *string `json:"uri"`
}

// Validate validates this get restapi v10 account account ID extension extension ID authz profile check default body
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateScopes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSuccessful(formats); err != nil {
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

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheck default"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	if o.Details != nil {

		if err := o.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheck default" + "." + "details")
			}
			return err
		}
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody) validateScopes(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheck default"+"."+"scopes", "body", o.Scopes); err != nil {
		return err
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody) validateSuccessful(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheck default"+"."+"successful", "body", o.Successful); err != nil {
		return err
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheckDefaultBody) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAuthzProfileCheck default"+"."+"uri", "body", o.URI); err != nil {
		return err
	}

	return nil
}
