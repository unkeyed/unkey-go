// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/unkeyed/unkey-go/internal/utils"
)

// Interval - Determines the rate at which verifications will be refilled. When 'daily' is set for 'interval' 'refillDay' will be set to null.
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

// Refill - Unkey allows you to refill remaining verifications on a key on a regular interval.
type Refill struct {
	// Determines the rate at which verifications will be refilled. When 'daily' is set for 'interval' 'refillDay' will be set to null.
	Interval Interval `json:"interval"`
	// Resets `remaining` to this value every interval.
	Amount int64 `json:"amount"`
	// The day verifications will refill each month, when interval is set to 'monthly'. Value is not zero-indexed making 1 the first day of the month. If left blank it will default to the first day of the month. When 'daily' is set for 'interval' 'refillDay' will be set to null.
	RefillDay *float64 `default:"1" json:"refillDay"`
	// The unix timestamp in miliseconds when the key was last refilled.
	LastRefillAt *int64 `json:"lastRefillAt,omitempty"`
}

func (r Refill) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Refill) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
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

func (o *Refill) GetRefillDay() *float64 {
	if o == nil {
		return nil
	}
	return o.RefillDay
}

func (o *Refill) GetLastRefillAt() *int64 {
	if o == nil {
		return nil
	}
	return o.LastRefillAt
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
	Async bool `json:"async"`
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
	Type *Type `json:"type,omitempty"`
	// The total amount of burstable requests.
	Limit int64 `json:"limit"`
	// How many tokens to refill during each refillInterval.
	RefillRate *int64 `json:"refillRate,omitempty"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval *int64 `json:"refillInterval,omitempty"`
	// The duration of the ratelimit window, in milliseconds.
	Duration int64 `json:"duration"`
}

func (o *Ratelimit) GetAsync() bool {
	if o == nil {
		return false
	}
	return o.Async
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

func (o *Ratelimit) GetRefillRate() *int64 {
	if o == nil {
		return nil
	}
	return o.RefillRate
}

func (o *Ratelimit) GetRefillInterval() *int64 {
	if o == nil {
		return nil
	}
	return o.RefillInterval
}

func (o *Ratelimit) GetDuration() int64 {
	if o == nil {
		return 0
	}
	return o.Duration
}

// Identity - The identity of the key
type Identity struct {
	// The id of the identity
	ID string `json:"id"`
	// The external id of the identity
	ExternalID string `json:"externalId"`
	// Any additional metadata attached to the identity
	Meta map[string]any `json:"meta,omitempty"`
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

func (o *Identity) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

type Key struct {
	// The id of the key
	ID string `json:"id"`
	// The first few characters of the key to visually identify it
	Start string `json:"start"`
	// The id of the workspace that owns the key
	WorkspaceID string `json:"workspaceId"`
	// The id of the api that this key is for
	APIID *string `json:"apiId,omitempty"`
	// The name of the key, give keys a name to easily identify their purpose
	Name *string `json:"name,omitempty"`
	// The id of the tenant associated with this key. Use whatever reference you have in your system to identify the tenant. When verifying the key, we will send this field back to you, so you know who is accessing your API.
	OwnerID *string `json:"ownerId,omitempty"`
	// Any additional metadata you want to store with the key
	Meta map[string]any `json:"meta,omitempty"`
	// The unix timestamp in milliseconds when the key was created
	CreatedAt int64 `json:"createdAt"`
	// The unix timestamp in milliseconds when the key was last updated
	UpdatedAt *int64 `json:"updatedAt,omitempty"`
	// The unix timestamp in milliseconds when the key will expire. If this field is null or undefined, the key is not expiring.
	Expires *int64 `json:"expires,omitempty"`
	// The number of requests that can be made with this key before it becomes invalid. If this field is null or undefined, the key has no request limit.
	Remaining *int64 `json:"remaining,omitempty"`
	// Unkey allows you to refill remaining verifications on a key on a regular interval.
	Refill *Refill `json:"refill,omitempty"`
	// Unkey comes with per-key ratelimiting out of the box.
	Ratelimit *Ratelimit `json:"ratelimit,omitempty"`
	// All roles this key belongs to
	Roles []string `json:"roles,omitempty"`
	// All permissions this key has
	Permissions []string `json:"permissions,omitempty"`
	// Sets if key is enabled or disabled. Disabled keys are not valid.
	Enabled *bool `json:"enabled,omitempty"`
	// The key in plaintext
	Plaintext *string `json:"plaintext,omitempty"`
	// The identity of the key
	Identity *Identity `json:"identity,omitempty"`
}

func (o *Key) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Key) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}

func (o *Key) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *Key) GetAPIID() *string {
	if o == nil {
		return nil
	}
	return o.APIID
}

func (o *Key) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Key) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *Key) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *Key) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *Key) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Key) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *Key) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *Key) GetRefill() *Refill {
	if o == nil {
		return nil
	}
	return o.Refill
}

func (o *Key) GetRatelimit() *Ratelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

func (o *Key) GetRoles() []string {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *Key) GetPermissions() []string {
	if o == nil {
		return nil
	}
	return o.Permissions
}

func (o *Key) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Key) GetPlaintext() *string {
	if o == nil {
		return nil
	}
	return o.Plaintext
}

func (o *Key) GetIdentity() *Identity {
	if o == nil {
		return nil
	}
	return o.Identity
}
