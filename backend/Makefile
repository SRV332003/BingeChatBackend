HangkAroundBackend: docs/swagger.json
	@echo "Building HangkAroundBackend
	 @go build -o HangkAroundBackend cmd/main.go

docs/swagger.json: $(shell find . -name '*.go')
	@echo "Generating Swagger Docs"
	@swag init