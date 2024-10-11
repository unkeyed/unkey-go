# WhoamiResponseBody

The configuration for a single key


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ID`                                                        | *string*                                                    | :heavy_check_mark:                                          | The ID of the key                                           | key_123                                                     |
| `Name`                                                      | **string*                                                   | :heavy_minus_sign:                                          | The name of the key                                         | API Key 1                                                   |
| `Remaining`                                                 | **int64*                                                    | :heavy_minus_sign:                                          | The remaining number of requests for the key                | 1000                                                        |
| `Identity`                                                  | [*operations.Identity](../../models/operations/identity.md) | :heavy_minus_sign:                                          | The identity object associated with the key                 |                                                             |
| `Meta`                                                      | map[string]*any*                                            | :heavy_minus_sign:                                          | Metadata associated with the key                            | {<br/>"role": "admin",<br/>"plan": "premium"<br/>}          |
| `CreatedAt`                                                 | *int64*                                                     | :heavy_check_mark:                                          | The timestamp in milliseconds when the key was created      | 1620000000000                                               |
| `Enabled`                                                   | *bool*                                                      | :heavy_check_mark:                                          | Whether the key is enabled                                  | true                                                        |
| `Environment`                                               | **string*                                                   | :heavy_minus_sign:                                          | The environment the key is associated with                  | production                                                  |