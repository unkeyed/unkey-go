<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"log"
)

func main() {
	s := unkeygo.New(
		unkeygo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	request := operations.CreateAPIRequestBody{
		Name: "my-api",
	}
	ctx := context.Background()
	res, err := s.CreateAPI(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->