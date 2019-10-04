validate: ## validate schema.yml
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli validate -i /local/schema/schema.yml
generate: ## generate go code
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/schema/schema.yml -g go-gin-server -o /local/server
