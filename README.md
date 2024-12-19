<div align="center">
    <img src="https://github.com/unkeyed/unkey-go/assets/68016351/9df64b35-78cd-4404-8b4f-1796b8c685b1" width="200">
    <p>Redefined API management for developers</p>
    <a href="https://godoc.org/github.com/unkeyed/unkey-go"><img src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge" /></a>
    <a href="https://www.unkey.com/docs/api-reference/overview"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>


<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Retries](#retries)
  * [Pagination](#pagination)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/unkeyed/unkey-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := unkeygo.New(
		unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
	)

	res, err := s.Liveness.V1Liveness(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Apis](docs/sdks/apis/README.md)

* [GetAPI](docs/sdks/apis/README.md#getapi)
* [CreateAPI](docs/sdks/apis/README.md#createapi)
* [ListKeys](docs/sdks/apis/README.md#listkeys)
* [DeleteAPI](docs/sdks/apis/README.md#deleteapi)
* [DeleteKeys](docs/sdks/apis/README.md#deletekeys)

### [Identities](docs/sdks/identities/README.md)

* [CreateIdentity](docs/sdks/identities/README.md#createidentity)
* [GetIdentity](docs/sdks/identities/README.md#getidentity)
* [ListIdentities](docs/sdks/identities/README.md#listidentities)
* [UpdateIdentity](docs/sdks/identities/README.md#updateidentity)
* [DeleteIdentity](docs/sdks/identities/README.md#deleteidentity)

### [Keys](docs/sdks/keys/README.md)

* [GetKey](docs/sdks/keys/README.md#getkey)
* [Whoami](docs/sdks/keys/README.md#whoami)
* [DeleteKey](docs/sdks/keys/README.md#deletekey)
* [CreateKey](docs/sdks/keys/README.md#createkey)
* [VerifyKey](docs/sdks/keys/README.md#verifykey)
* [UpdateKey](docs/sdks/keys/README.md#updatekey)
* [UpdateRemaining](docs/sdks/keys/README.md#updateremaining)
* [GetVerifications](docs/sdks/keys/README.md#getverifications)
* [AddPermissions](docs/sdks/keys/README.md#addpermissions)
* [RemovePermissions](docs/sdks/keys/README.md#removepermissions)
* [SetPermissions](docs/sdks/keys/README.md#setpermissions)
* [AddRoles](docs/sdks/keys/README.md#addroles)
* [RemoveRoles](docs/sdks/keys/README.md#removeroles)
* [SetRoles](docs/sdks/keys/README.md#setroles)

### [Liveness](docs/sdks/liveness/README.md)

* [V1Liveness](docs/sdks/liveness/README.md#v1liveness)

### [Migrations](docs/sdks/migrations/README.md)

* [V1MigrationsCreateKeys](docs/sdks/migrations/README.md#v1migrationscreatekeys)
* [V1MigrationsEnqueueKeys](docs/sdks/migrations/README.md#v1migrationsenqueuekeys)

### [Permissions](docs/sdks/permissions/README.md)

* [CreatePermission](docs/sdks/permissions/README.md#createpermission)
* [DeletePermission](docs/sdks/permissions/README.md#deletepermission)
* [GetPermission](docs/sdks/permissions/README.md#getpermission)
* [ListPermissions](docs/sdks/permissions/README.md#listpermissions)
* [CreateRole](docs/sdks/permissions/README.md#createrole)
* [DeleteRole](docs/sdks/permissions/README.md#deleterole)
* [GetRole](docs/sdks/permissions/README.md#getrole)
* [ListRoles](docs/sdks/permissions/README.md#listroles)

### [Ratelimit](docs/sdks/ratelimit/README.md)

* [SetOverride](docs/sdks/ratelimit/README.md#setoverride)
* [ListOverrides](docs/sdks/ratelimit/README.md#listoverrides)
* [GetOverride](docs/sdks/ratelimit/README.md#getoverride)

### [Ratelimits](docs/sdks/ratelimits/README.md)

* [Limit](docs/sdks/ratelimits/README.md#limit)
* [DeleteOverride](docs/sdks/ratelimits/README.md#deleteoverride)


</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `V1Liveness` function may return the following errors:

| Error Type                       | Status Code | Content Type     |
| -------------------------------- | ----------- | ---------------- |
| sdkerrors.ErrBadRequest          | 400         | application/json |
| sdkerrors.ErrUnauthorized        | 401         | application/json |
| sdkerrors.ErrForbidden           | 403         | application/json |
| sdkerrors.ErrNotFound            | 404         | application/json |
| sdkerrors.ErrConflict            | 409         | application/json |
| sdkerrors.ErrTooManyRequests     | 429         | application/json |
| sdkerrors.ErrInternalServerError | 500         | application/json |
| sdkerrors.SDKError               | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := unkeygo.New(
		unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
	)

	res, err := s.Liveness.V1Liveness(ctx)
	if err != nil {

		var e *sdkerrors.ErrBadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrUnauthorized
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrForbidden
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrNotFound
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrConflict
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrTooManyRequests
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ErrInternalServerError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := unkeygo.New(
		unkeygo.WithServerURL("https://api.unkey.dev"),
		unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
	)

	res, err := s.Liveness.V1Liveness(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      |
| ------------ | ---- | ----------- |
| `BearerAuth` | http | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := unkeygo.New(
		unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
	)

	res, err := s.Liveness.V1Liveness(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := unkeygo.New(
		unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
	)

	res, err := s.Liveness.V1Liveness(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := unkeygo.New(
		unkeygo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
	)

	res, err := s.Liveness.V1Liveness(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
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

	res, err := s.Identities.ListIdentities(ctx, operations.ListIdentitiesRequest{
		Limit: unkeygo.Int64(100),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!
