package extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// GetRestapiV10AccountAccountIDExtensionExtensionIDReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionID structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDDefault struct {
	_statusCode int

	Payload *models.ExtensionInfo
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}][%d] GetRestapiV10AccountAccountIDExtensionExtensionID default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
