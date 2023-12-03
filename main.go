package main

import (
	"os"
	"log"
	"context"
	"github.com/redis/go-redis/v9"
)

type config struct {
	Database int
	Password string
	NonClusterAddress string `yaml:"nonClusterAddress"`
	Verbose bool
	StopIfUnvavilable bool `yaml:"stopIfUnvavilable"`
	Cluster struct {
		Enabled bool
		RandomRouting bool `yaml:"randomRouting"`
		ClusterAddresses []string `yaml:"clusterAddresses"`
	}
}

type clientConnections struct {
	ClusterClient *redis.ClusterClient
	IndependentClient *redis.Client
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("[ FATAL ] Did not find the config:(")
	}

	config, err := readConf(os.Args[1])
	logFatalErr(err)

	validateConf(config)
	logTraits(config)

	var cc clientConnections
	context := context.Background()

	if config.Cluster.Enabled {
		cc.ClusterClient = connectCluster(config)
		loadClusterClient(config, context, cc.ClusterClient)
	} else {
		cc.IndependentClient = connectNonCluster(config)
		loadClient(config, context, cc.IndependentClient)
	}
}

