.PHONY: up down run test vet build logs db-shell

up:        ## Start PostgreSQL and wait until it is healthy
	docker compose up -d
	@echo "Waiting for Postgres to be healthy..."
	@until [ "$$(docker inspect -f '{{.State.Health.Status}}' pangea-postgres 2>/dev/null)" = "healthy" ]; do \
		sleep 1; printf "."; \
	done; echo " ready."

down:      ## Stop PostgreSQL
	docker compose down

run:       ## Run the marketplace service
	go run ./cmd/marketplace

test:      ## Run all tests
	go test ./...

vet:       ## Static checks
	go vet ./...

build:     ## Build the binary into ./bin
	go build -o bin/marketplace ./cmd/marketplace

logs:      ## Tail Postgres logs
	docker compose logs -f postgres

db-shell:  ## Open a psql shell
	docker exec -it pangea-postgres psql -U pangea -d pangea