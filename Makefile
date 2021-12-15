dev-setup:
	docker-compose -f docker-compose.yml up -d postgres

dev-setup-down:
	docker-compose -f docker-compose.yml down

run-pre-commit:
	gofmt -w .
