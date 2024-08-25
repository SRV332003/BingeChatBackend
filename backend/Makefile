HangAroundBackend: docs/swagger.json go.sum
	@echo "Building HangAroundBackend"
	@go build -o HangAroundBackend .

docs/swagger.json: $(shell find . -name '*.go') 
	@echo "Generating Swagger Docs"
	@swag init

go.sum, go.mod: go.mod
	@echo "Updating go.sum"
	@go mod tidy
