// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/internal/utils"
	"github.com/unkeyed/unkey-go/models/components"
)

type SetOverrideRequestBody struct {
	// The id of the namespace. Either namespaceId or namespaceName must be provided
	NamespaceID *string `json:"namespaceId,omitempty"`
	// Namespaces group different limits together for better analytics. You might have a namespace for your public API and one for internal tRPC routes. Wildcards can also be used, more info can be found at https://www.unkey.com/docs/ratelimiting/overrides#wildcard-rules
	NamespaceName *string `json:"namespaceName,omitempty"`
	// Identifier of your user, this can be their userId, an email, an ip or anything else. Wildcards ( * ) can be used to match multiple identifiers, More info can be found at https://www.unkey.com/docs/ratelimiting/overrides#wildcard-rules
	Identifier string `json:"identifier"`
	// How many requests may pass in a given window.
	Limit int64 `json:"limit"`
	// The window duration in milliseconds
	Duration int64 `json:"duration"`
	// Async will return a response immediately, lowering latency at the cost of accuracy.
	Async *bool `default:"false" json:"async"`
}

func (s SetOverrideRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SetOverrideRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SetOverrideRequestBody) GetNamespaceID() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceID
}

func (o *SetOverrideRequestBody) GetNamespaceName() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceName
}

func (o *SetOverrideRequestBody) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *SetOverrideRequestBody) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *SetOverrideRequestBody) GetDuration() int64 {
	if o == nil {
		return 0
	}
	return o.Duration
}

func (o *SetOverrideRequestBody) GetAsync() *bool {
	if o == nil {
		return nil
	}
	return o.Async
}

// SetOverrideResponseBody - Sucessfully created a ratelimit override
type SetOverrideResponseBody struct {
	// The id of the override. This is used internally
	OverrideID string `json:"overrideId"`
}

func (o *SetOverrideResponseBody) GetOverrideID() string {
	if o == nil {
		return ""
	}
	return o.OverrideID
}

type SetOverrideResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Sucessfully created a ratelimit override
	Object *SetOverrideResponseBody
}

func (o *SetOverrideResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetOverrideResponse) GetObject() *SetOverrideResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
