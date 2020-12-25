package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:36379",
		// Addr:     "192.168.80.238:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)
	client.Do("AUTH", "123456")
	cmd := client.Do("select", "15")
	fmt.Println(cmd)

	// cmd = client.Do("set", "kk", "15")
	// fmt.Println(cmd)
	// str, err := client.Pipeline().Select(15).Result()
	// fmt.Println(str, err)
	// client.Set("test", "test11", time.Second*10)

	for i := 0; i < 100; i++ {
		client.Do("AUTH", "123456")
		cmd := client.Do("select", "15")
		fmt.Println(cmd)
		client.HSet(fmt.Sprintf("key%d", i), "field", ":dfafdasfdas")
		client.Expire(fmt.Sprintf("key%d", i), time.Second*100)
	}

}
