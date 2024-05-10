// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unkeyed/sdk-go/internal/utils"
)

// Interval - Unkey will automatically refill verifications at the set interval.
type Interval string

const (
	IntervalDaily   Interval = "daily"
	IntervalMonthly Interval = "monthly"
)

func (e Interval) ToPointer() *Interval {
	return &e
}

func (e *Interval) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "daily":
		fallthrough
	case "monthly":
		*e = Interval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Interval: %v", v)
	}
}

// Refill - Unkey enables you to refill verifications for each key at regular intervals.
type Refill struct {
	// Unkey will automatically refill verifications at the set interval.
	Interval Interval `json:"interval"`
	// The number of verifications to refill for each occurrence is determined individually for each key.
	Amount int64 `json:"amount"`
}

func (o *Refill) GetInterval() Interval {
	if o == nil {
		return Interval("")
	}
	return o.Interval
}

func (o *Refill) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

// Type - Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
//
// https://unkey.dev/docs/features/ratelimiting - Learn more
type Type string

const (
	TypeFast       Type = "fast"
	TypeConsistent Type = "consistent"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fast":
		fallthrough
	case "consistent":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// Ratelimit - Unkey comes with per-key ratelimiting out of the box.
type Ratelimit struct {
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
	Type *Type `default:"fast" json:"type"`
	// The total amount of burstable requests.
	Limit int64 `json:"limit"`
	// How many tokens to refill during each refillInterval.
	RefillRate int64 `json:"refillRate"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval int64 `json:"refillInterval"`
}

func (r Ratelimit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Ratelimit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Ratelimit) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Ratelimit) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *Ratelimit) GetRefillRate() int64 {
	if o == nil {
		return 0
	}
	return o.RefillRate
}

func (o *Ratelimit) GetRefillInterval() int64 {
	if o == nil {
		return 0
	}
	return o.RefillInterval
}

type V1KeysCreateKeyRequestBody struct {
	// Choose an `API` where this key should be created.
	APIID string `json:"apiId"`
	// To make it easier for your users to understand which product an api key belongs to, you can add prefix them.
	//
	// For example Stripe famously prefixes their customer ids with cus_ or their api keys with sk_live_.
	//
	// The underscore is automatically added if you are defining a prefix, for example: "prefix": "abc" will result in a key like abc_xxxxxxxxx
	//
	Prefix *string `json:"prefix,omitempty"`
	// The name for your Key. This is not customer facing.
	Name *string `json:"name,omitempty"`
	// The byte length used to generate your key determines its entropy as well as its length. Higher is better, but keys become longer and more annoying to handle. The default is 16 bytes, or 2^^128 possible combinations.
	ByteLength *int64 `default:"16" json:"byteLength"`
	// Your user’s Id. This will provide a link between Unkey and your customer record.
	// When validating a key, we will return this back to you, so you can clearly identify your user from their api key.
	OwnerID *string `json:"ownerId,omitempty"`
	// This is a place for dynamic meta data, anything that feels useful for you should go here
	Meta map[string]interface{} `json:"meta,omitempty"`
	// A list of roles that this key should have. If the role does not exist, an error is thrown
	Roles []string `json:"roles,omitempty"`
	// You can auto expire keys by providing a unix timestamp in milliseconds. Once Keys expire they will automatically be disabled and are no longer valid unless you enable them again.
	Expires *int64 `json:"expires,omitempty"`
	// You can limit the number of requests a key can make. Once a key reaches 0 remaining requests, it will automatically be disabled and is no longer valid unless you update it.
	Remaining *int64 `json:"remaining,omitempty"`
	// Unkey enables you to refill verifications for each key at regular intervals.
	Refill *Refill `json:"refill,omitempty"`
	// Unkey comes with per-key ratelimiting out of the box.
	Ratelimit *Ratelimit `json:"ratelimit,omitempty"`
	// Sets if key is enabled or disabled. Disabled keys are not valid.
	Enabled *bool `default:"true" json:"enabled"`
	// Environments allow you to divide your keyspace.
	//
	// Some applications like Stripe, Clerk, WorkOS and others have a concept of "live" and "test" keys to
	// give the developer a way to develop their own application without the risk of modifying real world
	// resources.
	//
	// When you set an environment, we will return it back to you when validating the key, so you can
	// handle it correctly.
	//
	Environment *string `json:"environment,omitempty"`
}

func (v V1KeysCreateKeyRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V1KeysCreateKeyRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V1KeysCreateKeyRequestBody) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *V1KeysCreateKeyRequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *V1KeysCreateKeyRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *V1KeysCreateKeyRequestBody) GetByteLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ByteLength
}

func (o *V1KeysCreateKeyRequestBody) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *V1KeysCreateKeyRequestBody) GetMeta() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *V1KeysCreateKeyRequestBody) GetRoles() []string {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *V1KeysCreateKeyRequestBody) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *V1KeysCreateKeyRequestBody) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *V1KeysCreateKeyRequestBody) GetRefill() *Refill {
	if o == nil {
		return nil
	}
	return o.Refill
}

func (o *V1KeysCreateKeyRequestBody) GetRatelimit() *Ratelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

func (o *V1KeysCreateKeyRequestBody) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *V1KeysCreateKeyRequestBody) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

// V1KeysCreateKeyResponseBody - The configuration for an api
type V1KeysCreateKeyResponseBody struct {
	// The id of the key. This is not a secret and can be stored as a reference if you wish. You need the keyId to update or delete a key later.
	KeyID string `json:"keyId"`
	// The newly created api key, do not store this on your own system but pass it along to your user.
	Key string `json:"key"`
}

func (o *V1KeysCreateKeyResponseBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *V1KeysCreateKeyResponseBody) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}
