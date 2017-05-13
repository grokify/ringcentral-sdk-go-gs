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

// PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDReader is a Reader for the PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageID structure.
type PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault creates a PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault with default headers values
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault(code int) *PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault {
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault{
		_statusCode: code,
	}
}

/*PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault handles this case with default header values.

OK
*/
type PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault struct {
	_statusCode int

	Payload *models.MessageInfo
}

// Code gets the status code for the put restapi v10 account account ID extension extension ID message store message ID default response
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault) Code() int {
	return o._statusCode
}

func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault) Error() string {
	return fmt.Sprintf("[PUT /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId}][%d] PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageID default  %+v", o._statusCode, o.Payload)
}

func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDBody put restapi v10 account account ID extension extension ID message store message ID body
swagger:model PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDBody
*/
type PutRestapiV10AccountAccountIDExtensionExtensionIDMessageStoreMessageIDBody struct {

	// Read status of a message to be changed. Multiple values are accepted
	ReadStatus string `json:"readStatus,omitempty"`
}
