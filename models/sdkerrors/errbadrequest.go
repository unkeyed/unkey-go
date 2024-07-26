// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// Code - A machine readable error code.
type Code string

const (
	CodeBadRequest Code = "BAD_REQUEST"
)

func (e Code) ToPointer() *Code {
	return &e
}
func (e *Code) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BAD_REQUEST":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

type Error struct {
	// A machine readable error code.
	Code Code `json:"code"`
	// A link to our documentation with more details about this error code
	Docs string `json:"docs"`
	// A human readable explanation of what went wrong
	Message string `json:"message"`
	// Please always include the requestId in your error report
	RequestID string `json:"requestId"`
}

func (o *Error) GetCode() Code {
	if o == nil {
		return Code("")
	}
	return o.Code
}

func (o *Error) GetDocs() string {
	if o == nil {
		return ""
	}
	return o.Docs
}

func (o *Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *Error) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

// ErrBadRequest - The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
type ErrBadRequest struct {
	Error_ Error `json:"error"`
}

var _ error = &ErrBadRequest{}

func (e *ErrBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
