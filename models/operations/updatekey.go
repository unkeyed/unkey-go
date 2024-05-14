// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
)

// UpdateKeyType - Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
//
// https://unkey.dev/docs/features/ratelimiting - Learn more
type UpdateKeyType string

const (
	UpdateKeyTypeFast       UpdateKeyType = "fast"
	UpdateKeyTypeConsistent UpdateKeyType = "consistent"
)

func (e UpdateKeyType) ToPointer() *UpdateKeyType {
	return &e
}

func (e *UpdateKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fast":
		fallthrough
	case "consistent":
		*e = UpdateKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateKeyType: %v", v)
	}
}

// UpdateKeyRatelimit - Unkey comes with per-key ratelimiting out of the box. Set `null` to disable.
type UpdateKeyRatelimit struct {
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
	Type UpdateKeyType `json:"type"`
	// The total amount of burstable requests.
	Limit int64 `json:"limit"`
	// How many tokens to refill during each refillInterval.
	RefillRate int64 `json:"refillRate"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval int64 `json:"refillInterval"`
}

func (o *UpdateKeyRatelimit) GetType() UpdateKeyType {
	if o == nil {
		return UpdateKeyType("")
	}
	return o.Type
}

func (o *UpdateKeyRatelimit) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *UpdateKeyRatelimit) GetRefillRate() int64 {
	if o == nil {
		return 0
	}
	return o.RefillRate
}

func (o *UpdateKeyRatelimit) GetRefillInterval() int64 {
	if o == nil {
		return 0
	}
	return o.RefillInterval
}

// UpdateKeyInterval - Unkey will automatically refill verifications at the set interval. If null is used the refill functionality will be removed from the key.
type UpdateKeyInterval string

const (
	UpdateKeyIntervalDaily   UpdateKeyInterval = "daily"
	UpdateKeyIntervalMonthly UpdateKeyInterval = "monthly"
)

func (e UpdateKeyInterval) ToPointer() *UpdateKeyInterval {
	return &e
}

func (e *UpdateKeyInterval) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "daily":
		fallthrough
	case "monthly":
		*e = UpdateKeyInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateKeyInterval: %v", v)
	}
}

// UpdateKeyRefill - Unkey enables you to refill verifications for each key at regular intervals.
type UpdateKeyRefill struct {
	// Unkey will automatically refill verifications at the set interval. If null is used the refill functionality will be removed from the key.
	Interval UpdateKeyInterval `json:"interval"`
	// The amount of verifications to refill for each occurrence is determined individually for each key.
	Amount int64 `json:"amount"`
}

func (o *UpdateKeyRefill) GetInterval() UpdateKeyInterval {
	if o == nil {
		return UpdateKeyInterval("")
	}
	return o.Interval
}

func (o *UpdateKeyRefill) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

type UpdateKeyRequestBody struct {
	// The id of the key you want to modify
	KeyID string `json:"keyId"`
	// The name of the key
	Name *string `json:"name,omitempty"`
	// The id of the tenant associated with this key. Use whatever reference you have in your system to identify the tenant. When verifying the key, we will send this field back to you, so you know who is accessing your API.
	OwnerID *string `json:"ownerId,omitempty"`
	// Any additional metadata you want to store with the key
	Meta map[string]any `json:"meta,omitempty"`
	// The unix timestamp in milliseconds when the key will expire. If this field is null or undefined, the key is not expiring.
	Expires *float64 `json:"expires,omitempty"`
	// Unkey comes with per-key ratelimiting out of the box. Set `null` to disable.
	Ratelimit *UpdateKeyRatelimit `json:"ratelimit,omitempty"`
	// The number of requests that can be made with this key before it becomes invalid. Set `null` to disable.
	Remaining *float64 `json:"remaining,omitempty"`
	// Unkey enables you to refill verifications for each key at regular intervals.
	Refill *UpdateKeyRefill `json:"refill,omitempty"`
	// Set if key is enabled or disabled. If disabled, the key cannot be used to verify.
	Enabled *bool `json:"enabled,omitempty"`
}

func (o *UpdateKeyRequestBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *UpdateKeyRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateKeyRequestBody) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *UpdateKeyRequestBody) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *UpdateKeyRequestBody) GetExpires() *float64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *UpdateKeyRequestBody) GetRatelimit() *UpdateKeyRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

func (o *UpdateKeyRequestBody) GetRemaining() *float64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *UpdateKeyRequestBody) GetRefill() *UpdateKeyRefill {
	if o == nil {
		return nil
	}
	return o.Refill
}

func (o *UpdateKeyRequestBody) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

// UpdateKeyResponseBody - The key was successfully updated, it may take up to 30s for this to take effect in all regions
type UpdateKeyResponseBody struct {
}
