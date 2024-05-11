# Ratelimits
(*Ratelimits*)

### Available Operations

* [Limit](#limit)

## Limit

### Example Usage

```go
package main

import(
	"github.com/unkeyed/unkey-go/models/components"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    request := operations.LimitRequestBody{
        Namespace: unkeygo.String("email.outbound"),
        Identifier: "user_123",
        Limit: 10,
        Duration: 60000,
        Cost: unkeygo.Int64(2),
        Resources: []operations.Resources{
            operations.Resources{
                Type: "project",
                ID: "p_123",
                Name: unkeygo.String("dub"),
            },
        },
    }
    
    ctx := context.Background()
    res, err := s.Ratelimits.Limit(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.LimitRequestBody](../../models/operations/limitrequestbody.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.LimitResponseBody](../../models/operations/limitresponsebody.md), error**
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
