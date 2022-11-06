Simple cache module for GO
===========

Simple cache module for Go Lang

## Create new cache storage

```
cache.New()
```


## Add new item to cache

```
cache.Set(key string, value interfece{}, ttl time.Duration)
```

##### Example:
```
cache.Set("userId", 42, time.Second*5)
```

## Get item from cache storage

```
cache.Get(key string) (interface{}, error)
```

##### Example:
```
userId, err := cache.Get("userId")
```

## Delete item from cache storage

```
cache.Delete(key string) error
```

##### Example:
```
err := cache.Delete("userId")
```