.PHONY: openapi
openapi-codegen:
	oapi-codegen --config=api/server.cfg.yaml api/objects.yaml
	oapi-codegen --config=api/models.cfg.yaml api/objects.yaml

go-run:
	go run main.go