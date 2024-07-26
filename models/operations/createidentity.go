// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type Ratelimits struct {
	// The name of this limit. You will need to use this again when verifying a key.
	Name string `json:"name"`
	// How many requests may pass within a given window before requests are rejected.
	Limit int64 `json:"limit"`
	// The duration for each ratelimit window in milliseconds.
	Duration int64 `json:"duration"`
}

func (o *Ratelimits) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Ratelimits) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *Ratelimits) GetDuration() int64 {
	if o == nil {
		return 0
	}
	return o.Duration
}

type CreateIdentityRequestBody struct {
	// The id of this identity in your system.
	//
	// This usually comes from your authentication provider and could be a userId, organisationId or even an email.
	// It does not matter what you use, as long as it uniquely identifies something in your application.
	//
	ExternalID string `json:"externalId"`
	// Attach metadata to this identity that you need to have access to when verifying a key.
	//
	// This will be returned as part of the `verifyKey` response.
	//
	Meta map[string]any `json:"meta,omitempty"`
	// Attach ratelimits to this identity.
	//
	// When verifying keys, you can specify which limits you want to use and all keys attached to this identity, will share the limits.
	Ratelimits []Ratelimits `json:"ratelimits,omitempty"`
}

func (o *CreateIdentityRequestBody) GetExternalID() string {
	if o == nil {
		return ""
	}
	return o.ExternalID
}

func (o *CreateIdentityRequestBody) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *CreateIdentityRequestBody) GetRatelimits() []Ratelimits {
	if o == nil {
		return nil
	}
	return o.Ratelimits
}

// CreateIdentityResponseBody - The configuration for an api
type CreateIdentityResponseBody struct {
	// The id of the identity. Used internally, you do not need to store this.
	IdentityID string `json:"identityId"`
}

func (o *CreateIdentityResponseBody) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

type CreateIdentityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The configuration for an api
	Object *CreateIdentityResponseBody
}

func (o *CreateIdentityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIdentityResponse) GetObject() *CreateIdentityResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
