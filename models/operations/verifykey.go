// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type VerifyKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The verification result
	V1KeysVerifyKeyResponse *components.V1KeysVerifyKeyResponse
}

func (o *VerifyKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VerifyKeyResponse) GetV1KeysVerifyKeyResponse() *components.V1KeysVerifyKeyResponse {
	if o == nil {
		return nil
	}
	return o.V1KeysVerifyKeyResponse
}
