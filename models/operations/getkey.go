// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

type GetKeyRequest struct {
	KeyID   string `queryParam:"style=form,explode=true,name=keyId"`
	Decrypt *bool  `queryParam:"style=form,explode=true,name=decrypt"`
}

func (o *GetKeyRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *GetKeyRequest) GetDecrypt() *bool {
	if o == nil {
		return nil
	}
	return o.Decrypt
}
