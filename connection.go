package main

import(
	"log"
	"github.com/redis/go-redis/v9"
)

func connectNonCluster(c * config) (conn *redis.Client) {
	log.Printf("[ INFO ] Connecting to the Redis endpoint %s...\n", c.NonClusterAddress)
	conn = redis.NewClient(&redis.Options{
		Addr:	  c.NonClusterAddress,
		Password: c.Password,
		DB:		  c.Database,
	})
	log.Println("[ INFO ] Done.")
	return
}

func connectCluster(c *config) (conn *redis.ClusterClient) {
	lat := true
	rand := false

	if c.Cluster.RandomRouting {
		lat = false
		rand = true
	}

	log.Println("[ INFO ] Connecting to the Redis cluster...")
	conn = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: c.Cluster.ClusterAddresses,
		Password: c.Password,
		RouteByLatency: lat,
		RouteRandomly: rand,
	})
	log.Println("[ INFO ] Done.")
	return
}