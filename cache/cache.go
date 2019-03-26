package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {
	// 实例化一个cache对象，设置过期时过期后删除的时间
	c := cache.New(5*time.Minute, 30*time.Second)
	// 将kv放进cachee中
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}

}
