# Keys
(*Keys*)

### Available Operations

* [V1KeysGetKey](#v1keysgetkey)
* [V1KeysDeleteKey](#v1keysdeletekey)
* [V1KeysCreateKey](#v1keyscreatekey)
* [V1KeysVerifyKey](#v1keysverifykey)
* [V1KeysUpdateKey](#v1keysupdatekey)
* [V1KeysUpdateRemaining](#v1keysupdateremaining)
* [V1KeysGetVerifications](#v1keysgetverifications)

## V1KeysGetKey

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-sdk-go/models/components"
	unkeysdkgo "github.com/unkeyed/unkey-sdk-go"
	"github.com/unkeyed/unkey-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeysdkgo.New(
        unkeysdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.V1KeysGetKeyRequest{
        KeyID: "key_1234",
    }
    
    ctx := context.Background()
    res, err := s.Keys.V1KeysGetKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.V1KeysGetKeyRequest](../../models/operations/v1keysgetkeyrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*components.Key](../../models/components/key.md), error**
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

## V1KeysDeleteKey

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-sdk-go/models/components"
	unkeysdkgo "github.com/unkeyed/unkey-sdk-go"
	"github.com/unkeyed/unkey-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeysdkgo.New(
        unkeysdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.V1KeysDeleteKeyRequestBody{
        KeyID: "key_1234",
    }
    
    ctx := context.Background()
    res, err := s.Keys.V1KeysDeleteKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V1KeysDeleteKeyRequestBody](../../models/operations/v1keysdeletekeyrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.V1KeysDeleteKeyResponseBody](../../models/operations/v1keysdeletekeyresponsebody.md), error**
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

## V1KeysCreateKey

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-sdk-go/models/components"
	unkeysdkgo "github.com/unkeyed/unkey-sdk-go"
	"github.com/unkeyed/unkey-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeysdkgo.New(
        unkeysdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.V1KeysCreateKeyRequestBody{
        APIID: "api_123",
        Name: unkeysdkgo.String("my key"),
        OwnerID: unkeysdkgo.String("team_123"),
        Meta: map[string]interface{}{
            "billingTier": "PRO",
            "trialEnds": "2023-06-16T17:16:37.161Z",
        },
        Roles: []string{
            "admin",
            "finance",
        },
        Expires: unkeysdkgo.Int64(1623869797161),
        Remaining: unkeysdkgo.Int64(1000),
        Refill: &operations.Refill{
            Interval: operations.IntervalDaily,
            Amount: 100,
        },
        Ratelimit: &operations.Ratelimit{
            Type: operations.TypeFast.ToPointer(),
            Limit: 10,
            RefillRate: 1,
            RefillInterval: 60,
        },
        Enabled: unkeysdkgo.Bool(false),
    }
    
    ctx := context.Background()
    res, err := s.Keys.V1KeysCreateKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V1KeysCreateKeyRequestBody](../../models/operations/v1keyscreatekeyrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.V1KeysCreateKeyResponseBody](../../models/operations/v1keyscreatekeyresponsebody.md), error**
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

## V1KeysVerifyKey

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-sdk-go/models/components"
	unkeysdkgo "github.com/unkeyed/unkey-sdk-go"
	"context"
	"log"
)

func main() {
    s := unkeysdkgo.New(
        unkeysdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := components.V1KeysVerifyKeyRequest{
        APIID: unkeysdkgo.String("api_1234"),
        Key: "sk_1234",
        Authorization: &components.Authorization{
            Permissions: &components.Permissions{},
        },
    }
    
    ctx := context.Background()
    res, err := s.Keys.V1KeysVerifyKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.V1KeysVerifyKeyRequest](../../models/components/v1keysverifykeyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*components.V1KeysVerifyKeyResponse](../../models/components/v1keysverifykeyresponse.md), error**
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

## V1KeysUpdateKey

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-sdk-go/models/components"
	unkeysdkgo "github.com/unkeyed/unkey-sdk-go"
	"github.com/unkeyed/unkey-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeysdkgo.New(
        unkeysdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.V1KeysUpdateKeyRequestBody{
        KeyID: "key_123",
        Name: unkeysdkgo.String("Customer X"),
        OwnerID: unkeysdkgo.String("user_123"),
        Meta: map[string]interface{}{
            "roles": "<value>",
            "stripeCustomerId": "cus_1234",
        },
        Expires: unkeysdkgo.Float64(0),
        Ratelimit: &operations.V1KeysUpdateKeyRatelimit{
            Type: operations.V1KeysUpdateKeyTypeFast,
            Limit: 10,
            RefillRate: 1,
            RefillInterval: 60,
        },
        Remaining: unkeysdkgo.Float64(1000),
        Refill: &operations.V1KeysUpdateKeyRefill{
            Interval: operations.V1KeysUpdateKeyIntervalDaily,
            Amount: 100,
        },
        Enabled: unkeysdkgo.Bool(true),
    }
    
    ctx := context.Background()
    res, err := s.Keys.V1KeysUpdateKey(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V1KeysUpdateKeyRequestBody](../../models/operations/v1keysupdatekeyrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.V1KeysUpdateKeyResponseBody](../../models/operations/v1keysupdatekeyresponsebody.md), error**
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

## V1KeysUpdateRemaining

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-sdk-go/models/components"
	unkeysdkgo "github.com/unkeyed/unkey-sdk-go"
	"github.com/unkeyed/unkey-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeysdkgo.New(
        unkeysdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.V1KeysUpdateRemainingRequestBody{
        KeyID: "key_123",
        Op: operations.OpSet,
        Value: unkeysdkgo.Int64(1),
    }
    
    ctx := context.Background()
    res, err := s.Keys.V1KeysUpdateRemaining(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V1KeysUpdateRemainingRequestBody](../../models/operations/v1keysupdateremainingrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.V1KeysUpdateRemainingResponseBody](../../models/operations/v1keysupdateremainingresponsebody.md), error**
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

## V1KeysGetVerifications

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-sdk-go/models/components"
	unkeysdkgo "github.com/unkeyed/unkey-sdk-go"
	"github.com/unkeyed/unkey-sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeysdkgo.New(
        unkeysdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.V1KeysGetVerificationsRequest{
        KeyID: unkeysdkgo.String("key_1234"),
        OwnerID: unkeysdkgo.String("chronark"),
        Start: unkeysdkgo.Int64(1620000000000),
        End: unkeysdkgo.Int64(1620000000000),
        Granularity: operations.GranularityDay.ToPointer(),
    }
    
    ctx := context.Background()
    res, err := s.Keys.V1KeysGetVerifications(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V1KeysGetVerificationsRequest](../../models/operations/v1keysgetverificationsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.V1KeysGetVerificationsResponseBody](../../models/operations/v1keysgetverificationsresponsebody.md), error**
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
