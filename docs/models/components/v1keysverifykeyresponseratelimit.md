# V1KeysVerifyKeyResponseRatelimit

Multi ratelimits TODO:


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Limit`                                                      | *int64*                                                      | :heavy_check_mark:                                           | Maximum number of requests that can be made inside a window  | 10                                                           |
| `Remaining`                                                  | *int64*                                                      | :heavy_check_mark:                                           | Remaining requests after this verification                   | 9                                                            |
| `Reset`                                                      | *int64*                                                      | :heavy_check_mark:                                           | Unix timestamp in milliseconds when the ratelimit will reset | 3600000                                                      |