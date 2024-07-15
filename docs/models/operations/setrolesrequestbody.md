# SetRolesRequestBody


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `KeyID`                                                                     | *string*                                                                    | :heavy_check_mark:                                                          | The id of the key.                                                          |
| `Roles`                                                                     | [][operations.SetRolesRoles](../../models/operations/setrolesroles.md)      | :heavy_check_mark:                                                          | The roles you want to set for this key. This overwrites all existing roles. |