// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type GetKeyRequest struct {
	KeyID   string `queryParam:"style=form,explode=true,name=keyId"`
	Decrypt *bool  `queryParam:"style=form,explode=true,name=decrypt"`
}

func (o *GetKeyRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *GetKeyRequest) GetDecrypt() *bool {
	if o == nil {
		return nil
	}
	return o.Decrypt
}

type GetKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The configuration for a single key
	Key *components.Key
}

func (o *GetKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetKeyResponse) GetKey() *components.Key {
	if o == nil {
		return nil
	}
	return o.Key
}
