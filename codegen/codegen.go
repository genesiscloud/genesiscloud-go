package codegen

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=config/types.yaml https://api.genesiscloud.com/compute/v1/openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=config/client.yaml https://api.genesiscloud.com/compute/v1/openapi.yaml
