# KeyID

Only include data for a specific key or keys.

When you are providing zero or more than one key ids, all usage counts are aggregated and summed up. Send multiple requests with one keyId each if you need counts per key.




## Supported Types

### 

```go
keyID := operations.CreateKeyIDStr(string{/* values here */})
```

### 

```go
keyID := operations.CreateKeyIDArrayOfStr([]string{/* values here */})
```

