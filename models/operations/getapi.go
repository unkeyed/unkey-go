// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type GetAPIRequest struct {
	APIID string `queryParam:"style=form,explode=true,name=apiId"`
}

func (o *GetAPIRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

// GetAPIResponseBody - The configuration for an api
type GetAPIResponseBody struct {
	// The id of the key
	ID string `json:"id"`
	// The id of the workspace that owns the api
	WorkspaceID string `json:"workspaceId"`
	// The name of the api. This is internal and your users will not see this.
	Name *string `json:"name,omitempty"`
}

func (o *GetAPIResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetAPIResponseBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *GetAPIResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
