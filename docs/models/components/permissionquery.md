# PermissionQuery

A query for which permissions you require


## Supported Types

### 

```go
permissionQuery := components.CreatePermissionQueryStr(string{/* values here */})
```

### And

```go
permissionQuery := components.CreatePermissionQueryAnd(components.And{/* values here */})
```

### Or

```go
permissionQuery := components.CreatePermissionQueryOr(components.Or{/* values here */})
```

