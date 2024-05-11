<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/unkeyed/unkey"
	"github.com/unkeyed/unkey/models/components"
	"github.com/unkeyed/unkey/models/operations"
	"log"
)

func main() {
	s := unkey.New(
		unkey.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	request := operations.CreateAPIRequestBody{
		Name: "my-api",
	}

	ctx := context.Background()
	res, err := s.CreateAPI(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->