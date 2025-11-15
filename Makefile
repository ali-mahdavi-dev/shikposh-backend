run:
	go run cmd/main.go

# Testing Commands
test:
	go test ./test/... -v

test-unit:
	go test ./test/unit/... -v

test-integration: docker-test-up
	@echo "Waiting for PostgreSQL to be ready..."
	@sleep 5
	go test ./test/integration/... -v
	@$(MAKE) docker-test-down

test-e2e:
	go test ./test/e2e/... -v

test-acceptance:
	go test ./test/acceptance/... -v

test-coverage:
	go test ./test/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Docker Test Commands
docker-test-up:
	cd docker && docker-compose -f docker-compose.test.yml up -d

docker-test-down:
	cd docker && docker-compose -f docker-compose.test.yml down -v

# Database Migrations
migrate-up:
	go run ./cmd/main.go migrate up

migrate-down:
	go run ./cmd/main.go migrate down

# Documentation
swagger:
	swag fmt && swag init -g ./cmd/main.go -o ./docs --parseInternal=true --parseDependency=true
