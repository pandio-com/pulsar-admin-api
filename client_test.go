package pulsar

import (
	"context"
	"fmt"
	"log"
	"testing"

	swagger "github.com/pandio-com/pulsar-management-service/pkg/pulsar/client"
)

func TestClient(t *testing.T) {
	client := swagger.NewAPIClient(&swagger.Configuration{
		BasePath:      "http://localhost:8080/admin/v2",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	})
	clusters, _, err := client.ClustersApi.GetClusters(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", clusters)
}
