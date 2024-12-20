# Ratelimit
(*Ratelimit*)

## Overview

### Available Operations

* [SetOverride](#setoverride)
* [ListOverrides](#listoverrides)
* [GetOverride](#getoverride)

## SetOverride

### Example Usage

```go
package main

import(
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := unkeygo.New(
        unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
    )

    res, err := s.Ratelimit.SetOverride(ctx, operations.SetOverrideRequestBody{
        NamespaceID: unkeygo.String("rlns_1234"),
        NamespaceName: unkeygo.String("email.outbound"),
        Identifier: "user_123",
        Limit: 10,
        Duration: 60000,
    })
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
| `request`                                                                              | [operations.SetOverrideRequestBody](../../models/operations/setoverriderequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.SetOverrideResponse](../../models/operations/setoverrideresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4XX, 5XX                         | \*/\*                            |

## ListOverrides

### Example Usage

```go
package main

import(
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := unkeygo.New(
        unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
    )

    res, err := s.Ratelimit.ListOverrides(ctx, operations.ListOverridesRequest{
        NamespaceID: unkeygo.String("rlns_1234"),
        NamespaceName: unkeygo.String("email.outbound"),
        Limit: unkeygo.Int64(100),
    })
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
| `request`                                                                          | [operations.ListOverridesRequest](../../models/operations/listoverridesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListOverridesResponse](../../models/operations/listoverridesresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4XX, 5XX                         | \*/\*                            |

## GetOverride

### Example Usage

```go
package main

import(
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := unkeygo.New(
        unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
    )

    res, err := s.Ratelimit.GetOverride(ctx, operations.GetOverrideRequest{
        NamespaceID: unkeygo.String("rlns_1234"),
        NamespaceName: unkeygo.String("email.outbound"),
        Identifier: "user_123",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetOverrideRequest](../../models/operations/getoverriderequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetOverrideResponse](../../models/operations/getoverrideresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4XX, 5XX                         | \*/\*                            |