validate: ## validate schema.yml
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli validate -i /local/schema/schema.yml
