// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// V1KeysVerifyKeyResponseRatelimit - The ratelimit configuration for this key. If this field is null or undefined, the key has no ratelimit.
type V1KeysVerifyKeyResponseRatelimit struct {
	// Maximum number of requests that can be made inside a window
	Limit int64 `json:"limit"`
	// Remaining requests after this verification
	Remaining int64 `json:"remaining"`
	// Unix timestamp in milliseconds when the ratelimit will reset
	Reset int64 `json:"reset"`
}

func (o *V1KeysVerifyKeyResponseRatelimit) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *V1KeysVerifyKeyResponseRatelimit) GetRemaining() int64 {
	if o == nil {
		return 0
	}
	return o.Remaining
}

func (o *V1KeysVerifyKeyResponseRatelimit) GetReset() int64 {
	if o == nil {
		return 0
	}
	return o.Reset
}

// Code - A machine readable code why the key is not valid.
// Possible values are:
// - VALID: the key is valid and you should proceed
// - NOT_FOUND: the key does not exist or has expired
// - FORBIDDEN: the key is not allowed to access the api
// - USAGE_EXCEEDED: the key has exceeded its request limit
// - RATE_LIMITED: the key has been ratelimited
// - UNAUTHORIZED: the key is not authorized
// - DISABLED: the key is disabled
// - INSUFFICIENT_PERMISSIONS: you do not have the required permissions to perform this action
type Code string

const (
	CodeValid                   Code = "VALID"
	CodeNotFound                Code = "NOT_FOUND"
	CodeForbidden               Code = "FORBIDDEN"
	CodeUsageExceeded           Code = "USAGE_EXCEEDED"
	CodeRateLimited             Code = "RATE_LIMITED"
	CodeUnauthorized            Code = "UNAUTHORIZED"
	CodeDisabled                Code = "DISABLED"
	CodeInsufficientPermissions Code = "INSUFFICIENT_PERMISSIONS"
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
	case "VALID":
		fallthrough
	case "NOT_FOUND":
		fallthrough
	case "FORBIDDEN":
		fallthrough
	case "USAGE_EXCEEDED":
		fallthrough
	case "RATE_LIMITED":
		fallthrough
	case "UNAUTHORIZED":
		fallthrough
	case "DISABLED":
		fallthrough
	case "INSUFFICIENT_PERMISSIONS":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

type V1KeysVerifyKeyResponse struct {
	// The id of the key
	KeyID *string `json:"keyId,omitempty"`
	// Whether the key is valid or not.
	// A key could be invalid for a number of reasons, for example if it has expired, has no more verifications left or if it has been deleted.
	Valid bool `json:"valid"`
	// The name of the key, give keys a name to easily identifiy their purpose
	Name *string `json:"name,omitempty"`
	// The id of the tenant associated with this key. Use whatever reference you have in your system to identify the tenant. When verifying the key, we will send this field back to you, so you know who is accessing your API.
	OwnerID *string `json:"ownerId,omitempty"`
	// Any additional metadata you want to store with the key
	Meta map[string]any `json:"meta,omitempty"`
	// The unix timestamp in milliseconds when the key will expire. If this field is null or undefined, the key is not expiring.
	Expires *int64 `json:"expires,omitempty"`
	// The ratelimit configuration for this key. If this field is null or undefined, the key has no ratelimit.
	Ratelimit *V1KeysVerifyKeyResponseRatelimit `json:"ratelimit,omitempty"`
	// The number of requests that can be made with this key before it becomes invalid. If this field is null or undefined, the key has no request limit.
	Remaining *int64 `json:"remaining,omitempty"`
	// A machine readable code why the key is not valid.
	// Possible values are:
	// - VALID: the key is valid and you should proceed
	// - NOT_FOUND: the key does not exist or has expired
	// - FORBIDDEN: the key is not allowed to access the api
	// - USAGE_EXCEEDED: the key has exceeded its request limit
	// - RATE_LIMITED: the key has been ratelimited
	// - UNAUTHORIZED: the key is not authorized
	// - DISABLED: the key is disabled
	// - INSUFFICIENT_PERMISSIONS: you do not have the required permissions to perform this action
	//
	Code Code `json:"code"`
	// Sets the key to be enabled or disabled. Disabled keys will not verify.
	Enabled *bool `json:"enabled,omitempty"`
	// A list of all the permissions this key is connected to.
	Permissions []string `json:"permissions,omitempty"`
	// The environment of the key, this is what what you set when you crated the key
	Environment *string `json:"environment,omitempty"`
}

func (o *V1KeysVerifyKeyResponse) GetKeyID() *string {
	if o == nil {
		return nil
	}
	return o.KeyID
}

func (o *V1KeysVerifyKeyResponse) GetValid() bool {
	if o == nil {
		return false
	}
	return o.Valid
}

func (o *V1KeysVerifyKeyResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *V1KeysVerifyKeyResponse) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *V1KeysVerifyKeyResponse) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *V1KeysVerifyKeyResponse) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *V1KeysVerifyKeyResponse) GetRatelimit() *V1KeysVerifyKeyResponseRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

func (o *V1KeysVerifyKeyResponse) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *V1KeysVerifyKeyResponse) GetCode() Code {
	if o == nil {
		return Code("")
	}
	return o.Code
}

func (o *V1KeysVerifyKeyResponse) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *V1KeysVerifyKeyResponse) GetPermissions() []string {
	if o == nil {
		return nil
	}
	return o.Permissions
}

func (o *V1KeysVerifyKeyResponse) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}
