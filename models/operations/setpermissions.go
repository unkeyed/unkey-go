// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type SetPermissionsPermissions struct {
	// The id of the permission. Provide either `id` or `name`. If both are provided `id` is used.
	ID *string `json:"id,omitempty"`
	// Identify the permission via its name. Provide either `id` or `name`. If both are provided `id` is used.
	Name *string `json:"name,omitempty"`
	// Set to true to automatically create the permissions they do not exist yet. Only works when specifying `name`.
	//                 Autocreating permissions requires your root key to have the `rbac.*.create_permission` permission, otherwise the request will get rejected
	Create *bool `json:"create,omitempty"`
}

func (o *SetPermissionsPermissions) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetPermissionsPermissions) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SetPermissionsPermissions) GetCreate() *bool {
	if o == nil {
		return nil
	}
	return o.Create
}

type SetPermissionsRequestBody struct {
	// The id of the key.
	KeyID string `json:"keyId"`
	// The permissions you want to set for this key. This overwrites all existing permissions.
	//             Setting permissions requires the `rbac.*.add_permission_to_key` permission.
	Permissions []SetPermissionsPermissions `json:"permissions"`
}

func (o *SetPermissionsRequestBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *SetPermissionsRequestBody) GetPermissions() []SetPermissionsPermissions {
	if o == nil {
		return []SetPermissionsPermissions{}
	}
	return o.Permissions
}

type SetPermissionsResponseBody struct {
	// The id of the permission. This is used internally
	ID string `json:"id"`
	// The name of the permission
	Name string `json:"name"`
}

func (o *SetPermissionsResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SetPermissionsResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type SetPermissionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// All currently connected permissions
	ResponseBodies []SetPermissionsResponseBody
}

func (o *SetPermissionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetPermissionsResponse) GetResponseBodies() []SetPermissionsResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
