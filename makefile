build:
	@go build -o dist/user_service_server ./main.go
	@go build -o dist/user_service_migration ./cmd/migrate/main.go