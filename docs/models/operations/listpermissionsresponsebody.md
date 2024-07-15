# ListPermissionsResponseBody


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                    | *string*                                                                                                | :heavy_check_mark:                                                                                      | The id of the permission                                                                                | perm_123                                                                                                |
| `Name`                                                                                                  | *string*                                                                                                | :heavy_check_mark:                                                                                      | The name of the permission.                                                                             | domain.record.manager                                                                                   |
| `Description`                                                                                           | **string*                                                                                               | :heavy_minus_sign:                                                                                      | The description of what this permission does. This is just for your team, your users will not see this. | Can manage dns records                                                                                  |