// Download a path from Artifactory
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atlassian/go-artifactory/v2/artifactory"
	"github.com/atlassian/go-artifactory/v2/artifactory/transport"
)

func main() {
	tp := transport.BasicAuth{
		Username: os.Getenv("ARTIFACTORY_USERNAME"),
		Password: os.Getenv("ARTIFACTORY_PASSWORD"),
	}

	client, err := artifactory.NewClient(os.Getenv("ARTIFACTORY_URL"), tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	key := os.Getenv("ARTIFACT_PATH")
	n := filepath.Base(key)
	o, err := os.Create(fmt.Sprintf("/tmp/%s", n))
	defer o.Close()

	_, err = client.V1.Artifacts.RetrieveArtifact(context.Background(), key, o)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	fmt.Printf("Downloaded to: %s\n", o.Name())
}
