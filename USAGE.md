<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	unkeygo "github.com/unkeyed/unkey-go"
	"log"
)

func main() {
	s := unkeygo.New(
		unkeygo.WithSecurity("UNKEY_ROOT_KEY"),
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