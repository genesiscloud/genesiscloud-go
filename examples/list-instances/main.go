package main

import (
	"context"
	"fmt"
	"os"

	"github.com/genesiscloud/genesiscloud-go"
)

// Generic pointer in Go 1.18+. Alternatively use "k8s.io/utils/pointer"
func pointer[T any](v T) *T {
	return &v
}

func run() error {
	ctx := context.Background()

	client, err := genesiscloud.NewGenesisCloudClient(genesiscloud.ClientConfig{
		Token: os.Getenv("GENESISCLOUD_TOKEN"),
	})
	if err != nil {
		return err
	}

	resp, err := client.ListInstancesPaginatedWithResponse(ctx, &genesiscloud.ListInstancesPaginatedParams{
		Page:    pointer(1),
		PerPage: pointer(100),
	})
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("list-instances: %s", resp.Status())
	}

	for _, instance := range resp.JSON200.Instances {
		fmt.Printf("%s %s\n", instance.Id, instance.Name)
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
