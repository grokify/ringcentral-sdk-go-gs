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

// GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImage structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault struct {
	_statusCode int

	Payload *models.Binary
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID profile image default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image][%d] GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImage default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Binary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
