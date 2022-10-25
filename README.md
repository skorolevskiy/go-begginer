Simple cache module for GO
===========

Simple cache module for Go Lang

## Create new cache storage

```
cache.New()
```


## Add new item to cache

```
cache.Set(key string, value interfece{})
```

##### Example:
```
cache.Set("userId", 42)
```

## Get item from cache storage

```
cache.Get(key string)
```

##### Example:
```
cache.Get("userId")
```

## Delete item from cache storage

```
cache.Delete(key string)
```

##### Example:
```
cache.Delete("userId")
```