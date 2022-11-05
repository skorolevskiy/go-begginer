package main

import (
	"fmt"
	cache "github.com/skorolevskiy/go-begginer/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
