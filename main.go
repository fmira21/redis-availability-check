package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	clusterAddrs := flag.String("cluster", "", "Redis cluster addresses, comma separated")
	serverAddr := flag.String("server", "", "Redis single server address")
	username := flag.String("username", "", "Redis username")
	password := flag.String("password", "", "Redis password")
	flag.Parse()

	var conn redis.UniversalClient

	if *clusterAddrs != "" {
		conn = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:         strings.Split(*clusterAddrs, ","),
			Username:      *username,
			Password:      *password,
			RouteRandomly: true,
		})
	} else if *serverAddr != "" {
		conn = redis.NewClient(&redis.Options{
			Addr:     *serverAddr,
			Username: *username,
			Password: *password,
		})
	}

	ctx := context.Background()

	for {
		value := time.Now().Format(time.RFC3339Nano)

		if err := conn.Set(ctx, "redis-test", value, 0).Err(); err != nil {
			log.Println("Set error:", err)
		}

		if val, err := conn.Get(ctx, "redis-test").Result(); err != nil || value != val {
			log.Println("Get error:", err)
		} else {
			log.Println("Value:", val)
		}

		time.Sleep(200 * time.Millisecond)
	}
}
