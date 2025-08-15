include .env
export

database_url = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST_HOST):$(POSTGRES_HOST_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSL)

migration_create:
	migrate create -ext sql -dir migrations $(name)

migration_up:
	migrate -path migrations -database ${database_url} up

migration_down:
	migrate  -path migrations -database $(database_url) down $(n)

migration_force:
	migrate -path migrations -database $(database_url) force -- $(version)

oapi-gen:
	oapi-codegen -config oapi-codegen.yaml oas/openapi.yaml