package ring_out

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDReader is a Reader for the DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID structure.
type DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault creates a DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault with default headers values
func NewDeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault(code int) *DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault {
	return &DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault{
		_statusCode: code,
	}
}

/*DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault handles this case with default header values.

OK
*/
type DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete restapi v10 account account ID extension extension ID ringout ringout ID default response
func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId}][%d] DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutID default ", o._statusCode)
}

func (o *DeleteRestapiV10AccountAccountIDExtensionExtensionIDRingoutRingoutIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
