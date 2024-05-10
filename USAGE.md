<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkgo "github.com/unkeyed/sdk-go"
	"github.com/unkeyed/sdk-go/models/components"
	"github.com/unkeyed/sdk-go/models/operations"
	"log"
)

func main() {
	s := sdkgo.New(
		sdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	request := operations.V1ApisCreateAPIRequestBody{
		Name: "my-api",
	}

	ctx := context.Background()
	res, err := s.V1ApisCreateAPI(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->