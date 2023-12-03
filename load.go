package main

import(
	"os"
	"log"
	"context"
	"syscall"
	"os/signal"
	"github.com/redis/go-redis/v9"
)

func stopIfUnvavilable(c *config, err error) {
	if c.StopIfUnvavilable {
		logFatalErr(err)
	} else {
		logErrNonInterrupt(err)
	}
}

func clientSimplePair(c *config, ctx context.Context, conn *redis.Client) {
	if c.Verbose {
		log.Printf("[ DEBUG ] Creating simple key-value pair...")
	}
	err := conn.Set(ctx, "testingKey", "testingValue", 0).Err()
	stopIfUnvavilable(c, err)
	if c.Verbose {
		log.Printf("[ DEBUG ] Getting value by key...")
	}
	_, err = conn.Get(ctx, "testingKey").Result()
	stopIfUnvavilable(c, err)
}

func clusterSimplePair(c *config, ctx context.Context, conn *redis.ClusterClient) {
	if c.Verbose {
		log.Printf("[ DEBUG ] Creating simple key-value pair...")
	}
	err := conn.Set(ctx, "testingKey", "testingValue", 0).Err()
	stopIfUnvavilable(c, err)
	if c.Verbose {
		log.Printf("[ DEBUG ] Getting value by key...")
	}
	_, err = conn.Get(ctx, "testingKey").Result()
	stopIfUnvavilable(c, err)
}

func clientCleanup(c *config, ctx context.Context, conn *redis.Client, onexit bool) {
	if c.Verbose {
		log.Printf("[ DEBUG ] Cleaning up the testing key...")
	}
	_, err := conn.Del(ctx, "testingKey").Result()
	if onexit {
		logFatalErr(err)
	}
	stopIfUnvavilable(c, err)
}

func clusterCleanup(c *config, ctx context.Context, conn *redis.ClusterClient, onexit bool) {
	if c.Verbose {
		log.Printf("[ DEBUG ] Cleaning up the testing key...")
	}
	_, err := conn.Del(ctx, "testingKey").Result()
	if onexit {
		logFatalErr(err)
	}
	stopIfUnvavilable(c, err)
}

func loadClient(c *config, ctx context.Context, conn *redis.Client) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
    go func() {
		<-channel
		log.Println("[ INFO ] Caught SIGTERM, cleaning up...")
		clientCleanup(c, ctx, conn, true)
	}()
	log.Println("[ INFO ] Applying load to the cluster...")
	for {
		clientSimplePair(c, ctx, conn)
		clientCleanup(c, ctx, conn, false)
	}
}

func loadClusterClient(c *config, ctx context.Context, conn *redis.ClusterClient) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
    go func() {
		<-channel
		log.Println("[ INFO ] Caught SIGTERM, cleaning up...")
		clusterCleanup(c, ctx, conn, true)
	}()
	log.Println("[ INFO ] Applying load to the cluster...")
	for {
		clusterSimplePair(c, ctx, conn)
		clusterCleanup(c, ctx, conn, false)
	}
}