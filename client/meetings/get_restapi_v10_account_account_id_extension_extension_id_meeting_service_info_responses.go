package meetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfo structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault struct {
	_statusCode int

	Payload *models.MeetingServiceInfo
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID meeting service info default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/service-info][%d] GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDMeetingServiceInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeetingServiceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}