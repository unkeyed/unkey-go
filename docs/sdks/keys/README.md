# Keys
(*Keys*)

### Available Operations

* [GetKey](#getkey)
* [DeleteKey](#deletekey)
* [CreateKey](#createkey)
* [VerifyKey](#verifykey)
* [UpdateKey](#updatekey)
* [UpdateRemaining](#updateremaining)
* [GetVerifications](#getverifications)
* [AddPermissions](#addpermissions)
* [RemovePermissions](#removepermissions)
* [SetPermissions](#setpermissions)
* [AddRoles](#addroles)
* [RemoveRoles](#removeroles)
* [SetRoles](#setroles)

## GetKey

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.GetKeyRequest{
        KeyID: "key_1234",
    }
    ctx := context.Background()
    res, err := s.Keys.GetKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Key != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [operations.GetKeyRequest](../../models/operations/getkeyrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |


### Response

**[*operations.GetKeyResponse](../../models/operations/getkeyresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## DeleteKey

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.DeleteKeyRequestBody{
        KeyID: "key_1234",
    }
    ctx := context.Background()
    res, err := s.Keys.DeleteKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteKeyRequestBody](../../models/operations/deletekeyrequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.DeleteKeyResponse](../../models/operations/deletekeyresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## CreateKey

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.CreateKeyRequestBody{
        APIID: "api_123",
        Name: unkeygo.String("my key"),
        OwnerID: unkeygo.String("team_123"),
        Meta: map[string]any{
            "billingTier": "PRO",
            "trialEnds": "2023-06-16T17:16:37.161Z",
        },
        Roles: []string{
            "admin",
            "finance",
        },
        Permissions: []string{
            "domains.create_record",
            "say_hello",
        },
        Expires: unkeygo.Int64(1623869797161),
        Remaining: unkeygo.Int64(1000),
        Refill: &operations.Refill{
            Interval: operations.IntervalDaily,
            Amount: 100,
        },
        Ratelimit: &operations.Ratelimit{
            Limit: 10,
            Duration: unkeygo.Int64(60000),
        },
        Enabled: unkeygo.Bool(false),
    }
    ctx := context.Background()
    res, err := s.Keys.CreateKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateKeyRequestBody](../../models/operations/createkeyrequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.CreateKeyResponse](../../models/operations/createkeyresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## VerifyKey

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/components"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := components.V1KeysVerifyKeyRequest{
        APIID: unkeygo.String("api_1234"),
        Key: "sk_1234",
        Ratelimits: []components.Ratelimits{
            components.Ratelimits{
                Name: "tokens",
            },
        },
    }
    ctx := context.Background()
    res, err := s.Keys.VerifyKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.V1KeysVerifyKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.V1KeysVerifyKeyRequest](../../models/components/v1keysverifykeyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.VerifyKeyResponse](../../models/operations/verifykeyresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## UpdateKey

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.UpdateKeyRequestBody{
        KeyID: "key_123",
        Name: unkeygo.String("Customer X"),
        OwnerID: unkeygo.String("user_123"),
        Meta: map[string]any{
            "roles": []any{
                "admin",
                "user",
            },
            "stripeCustomerId": "cus_1234",
        },
        Expires: unkeygo.Float64(0),
        Ratelimit: &operations.UpdateKeyRatelimit{
            Limit: 10,
            RefillRate: unkeygo.Int64(1),
            RefillInterval: unkeygo.Int64(60),
        },
        Remaining: unkeygo.Float64(1000),
        Refill: &operations.UpdateKeyRefill{
            Interval: operations.UpdateKeyIntervalDaily,
            Amount: 100,
        },
        Enabled: unkeygo.Bool(true),
        Roles: []operations.Roles{
            operations.Roles{
                ID: unkeygo.String("perm_123"),
            },
            operations.Roles{
                Name: unkeygo.String("dns.record.create"),
            },
            operations.Roles{
                Name: unkeygo.String("dns.record.delete"),
                Create: unkeygo.Bool(true),
            },
        },
        Permissions: []operations.Permissions{
            operations.Permissions{
                ID: unkeygo.String("perm_123"),
            },
            operations.Permissions{
                Name: unkeygo.String("dns.record.create"),
            },
            operations.Permissions{
                Name: unkeygo.String("dns.record.delete"),
                Create: unkeygo.Bool(true),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Keys.UpdateKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateKeyRequestBody](../../models/operations/updatekeyrequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.UpdateKeyResponse](../../models/operations/updatekeyresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## UpdateRemaining

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.UpdateRemainingRequestBody{
        KeyID: "key_123",
        Op: operations.OpSet,
        Value: unkeygo.Int64(1),
    }
    ctx := context.Background()
    res, err := s.Keys.UpdateRemaining(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateRemainingRequestBody](../../models/operations/updateremainingrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.UpdateRemainingResponse](../../models/operations/updateremainingresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## GetVerifications

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.GetVerificationsRequest{
        KeyID: unkeygo.String("key_1234"),
        OwnerID: unkeygo.String("chronark"),
        Start: unkeygo.Int64(1620000000000),
        End: unkeygo.Int64(1620000000000),
        Granularity: operations.GranularityDay.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Keys.GetVerifications(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetVerificationsRequest](../../models/operations/getverificationsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.GetVerificationsResponse](../../models/operations/getverificationsresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## AddPermissions

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.AddPermissionsRequestBody{
        KeyID: "<value>",
        Permissions: []operations.AddPermissionsPermissions{
            operations.AddPermissionsPermissions{},
        },
    }
    ctx := context.Background()
    res, err := s.Keys.AddPermissions(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AddPermissionsRequestBody](../../models/operations/addpermissionsrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.AddPermissionsResponse](../../models/operations/addpermissionsresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## RemovePermissions

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.RemovePermissionsRequestBody{
        KeyID: "<value>",
        Permissions: []operations.RemovePermissionsPermissions{
            operations.RemovePermissionsPermissions{
                ID: unkeygo.String("perm_123"),
            },
            operations.RemovePermissionsPermissions{
                Name: unkeygo.String("dns.record.create"),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Keys.RemovePermissions(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemovePermissionsRequestBody](../../models/operations/removepermissionsrequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.RemovePermissionsResponse](../../models/operations/removepermissionsresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## SetPermissions

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.SetPermissionsRequestBody{
        KeyID: "<value>",
        Permissions: []operations.SetPermissionsPermissions{
            operations.SetPermissionsPermissions{
                ID: unkeygo.String("perm_123"),
            },
            operations.SetPermissionsPermissions{
                Name: unkeygo.String("dns.record.create"),
            },
            operations.SetPermissionsPermissions{
                Name: unkeygo.String("dns.record.delete"),
                Create: unkeygo.Bool(true),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Keys.SetPermissions(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.SetPermissionsRequestBody](../../models/operations/setpermissionsrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.SetPermissionsResponse](../../models/operations/setpermissionsresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## AddRoles

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.AddRolesRequestBody{
        KeyID: "<value>",
        Roles: []operations.AddRolesRoles{
            operations.AddRolesRoles{
                ID: unkeygo.String("role_123"),
            },
            operations.AddRolesRoles{
                Name: unkeygo.String("dns.record.create"),
            },
            operations.AddRolesRoles{
                Name: unkeygo.String("dns.record.delete"),
                Create: unkeygo.Bool(true),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Keys.AddRoles(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.AddRolesRequestBody](../../models/operations/addrolesrequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.AddRolesResponse](../../models/operations/addrolesresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## RemoveRoles

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.RemoveRolesRequestBody{
        KeyID: "<value>",
        Roles: []operations.RemoveRolesRoles{
            operations.RemoveRolesRoles{
                ID: unkeygo.String("role_123"),
            },
            operations.RemoveRolesRoles{
                Name: unkeygo.String("dns.record.create"),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Keys.RemoveRoles(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveRolesRequestBody](../../models/operations/removerolesrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.RemoveRolesResponse](../../models/operations/removerolesresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## SetRoles

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.SetRolesRequestBody{
        KeyID: "<value>",
        Roles: []operations.SetRolesRoles{
            operations.SetRolesRoles{
                ID: unkeygo.String("role_123"),
            },
            operations.SetRolesRoles{
                Name: unkeygo.String("dns.record.create"),
            },
            operations.SetRolesRoles{
                Name: unkeygo.String("dns.record.delete"),
                Create: unkeygo.Bool(true),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Keys.SetRoles(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SetRolesRequestBody](../../models/operations/setrolesrequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.SetRolesResponse](../../models/operations/setrolesresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |
