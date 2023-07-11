# Genesis Cloud Go Client

The Genesis Cloud Go client provides the ability to manage each services resources programmatically from scripts or applications.

- [Godoc](https://pkg.go.dev/github.com/genesiscloud/genesiscloud-go?tab=doc)
- [Genesis Cloud API Documentation](https://developers.genesiscloud.com/)
- [How to generate an API key?](https://support.genesiscloud.com/support/solutions/articles/47001126146-how-to-generate-an-api-token-)

This repository is licensed under Mozilla Public License 2.0 (no copyleft exception) (see [LICENSE.txt](./LICENSE.txt)).

# Maintainers

This client is maintained by Genesis Cloud. If you encounter any problems, feel free to create issues or pull requests.

## Requirements

- [Go](https://golang.org/doc/install) >= 1.19

## Installation

```bash
go get github.com/genesiscloud/genesiscloud-go
```

## Getting Started

```go
import "github.com/genesiscloud/genesiscloud-go"

// Generic pointer in Go 1.18+. Alternatively use "k8s.io/utils/pointer"
func pointer[T any](v T) *T {
	return &v
}

client, err := genesiscloud.NewGenesisCloudClient(genesiscloud.ClientConfig{
	Token: os.Getenv("GENESISCLOUD_TOKEN"),
})
if err != nil {
	// ...
}

resp, err := client.ListInstancesWithResponse(ctx, &genesiscloud.ListInstancesParams{
	Page:    pointer(1),
	PerPage: pointer(100),
})
if err != nil {
	// ...
}
```

## Examples

You can find additional examples in the [GoDoc](https://pkg.go.dev/github.com/genesiscloud/genesiscloud-go?tab=doc) or
check the [examples folder](./examples).

```sh
GENESISCLOUD_TOKEN="XXXX" go run ./examples/list-instances
```

## Development or update of OpenAPI document

```sh
# Update openapi.yaml (./codegen/openapi.yaml)

# Generate code
go generate ./...
```
