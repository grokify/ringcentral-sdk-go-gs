package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostRestapiOauthRevokeReader is a Reader for the PostRestapiOauthRevoke structure.
type PostRestapiOauthRevokeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRestapiOauthRevokeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostRestapiOauthRevokeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostRestapiOauthRevokeDefault creates a PostRestapiOauthRevokeDefault with default headers values
func NewPostRestapiOauthRevokeDefault(code int) *PostRestapiOauthRevokeDefault {
	return &PostRestapiOauthRevokeDefault{
		_statusCode: code,
	}
}

/*PostRestapiOauthRevokeDefault handles this case with default header values.

OK
*/
type PostRestapiOauthRevokeDefault struct {
	_statusCode int
}

// Code gets the status code for the post restapi oauth revoke default response
func (o *PostRestapiOauthRevokeDefault) Code() int {
	return o._statusCode
}

func (o *PostRestapiOauthRevokeDefault) Error() string {
	return fmt.Sprintf("[POST /restapi/oauth/revoke][%d] PostRestapiOauthRevoke default ", o._statusCode)
}

func (o *PostRestapiOauthRevokeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostRestapiOauthRevokeBody post restapi oauth revoke body
swagger:model PostRestapiOauthRevokeBody
*/
type PostRestapiOauthRevokeBody struct {

	// Active access or refresh token to be revoked
	Token string `json:"token,omitempty"`
}