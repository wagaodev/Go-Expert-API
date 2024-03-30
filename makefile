.PHONY: default run build test docs clean
#Variables
APP_NAME=goapi

## Tasks
default: run

run:
	@go run cmd/server/main.go
run-with-docs: 
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test -v ./...
docs:
	@echo "Generating docs"
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -f ./docs
	@rm -f ./docs/docs_swagger.go
	@rm -f ./docs/docs.go
	@rm -f ./docs/swagger.json
	@rm -f ./docs/swagger.yaml
	@rm -f ./docs/swagger.yml
	@rm -f ./docs/swagger_doc.go