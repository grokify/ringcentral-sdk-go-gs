package answering_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/grokify/go-ringcentral/models"
)

// PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDReader is a Reader for the PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleID structure.
type PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault creates a PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault with default headers values
func NewPutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault(code int) *PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault {
	return &PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault{
		_statusCode: code,
	}
}

/*PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault handles this case with default header values.

OK
*/
type PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault struct {
	_statusCode int

	Payload *models.AnsweringRuleInfo
}

// Code gets the status code for the put restapi v10 account account ID extension extension ID answering rule answering rule ID default response
func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault) Error() string {
	return fmt.Sprintf("[PUT /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{answeringRuleId}][%d] PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnsweringRuleInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDBody put restapi v10 account account ID extension extension ID answering rule answering rule ID body
swagger:model PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDBody
*/
type PutRestapiV10AccountAccountIDExtensionExtensionIDAnsweringRuleAnsweringRuleIDBody struct {

	// Specifies if the answering rule is active or not
	Enabled bool `json:"enabled,omitempty"`

	// Forwarding parameters. Returned if 'ForwardCalls' is specified in 'callHandlingAction'. These settings determine the forwarding numbers to which the call will be forwarded
	Forwarding *models.ForwardingInfo `json:"forwarding,omitempty"`

	// Predefined greetings applied for an answering rule
	Greetings []*models.GreetingInfo `json:"greetings"`

	// Custom name of an answering rule. The maximum number of characters is 64
	Name string `json:"name,omitempty"`
}
