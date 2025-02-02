# Tag

Only include data for a specific tag or tags.

When you are providing zero or more than one tag, all usage counts are aggregated and summed up. Send multiple requests with one tag each if you need counts per tag.


## Supported Types

### 

```go
tag := operations.CreateTagStr(string{/* values here */})
```

### 

```go
tag := operations.CreateTagArrayOfStr([]string{/* values here */})
```

