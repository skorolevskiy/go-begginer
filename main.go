package main

import (
	"fmt"
	"log"
	"time"

	cache "github.com/skorolevskiy/go-begginer/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*5)
	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	fmt.Println(userId) // Output: 42

	time.Sleep(time.Second * 6) // прошло 5 секунд

	//userId = cache.Get("userId")
	userId, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}
}
