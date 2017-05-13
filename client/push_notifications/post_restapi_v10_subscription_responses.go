package push_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// PostRestapiV10SubscriptionReader is a Reader for the PostRestapiV10Subscription structure.
type PostRestapiV10SubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRestapiV10SubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostRestapiV10SubscriptionDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostRestapiV10SubscriptionDefault creates a PostRestapiV10SubscriptionDefault with default headers values
func NewPostRestapiV10SubscriptionDefault(code int) *PostRestapiV10SubscriptionDefault {
	return &PostRestapiV10SubscriptionDefault{
		_statusCode: code,
	}
}

/*PostRestapiV10SubscriptionDefault handles this case with default header values.

OK
*/
type PostRestapiV10SubscriptionDefault struct {
	_statusCode int

	Payload *models.SubscriptionInfo
}

// Code gets the status code for the post restapi v10 subscription default response
func (o *PostRestapiV10SubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *PostRestapiV10SubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /restapi/v1.0/subscription][%d] PostRestapiV10Subscription default  %+v", o._statusCode, o.Payload)
}

func (o *PostRestapiV10SubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRestapiV10SubscriptionBody post restapi v10 subscription body
swagger:model PostRestapiV10SubscriptionBody
*/
type PostRestapiV10SubscriptionBody struct {

	// Notification delivery settings
	DeliveryMode *models.SubscriptionRequestDeliveryMode `json:"deliveryMode,omitempty"`

	// Mandatory. Collection of URIs to API resources (see Event Types for details). For APNS transport type only message event filter is available
	EventFilters []string `json:"eventFilters"`
}
