package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	setErr := client.Set("key", "value", 0).Err()
	if setErr != nil {
		panic(err)
	}

	val, getErr := client.Get("key").Result()
	if getErr != nil {
		panic(getErr)
	}
	fmt.Println("key:", val)
}
