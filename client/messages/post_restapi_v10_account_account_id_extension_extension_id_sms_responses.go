package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// PostRestapiV10AccountAccountIDExtensionExtensionIDSmsReader is a Reader for the PostRestapiV10AccountAccountIDExtensionExtensionIDSms structure.
type PostRestapiV10AccountAccountIDExtensionExtensionIDSmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDSmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault creates a PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault with default headers values
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault(code int) *PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault {
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault{
		_statusCode: code,
	}
}

/*PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault handles this case with default header values.

OK
*/
type PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault struct {
	_statusCode int

	Payload *models.MessageInfo
}

// Code gets the status code for the post restapi v10 account account ID extension extension ID sms default response
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault) Code() int {
	return o._statusCode
}

func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault) Error() string {
	return fmt.Sprintf("[POST /restapi/v1.0/account/{accountId}/extension/{extensionId}/sms][%d] PostRestapiV10AccountAccountIDExtensionExtensionIDSms default  %+v", o._statusCode, o.Payload)
}

func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDSmsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRestapiV10AccountAccountIDExtensionExtensionIDSmsBody post restapi v10 account account ID extension extension ID sms body
swagger:model PostRestapiV10AccountAccountIDExtensionExtensionIDSmsBody
*/
type PostRestapiV10AccountAccountIDExtensionExtensionIDSmsBody struct {

	// Sender of an SMS message. The phoneNumber property must be filled to correspond to one of the account phone numbers which is allowed to send SMS
	From *models.CallerInfo `json:"from,omitempty"`

	// Text of a message. Max length is 1000 symbols (2-byte UTF-16 encoded). If a character is encoded in 4 bytes in UTF-16 it is treated as 2 characters, thus restricting the maximum message length to 500 symbols
	Text string `json:"text,omitempty"`

	// Receiver of an SMS message. The phoneNumber property must be filled
	To []*models.CallerInfo `json:"to"`
}