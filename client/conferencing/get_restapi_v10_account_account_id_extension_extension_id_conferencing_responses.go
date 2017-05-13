package conferencing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingReader is a Reader for the GetRestapiV10AccountAccountIDExtensionExtensionIDConferencing structure.
type GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault creates a GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault with default headers values
func NewGetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault(code int) *GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault {
	return &GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault{
		_statusCode: code,
	}
}

/*GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault handles this case with default header values.

OK
*/
type GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault struct {
	_statusCode int

	Payload *models.ConferencingInfo
}

// Code gets the status code for the get restapi v10 account account ID extension extension ID conferencing default response
func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault) Code() int {
	return o._statusCode
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault) Error() string {
	return fmt.Sprintf("[GET /restapi/v1.0/account/{accountId}/extension/{extensionId}/conferencing][%d] GetRestapiV10AccountAccountIDExtensionExtensionIDConferencing default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestapiV10AccountAccountIDExtensionExtensionIDConferencingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConferencingInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
