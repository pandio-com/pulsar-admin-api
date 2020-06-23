package pulsar

import (
	"context"
	"fmt"
	"log"
	"testing"

	pulsar "github.com/pandio-com/pulsar-admin-api/client"
)

func TestClient(t *testing.T) {
	c := pulsar.NewAPIClient(&pulsar.Configuration{
		BasePath:      "http://localhost:7000/admin/v2",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	})
	clusters, _, err := c.ClustersApi.GetClusters(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", clusters)
}
