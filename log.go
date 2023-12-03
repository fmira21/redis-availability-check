package main

import(
	"log"
)

func logFatalErr(err error) {
	if err != nil {
		log.Fatalf("[ FATAL ] %s", err)
	}
}

func logErrNonInterrupt(err error) {
	if err != nil {
		log.Printf("[ ERROR ] %s", err)
	}
}

func logTraits(c *config) {
	log.Println("[ INFO ] Configuration is valid.")
	if c.Cluster.Enabled {
		log.Println("[ INFO ] Running for cluster Redis. 'nonClusterAddress' parameter is ignored.")
		log.Printf("[ INFO ] Cluster endpoints: %v\n", c.Cluster.ClusterAddresses)
		if c.Cluster.RandomRouting {
			log.Println("[ INFO ] Connecting randomly to cluster endpoints.")
		} else {
			log.Println("[ INFO ] Connecting to cluster endpoints by latency.")
		}
	} else {
		log.Println("[ INFO ] Running for non-cluster Redis. All parameters from the 'cluster' config section are ignored.")
		log.Printf("[ INFO ] DB name: %d\n", c.Database)
		log.Printf("[ INFO ] Endpoint: %s\n", c.NonClusterAddress)
	}
}