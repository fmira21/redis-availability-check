package main

import(
	"os"
	"log"
	"gopkg.in/yaml.v2"
)

func readConf(fname string) (*config, error) {
	file, err := os.ReadFile(fname)
	logFatalErr(err)
	c := &config{}
	err = yaml.Unmarshal(file, c)
	logFatalErr(err)
	return c, err
}

func validateConf(c *config) () {
	if !c.Cluster.Enabled {
		if c.Password == "" {
			log.Println("[ INFO ] Empty DB password.")
		}
		if c.NonClusterAddress == "" {
			log.Fatalln("[ FATAL ] 'nonClusterAddress' field is empty, aborting.")
		}
	} else {
		if c.Cluster.ClusterAddresses == nil {
			log.Fatalln("[ FATAL ] 'clusterAddresses' list is empty, aborting.")
		}
	}
}