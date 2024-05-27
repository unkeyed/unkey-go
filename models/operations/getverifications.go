// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unkeyed/unkey-go/internal/utils"
)

// Granularity - The granularity of the usage data to fetch, currently only `day` is supported
type Granularity string

const (
	GranularityDay Granularity = "day"
)

func (e Granularity) ToPointer() *Granularity {
	return &e
}
func (e *Granularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "day":
		*e = Granularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Granularity: %v", v)
	}
}

type GetVerificationsRequest struct {
	KeyID   *string `queryParam:"style=form,explode=true,name=keyId"`
	OwnerID *string `queryParam:"style=form,explode=true,name=ownerId"`
	Start   *int64  `queryParam:"style=form,explode=true,name=start"`
	End     *int64  `queryParam:"style=form,explode=true,name=end"`
	// The granularity of the usage data to fetch, currently only `day` is supported
	Granularity *Granularity `default:"day" queryParam:"style=form,explode=true,name=granularity"`
}

func (g GetVerificationsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetVerificationsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetVerificationsRequest) GetKeyID() *string {
	if o == nil {
		return nil
	}
	return o.KeyID
}

func (o *GetVerificationsRequest) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *GetVerificationsRequest) GetStart() *int64 {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *GetVerificationsRequest) GetEnd() *int64 {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *GetVerificationsRequest) GetGranularity() *Granularity {
	if o == nil {
		return nil
	}
	return o.Granularity
}

type Verifications struct {
	// The timestamp of the usage data
	Time int64 `json:"time"`
	// The number of successful requests
	Success float64 `json:"success"`
	// The number of requests that were rate limited
	RateLimited float64 `json:"rateLimited"`
	// The number of requests that exceeded the usage limit
	UsageExceeded float64 `json:"usageExceeded"`
}

func (o *Verifications) GetTime() int64 {
	if o == nil {
		return 0
	}
	return o.Time
}

func (o *Verifications) GetSuccess() float64 {
	if o == nil {
		return 0.0
	}
	return o.Success
}

func (o *Verifications) GetRateLimited() float64 {
	if o == nil {
		return 0.0
	}
	return o.RateLimited
}

func (o *Verifications) GetUsageExceeded() float64 {
	if o == nil {
		return 0.0
	}
	return o.UsageExceeded
}

// GetVerificationsResponseBody - Usage numbers over time
type GetVerificationsResponseBody struct {
	Verifications []Verifications `json:"verifications"`
}

func (o *GetVerificationsResponseBody) GetVerifications() []Verifications {
	if o == nil {
		return []Verifications{}
	}
	return o.Verifications
}
