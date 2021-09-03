package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.TODO()
var redisAddr = "redis-11202.c60.us-west-1-2.ec2.cloud.redislabs.com:11202" //redis uri

type Database struct {
	Client *redis.Client
}

func Connect() *Database {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "nV1xvb7XxM7eUJngCpseJMGkUXcHO2Qf", //password of our redis database
		DB:       0,
	})

	//if err := client.Ping(Ctx); err != nil {
	//	log.Fatal(err.Args()...)
	//	fmt.Println("could not connect to redis...")
	//}
	return &Database{
		Client: client,
	}

}
