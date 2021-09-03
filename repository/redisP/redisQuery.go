package redisP

import (
	"auth/database"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func Store(key string, val string) {
	if key == "" || val == "" {
		fmt.Println("no value provided for key and val")

	}
	//	err := database.Connect().Client.SetEX(context.TODO(), key, val, time.Duration(5*time.Minute))
	//fmt.Println(err)
	e := database.Connect().Client.Set(context.TODO(), key, val, time.Duration(10*time.Minute)).Err()
	if e != nil {
		fmt.Println(e)
		log.Fatal(e)
	}

	fmt.Println("stored data in redis successfully")

}

func FindId(nonce string) string {
	val, err := database.Connect().Client.Get(context.TODO(), nonce).Result()
	if err == redis.Nil {
		fmt.Println("this is get error and key does not exist", err)
	}
	phno, err := database.Connect().Client.Get(context.TODO(), val).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	}
	fmt.Println("this is  get result", val)
	return phno
}
