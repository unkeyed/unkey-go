// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type WhoamiRequestBody struct {
	// The actual key to fetch
	Key string `json:"key"`
}

func (o *WhoamiRequestBody) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

// Identity - The identity object associated with the key
type Identity struct {
	// The identity ID associated with the key
	ID string `json:"id"`
	// The external identity ID associated with the key
	ExternalID string `json:"externalId"`
}

func (o *Identity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Identity) GetExternalID() string {
	if o == nil {
		return ""
	}
	return o.ExternalID
}

// WhoamiResponseBody - The configuration for a single key
type WhoamiResponseBody struct {
	// The ID of the key
	ID string `json:"id"`
	// The name of the key
	Name *string `json:"name,omitempty"`
	// The remaining number of requests for the key
	Remaining *int64 `json:"remaining,omitempty"`
	// The identity object associated with the key
	Identity *Identity `json:"identity,omitempty"`
	// Metadata associated with the key
	Meta map[string]any `json:"meta,omitempty"`
	// The timestamp in milliseconds when the key was created
	CreatedAt int64 `json:"createdAt"`
	// Whether the key is enabled
	Enabled bool `json:"enabled"`
	// The environment the key is associated with
	Environment *string `json:"environment,omitempty"`
}

func (o *WhoamiResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *WhoamiResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WhoamiResponseBody) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *WhoamiResponseBody) GetIdentity() *Identity {
	if o == nil {
		return nil
	}
	return o.Identity
}

func (o *WhoamiResponseBody) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *WhoamiResponseBody) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *WhoamiResponseBody) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *WhoamiResponseBody) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

type WhoamiResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The configuration for a single key
	Object *WhoamiResponseBody
}

func (o *WhoamiResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *WhoamiResponse) GetObject() *WhoamiResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
