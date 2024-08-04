package discovery

import (
	"log"

	"github.com/hashicorp/consul/api"
)

var client *api.Client

func InitDiscovery(config *api.Config) {
	var err error
	client, err = api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCacheNodes() []string {
	services, _, err := client.Catalog().Service("cache-node", "", nil)
	if err != nil {
		log.Fatal(err)
	}

	var nodes []string
	for _, service := range services {
		nodes = append(nodes, service.ServiceAddress)
	}
	return nodes
}
