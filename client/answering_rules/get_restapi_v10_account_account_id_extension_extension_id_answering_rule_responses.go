package answering_rules

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

// GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault struct {
	_statusCode int

	Payload GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID answering rule default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule][%d] GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody get restapi v10 account account ID extension extension ID answering rule default body
swagger:model GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody struct {

	// Information on navigation
	// Required: true
	Navigation *models.NavigationInfo `json:"navigation"`

	// Information on paging
	// Required: true
	Paging *models.PagingInfo `json:"paging"`

	// List of answering rules
	// Required: true
	Records []*models.AnsweringRuleInfo `json:"records"`

	// Canonical URI of an answering rule resource
	// Required: true
	URI *string `json:"uri"`
}

// Validate validates this get restapi v10 account account ID extension extension ID answering rule default body
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNavigation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePaging(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRecords(formats); err != nil {
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

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody) validateNavigation(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default"+"."+"navigation", "body", o.Navigation); err != nil {
		return err
	}

	if o.Navigation != nil {

		if err := o.Navigation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default" + "." + "navigation")
			}
			return err
		}
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody) validatePaging(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default"+"."+"paging", "body", o.Paging); err != nil {
		return err
	}

	if o.Paging != nil {

		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody) validateRecords(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default"+"."+"records", "body", o.Records); err != nil {
		return err
	}

	for i := 0; i < len(o.Records); i++ {

		if swag.IsZero(o.Records[i]) { // not required
			continue
		}

		if o.Records[i] != nil {

			if err := o.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleDefaultBody) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRule default"+"."+"uri", "body", o.URI); err != nil {
		return err
	}

	return nil
}