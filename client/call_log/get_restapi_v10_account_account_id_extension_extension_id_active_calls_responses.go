package call_log

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

// GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault struct {
	_statusCode int

	Payload GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID active calls default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}/active-calls][%d] GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody get restapi v10 account account ID extension extension ID active calls default body
swagger:model GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody struct {

	// Information on navigation
	// Required: true
	Navigation *models.NavigationInfo `json:"navigation"`

	// Information on paging
	// Required: true
	Paging *models.PagingInfo `json:"paging"`

	// List of call log records
	// Required: true
	Records []*models.CallLogRecord `json:"records"`
}

// Validate validates this get restapi v10 account account ID extension extension ID active calls default body
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody) validateNavigation(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls default"+"."+"navigation", "body", o.Navigation); err != nil {
		return err
	}

	if o.Navigation != nil {

		if err := o.Navigation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls default" + "." + "navigation")
			}
			return err
		}
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody) validatePaging(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls default"+"."+"paging", "body", o.Paging); err != nil {
		return err
	}

	if o.Paging != nil {

		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls default" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCallsDefaultBody) validateRecords(formats strfmt.Registry) error {

	if err := validate.Required("GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls default"+"."+"records", "body", o.Records); err != nil {
		return err
	}

	for i := 0; i < len(o.Records); i++ {

		if swag.IsZero(o.Records[i]) { // not required
			continue
		}

		if o.Records[i] != nil {

			if err := o.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetRestapiV10AccountAccountIDExtensionExtensionIDActiveCalls default" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
