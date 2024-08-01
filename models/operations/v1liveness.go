// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type Services struct {
	// The name of the connected metrics service
	Metrics string `json:"metrics"`
	// The name of the connected logger service
	Logger string `json:"logger"`
	// The name of the connected ratelimit service
	Ratelimit string `json:"ratelimit"`
	// The name of the connected usagelimit service
	Usagelimit string `json:"usagelimit"`
	// The name of the connected analytics service
	Analytics string `json:"analytics"`
}

func (o *Services) GetMetrics() string {
	if o == nil {
		return ""
	}
	return o.Metrics
}

func (o *Services) GetLogger() string {
	if o == nil {
		return ""
	}
	return o.Logger
}

func (o *Services) GetRatelimit() string {
	if o == nil {
		return ""
	}
	return o.Ratelimit
}

func (o *Services) GetUsagelimit() string {
	if o == nil {
		return ""
	}
	return o.Usagelimit
}

func (o *Services) GetAnalytics() string {
	if o == nil {
		return ""
	}
	return o.Analytics
}

// V1LivenessResponseBody - The configured services and their status
type V1LivenessResponseBody struct {
	// The status of the server
	Status   string   `json:"status"`
	Services Services `json:"services"`
}

func (o *V1LivenessResponseBody) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *V1LivenessResponseBody) GetServices() Services {
	if o == nil {
		return Services{}
	}
	return o.Services
}

type V1LivenessResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The configured services and their status
	Object *V1LivenessResponseBody
}

func (o *V1LivenessResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V1LivenessResponse) GetObject() *V1LivenessResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
