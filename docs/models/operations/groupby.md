# GroupBy

By default, datapoints are not aggregated, however you probably want to get a breakdown per time, key or identity.

Grouping by tags and by tag is mutually exclusive.


## Supported Types

### One

```go
groupBy := operations.CreateGroupByOne(operations.One{/* values here */})
```

### 

```go
groupBy := operations.CreateGroupByArrayOf2([]operations.Two{/* values here */})
```

