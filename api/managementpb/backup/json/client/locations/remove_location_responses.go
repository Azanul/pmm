// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoveLocationReader is a Reader for the RemoveLocation structure.
type RemoveLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveLocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveLocationOK creates a RemoveLocationOK with default headers values
func NewRemoveLocationOK() *RemoveLocationOK {
	return &RemoveLocationOK{}
}

/* RemoveLocationOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveLocationOK struct {
	Payload interface{}
}

func (o *RemoveLocationOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/Remove][%d] removeLocationOk  %+v", 200, o.Payload)
}

func (o *RemoveLocationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLocationDefault creates a RemoveLocationDefault with default headers values
func NewRemoveLocationDefault(code int) *RemoveLocationDefault {
	return &RemoveLocationDefault{
		_statusCode: code,
	}
}

/* RemoveLocationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveLocationDefault struct {
	_statusCode int

	Payload *RemoveLocationDefaultBody
}

// Code gets the status code for the remove location default response
func (o *RemoveLocationDefault) Code() int {
	return o._statusCode
}

func (o *RemoveLocationDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/Remove][%d] RemoveLocation default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveLocationDefault) GetPayload() *RemoveLocationDefaultBody {
	return o.Payload
}

func (o *RemoveLocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RemoveLocationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveLocationBody remove location body
swagger:model RemoveLocationBody
*/
type RemoveLocationBody struct {
	// Machine-readable ID.
	LocationID string `json:"location_id,omitempty"`

	// Force mode
	Force bool `json:"force,omitempty"`
}

// Validate validates this remove location body
func (o *RemoveLocationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove location body based on context it is used
func (o *RemoveLocationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveLocationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveLocationBody) UnmarshalBinary(b []byte) error {
	var res RemoveLocationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveLocationDefaultBody remove location default body
swagger:model RemoveLocationDefaultBody
*/
type RemoveLocationDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RemoveLocationDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this remove location default body
func (o *RemoveLocationDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveLocationDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveLocation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveLocation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remove location default body based on the context it is used
func (o *RemoveLocationDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveLocationDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveLocation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveLocation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveLocationDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveLocationDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveLocationDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveLocationDefaultBodyDetailsItems0 remove location default body details items0
swagger:model RemoveLocationDefaultBodyDetailsItems0
*/
type RemoveLocationDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this remove location default body details items0
func (o *RemoveLocationDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove location default body details items0 based on context it is used
func (o *RemoveLocationDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveLocationDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveLocationDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RemoveLocationDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
