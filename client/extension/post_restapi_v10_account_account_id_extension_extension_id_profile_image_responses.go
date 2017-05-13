package extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageReader is a Reader for the PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImage structure.
type PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault creates a PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault with default headers values
func NewPostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault(code int) *PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault {
	return &PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault{
		_statusCode: code,
	}
}

/*PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault handles this case with default header values.

OK
*/
type PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault struct {
	_statusCode int
}

// Code gets the status code for the post restapi v10 account account ID extension extension ID profile image default response
func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault) Code() int {
	return o._statusCode
}

func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault) Error() string {
	return fmt.Sprintf("[POST /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image][%d] PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImage default ", o._statusCode)
}

func (o *PostRestapiV10AccountAccountIDExtensionExtensionIDProfileImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}