Simple cache module for GO
===========

Simple cache module for Go Lang

## Create new cache storage

```go
cache.New()
```


## Add new item to cache

```go
cache.Set(key string, value interfece{}, ttl time.Duration)
```

##### Example:
```go
cache.Set("userId", 42, time.Second*5)
```

## Get item from cache storage

```go
cache.Get(key string) (interface{}, error)
```

##### Example:
```go
userId, err := cache.Get("userId")
```

## Delete item from cache storage

```go
cache.Delete(key string) error
```

##### Example:
```go
err := cache.Delete("userId")
```