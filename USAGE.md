<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"log"
	"os"
)

func main() {
	s := unkeygo.New(
		unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
	)

	ctx := context.Background()
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