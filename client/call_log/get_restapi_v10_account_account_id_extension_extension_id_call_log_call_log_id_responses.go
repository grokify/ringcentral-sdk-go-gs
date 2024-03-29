package call_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogID structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault struct {
	_statusCode int

	Payload *models.CallLogInfo
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID call log call log ID default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log/{callLogId}][%d] GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogID default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDCallLogCallLogIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallLogInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
